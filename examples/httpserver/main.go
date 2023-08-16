package main

import (
	"context"
	"crypto/rand"
	"fmt"
	host_http "github.com/taction/wit-dapr/pkg/imports/host-http"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
)

func main() {
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
		os.Exit(0)
	}()
	run(os.Args[1])
}

func run(wasmFile string) error {
	if len(wasmFile) == 0 {
		return fmt.Errorf("no wasm file specified")
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request, wasmFile)
		if err != nil {
			writer.WriteHeader(500)
			writer.Write([]byte(err.Error()))
		}
	})

	return http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request, wasmFile string) error {
	wasmName := filepath.Base(wasmFile)
	wasmCode, err := os.ReadFile(wasmFile)
	if err != nil {
		return fmt.Errorf("could not read WASM file '%s': %w", wasmFile, err)
	}
	ctx := context.Background()
	runtime := wazero.NewRuntime(ctx)
	defer runtime.Close(ctx)

	// todo detect wasm module valid
	wasmModule, err := runtime.CompileModule(ctx, wasmCode)
	if err != nil {
		return err
	}
	defer wasmModule.Close(ctx)
	wasi_snapshot_preview1.Instantiate(ctx, runtime)

	ins, err := runtime.InstantiateModule(ctx, wasmModule,
		wazero.NewModuleConfig().WithName(wasmName).
			WithRandSource(rand.Reader).
			WithSysNanotime().
			WithSysWalltime().
			WithSysNanosleep())
	if err != nil {
		return err
	}
	ws := host_http.WasmServer{Module: ins}
	ws.ServeHTTP(w, r)
	return nil
}
