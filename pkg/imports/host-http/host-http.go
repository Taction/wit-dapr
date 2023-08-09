package host_http

import (
	"context"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"net/http"
)

type HTTP struct {
}

func MakeHTTP() *HTTP {

	return &HTTP{}
}

func (w *HTTP) Instantiate(ctx context.Context, rt wazero.Runtime) error {
	if err := default_http.Instantiate(ctx, rt, w.r, w.rs, w.f); err != nil {
		return err
	}
	return nil
}

func (w *HTTP) MakeHandler(m api.Module) http.Handler {
	return WasmServer{
		Module: m,
	}
}

func (w *HTTP) HandleHTTP(writer http.ResponseWriter, req *http.Request, m api.Module) {
	handler := w.MakeHandler(m)
	handler.ServeHTTP(writer, req)
}
