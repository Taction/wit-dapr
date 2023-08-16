package host_http

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/taction/wit-dapr/pkg/common"
	"github.com/tetratelabs/wazero/api"
	"io"
	"log"
	"net/http"
)

const returnedResponseLen = 28 // from http.c static uint8_t RET_AREA[28];

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
	methond := getMethod(req.Method)
	urlPtr, urlLen, headerPtr, headerLen, paramPtr, paramLen, hasBody, bodyPtr, bodyLen, err := formatRequestAsArgs(w.Module, req)
	resp, err := fn.Call(context.TODO(), uint64(methond), urlPtr, urlLen, headerPtr, headerLen, paramPtr, paramLen, hasBody, bodyPtr, bodyLen)
	if err != nil {
		res.WriteHeader(500)
		res.Write([]byte(err.Error()))
		return
	}
	data, ok := w.Module.Memory().Read(uint32(resp[0]), returnedResponseLen)
	if !ok {
		fmt.Println("Error reading headers.")
		res.WriteHeader(500)
		res.Write([]byte("Error reading response"))
		return
	}
	response := decodeResponse(w.Module, data)
	if response == nil {
		res.WriteHeader(500)
		res.Write([]byte("Error decoding response"))
		return
	}
	if response.Headers != nil {
		// todo handle headers properly
		for k, v := range *response.Headers {
			res.Header().Set(k, v[0])
		}
	}
	res.WriteHeader(int(response.StatusCode))
	if response.Body != nil {
		res.Write(response.Body)
	}
}

// out parameters are urlPtr, urlLen, headerPtr, headerLen, paramPtr, paramLen, hasBody, bodyPtr, bodyLen
func formatRequestAsArgs(mod api.Module, req *http.Request) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64, error) {
	ctx := context.Background()
	url := req.URL.String()
	urlLen := uint32(len(url))
	urlPtr, err := Malloc(ctx, mod, uint32(urlLen))
	if err != nil {
		return 0, 0, 0, 0, 0, 0, 0, 0, 0, err
	}
	mod.Memory().Write(uint32(urlPtr), []byte(url))

	// headers
	le := binary.LittleEndian
	headerLen := uint32(len(req.Header))
	// 8 bytes per string/string
	headerPtr, err := Malloc(ctx, mod, headerLen*16)
	if err != nil {
		return 0, 0, 0, 0, 0, 0, 0, 0, 0, err
	}
	// ok now allocate and write the strings.
	data := []byte{}
	for k, v := range req.Header {
		ptr, err := allocateWriteString(ctx, mod, k)
		if err != nil {
			return 0, 0, 0, 0, 0, 0, 0, 0, 0, err
		}
		data = le.AppendUint32(data, ptr)
		data = le.AppendUint32(data, uint32(len(k)))
		ptr, err = allocateWriteString(ctx, mod, v[0])
		if err != nil {
			return 0, 0, 0, 0, 0, 0, 0, 0, 0, err
		}
		data = le.AppendUint32(data, ptr)
		data = le.AppendUint32(data, uint32(len(v[0])))
	}
	mod.Memory().Write(headerPtr, data)
	// params will be contained in the url, so we don't need to do anything here.
	paramPtr, paramLen := 0, 0
	hasBody := 0
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return 0, 0, 0, 0, 0, 0, 0, 0, 0, err
	}
	var bodyPtr, bodyLen uint32 = 0, uint32(len(body))
	if len(body) > 0 {
		bodyPtr, err = Malloc(ctx, mod, uint32(len(body)))
		if bodyPtr != 0 {
			hasBody = 1
		}
	}
	return uint64(urlPtr), uint64(urlLen), uint64(headerPtr), uint64(headerLen), uint64(paramPtr), uint64(paramLen), uint64(hasBody), uint64(bodyPtr), uint64(bodyLen), nil
}

func Malloc(ctx context.Context, m api.Module, size uint32) (uint32, error) {
	malloc := m.ExportedFunction("malloc")
	result, err := malloc.Call(ctx, uint64(size))
	if err != nil {
		log.Println(err.Error())
	}
	return uint32(result[0]), err
}
func allocateWriteString(ctx context.Context, m api.Module, s string) (uint32, error) {
	ptr, err := Malloc(ctx, m, uint32(len(s)))
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	m.Memory().Write(ptr, []byte(s))
	return ptr, nil
}

func getMethod(method string) DaprHttpHttpTypesMethodKind {
	switch method {
	case "GET":
		return DaprHttpHttpTypesMethodKindGet
	case "PUT":
		return DaprHttpHttpTypesMethodKindPut
	case "POST":
		return DaprHttpHttpTypesMethodKindPost
	case "DELETE":
		return DaprHttpHttpTypesMethodKindDelete
	default:
		panic("unsupported method")
	}
}

func decodeResponse(m api.Module, res []byte) *Response {
	if len(res) != returnedResponseLen {
		fmt.Printf("Error response length %d, expected %v\n", len(res), returnedResponseLen)
		return nil
	}
	return &Response{
		StatusCode: uint16(binary.LittleEndian.Uint32(res[0:4])), // 内存对其原因
		Headers:    decodeHeaders(m, res[4:16]),
		Body:       decodeBody(m, res[16:]),
	}
}

func decodeHeaders(mod api.Module, h []byte) *Headers {
	if h[0] == 0 {
		// nil
		return nil
	}
	ptr := binary.LittleEndian.Uint32(h[4:8])
	len := binary.LittleEndian.Uint32(h[8:12])
	data, ok := mod.Memory().Read(ptr, len*16) // len is header count 16 bytes per header, 8 bytes for key, 8 bytes for value
	if !ok {
		fmt.Println("Error reading header data")
		return nil
	}
	fields := make(Headers)
	for i := uint32(0); i < len; i++ {
		key_ptr := binary.LittleEndian.Uint32(data[i*16 : i*16+4])
		key_len := binary.LittleEndian.Uint32(data[i*16+4 : i*16+8])
		key, ok := common.ReadString(mod, key_ptr, key_len)
		if !ok {
			fmt.Println("Error reading key")
			return nil
		}
		val_ptr := binary.LittleEndian.Uint32(data[i*16+8 : i*16+12])
		val_len := binary.LittleEndian.Uint32(data[i*16+12 : i*16+16])
		val, ok := common.ReadString(mod, val_ptr, val_len)
		if !ok {
			fmt.Println("Error reading value")
			return nil
		}
		if _, found := fields[key]; !found {
			fields[key] = []string{}
		}
		fields[key] = append(fields[key], val)
	}
	return &fields
}

func decodeBody(mod api.Module, h []byte) []byte {
	if h[0] == 0 {
		// nil
		return nil
	}
	ptr := binary.LittleEndian.Uint32(h[4:8])
	len := binary.LittleEndian.Uint32(h[8:12])
	data, ok := mod.Memory().Read(ptr, len)
	if !ok {
		fmt.Println("Error reading header data")
		return nil
	}
	return data
}
