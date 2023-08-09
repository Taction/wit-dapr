package main

import (
	"github.com/taction/wit-dapr/dapr/http"
)

func init() {
	a := &HostImpl{}
	http.SetExportsDaprHttpHttpHandler(a)
}

type HostImpl struct {
}

func (h *HostImpl) HandleHttpRequest(req http.DaprHttpHttpTypesRequest) http.DaprHttpHttpTypesResponse {
	headers := []http.DaprHttpHttpTypesTuple2StringStringT{http.DaprHttpHttpTypesTuple2StringStringT{"Content-Type", "text/plain"}}
	return http.DaprHttpHttpTypesResponse{
		Status:  200,
		Body:    http.Some([]byte("Hello from WASM!")),
		Headers: http.Some(headers),
	}
}

func main() {}
