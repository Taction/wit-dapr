package host_http

import (
	"context"
	"github.com/tetratelabs/wazero/api"
	"net/http"
)

type WasmServer struct {
	Module api.Module
}

func (w WasmServer) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fn := w.Module.ExportedFunction("dapr:http/http-handler#handle-http-request")
	if fn == nil {
		res.WriteHeader(500)
		res.Write([]byte("Handler not found"))
		return
	}
	id := w.Requests.MakeRequest(req)
	out := w.OutParams.MakeOutparameter()

	_, err := fn.Call(context.TODO(), uint64(id), uint64(out))
	if err != nil {
		res.WriteHeader(500)
		res.Write([]byte(err.Error()))
		return
	}
	responseId, found := w.OutParams.GetResponseByOutparameter(out)
	if !found {
		res.WriteHeader(500)
		res.Write([]byte("Couldn't find outparameter mapping"))
		return
	}
	r, found := w.Responses.GetResponse(responseId)
	if !found || r == nil {
		res.WriteHeader(500)
		res.Write([]byte("Couldn't find response"))
		return
	}
	if headers, found := w.Fields.GetFields(r.HeaderHandle); found {
		for key, value := range headers {
			for ix := range value {
				res.Header().Add(key, value[ix])
			}
		}
		w.Fields.DeleteFields(r.HeaderHandle)
	}
	res.WriteHeader(r.StatusCode)
	data := r.Buffer.Bytes()
	res.Write(data)

	w.Responses.DeleteResponse(responseId)
}
