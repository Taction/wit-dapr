package main

import (
	"context"
	"crypto/rand"
	"errors"
	"flag"
	"fmt"
	host_state "github.com/taction/wit-dapr/pkg/imports/host-state"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/dapr/dapr/pkg/components"
)

func main() {
	var resourcesPath stringSliceFlag
	flag.Var(&resourcesPath, "resources-path", "Path for resources directory. If not specified, no resources will be loaded. Can be passed multiple times")
	wasmfile := flag.String("wasmfile", "", "wasm file path")
	flag.Parse()

	// ...
	loader := components.NewLocalComponents(resourcesPath...)
	log.Println("Loading components…")
	comps, err := loader.LoadComponents()
	if err != nil {
		panic(err)
	}
	var compList []host_state.Component
	for _, comp := range comps {
		// copy comp to host_state.Component
		md := make([]host_state.MetadataItem, len(comp.Spec.Metadata))
		for i, m := range comp.Spec.Metadata {
			md[i] = host_state.MetadataItem{
				Name:  m.Name,
				Value: m.Value.String(),
				SecretKeyRef: host_state.SecretKeyRef{
					Name: m.SecretKeyRef.Name,
					Key:  m.SecretKeyRef.Key,
				},
			}
		}
		compList = append(compList, host_state.Component{
			TypeMeta:   comp.TypeMeta,
			ObjectMeta: comp.ObjectMeta,
			Spec: host_state.ComponentSpec{
				Type:     comp.Spec.Type,
				Version:  comp.Spec.Version,
				Metadata: md,
			},
		})
	}
	host_state.AddComp(compList...)
	run(*wasmfile)
}

func run(wasmFile string) error {
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
	host_state.Instantiate(ctx, runtime)

	ins, err := runtime.InstantiateModule(ctx, wasmModule,
		wazero.NewModuleConfig().WithName(wasmName).
			WithRandSource(rand.Reader).
			WithStdout(os.Stdout).
			WithSysNanotime().
			WithSysWalltime().
			WithSysNanosleep())
	if err != nil {
		return err
	}

	ins.Close(ctx)
	return nil
}

// Flag type. Allows passing a flag multiple times to get a slice of strings.
// It implements the flag.Value interface.
type stringSliceFlag []string

// String formats the flag value.
func (f stringSliceFlag) String() string {
	return strings.Join(f, ",")
}

// Set the flag value.
func (f *stringSliceFlag) Set(value string) error {
	if value == "" {
		return errors.New("value is empty")
	}
	*f = append(*f, value)
	return nil
}
