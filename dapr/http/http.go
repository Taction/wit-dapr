// Generated by `wit-bindgen` 0.7.0. DO NOT EDIT!
package http

// #include "http.h"
import "C"

import "unsafe"

type DaprHttpHttpTypesHttpStatus = uint16
type DaprHttpHttpTypesBody = uint8
type DaprHttpHttpTypesTuple2StringStringT = DaprHttpHttpHandlerTuple2StringStringT
type DaprHttpHttpHandlerTuple2StringStringT struct {
  F0 string
  F1 string
}

type DaprHttpHttpTypesHeaders = DaprHttpHttpTypesTuple2StringStringT
type DaprHttpHttpTypesParams = DaprHttpHttpTypesTuple2StringStringT
type DaprHttpHttpTypesUri = string
type DaprHttpHttpTypesMethodKind int

const (
DaprHttpHttpTypesMethodKindGet DaprHttpHttpTypesMethodKind = iota
DaprHttpHttpTypesMethodKindPost
DaprHttpHttpTypesMethodKindPut
DaprHttpHttpTypesMethodKindDelete
DaprHttpHttpTypesMethodKindPatch
DaprHttpHttpTypesMethodKindHead
DaprHttpHttpTypesMethodKindOptions
)

type DaprHttpHttpTypesMethod struct {
  kind DaprHttpHttpTypesMethodKind
}

func (n DaprHttpHttpTypesMethod) Kind() DaprHttpHttpTypesMethodKind {
  return n.kind
}

func DaprHttpHttpTypesMethodGet() DaprHttpHttpTypesMethod{
  return DaprHttpHttpTypesMethod{kind: DaprHttpHttpTypesMethodKindGet}
}

func DaprHttpHttpTypesMethodPost() DaprHttpHttpTypesMethod{
  return DaprHttpHttpTypesMethod{kind: DaprHttpHttpTypesMethodKindPost}
}

func DaprHttpHttpTypesMethodPut() DaprHttpHttpTypesMethod{
  return DaprHttpHttpTypesMethod{kind: DaprHttpHttpTypesMethodKindPut}
}

func DaprHttpHttpTypesMethodDelete() DaprHttpHttpTypesMethod{
  return DaprHttpHttpTypesMethod{kind: DaprHttpHttpTypesMethodKindDelete}
}

func DaprHttpHttpTypesMethodPatch() DaprHttpHttpTypesMethod{
  return DaprHttpHttpTypesMethod{kind: DaprHttpHttpTypesMethodKindPatch}
}

func DaprHttpHttpTypesMethodHead() DaprHttpHttpTypesMethod{
  return DaprHttpHttpTypesMethod{kind: DaprHttpHttpTypesMethodKindHead}
}

func DaprHttpHttpTypesMethodOptions() DaprHttpHttpTypesMethod{
  return DaprHttpHttpTypesMethod{kind: DaprHttpHttpTypesMethodKindOptions}
}

type DaprHttpHttpTypesRequest struct {
  Method DaprHttpHttpTypesMethod
  Uri string
  Headers []DaprHttpHttpTypesTuple2StringStringT
  Params []DaprHttpHttpTypesTuple2StringStringT
  Body Option[[]uint8]
}

type DaprHttpHttpTypesResponse struct {
  Status uint16
  Headers Option[[]DaprHttpHttpTypesTuple2StringStringT]
  Body Option[[]uint8]
}

type DaprHttpHttpTypesHttpErrorKind int

const (
DaprHttpHttpTypesHttpErrorKindSuccess DaprHttpHttpTypesHttpErrorKind = iota
DaprHttpHttpTypesHttpErrorKindDestinationNotAllowed
DaprHttpHttpTypesHttpErrorKindInvalidUrl
DaprHttpHttpTypesHttpErrorKindRequestError
DaprHttpHttpTypesHttpErrorKindRuntimeError
DaprHttpHttpTypesHttpErrorKindTooManyRequests
)

type DaprHttpHttpTypesHttpError struct {
  kind DaprHttpHttpTypesHttpErrorKind
}

func (n DaprHttpHttpTypesHttpError) Kind() DaprHttpHttpTypesHttpErrorKind {
  return n.kind
}

func DaprHttpHttpTypesHttpErrorSuccess() DaprHttpHttpTypesHttpError{
  return DaprHttpHttpTypesHttpError{kind: DaprHttpHttpTypesHttpErrorKindSuccess}
}

func DaprHttpHttpTypesHttpErrorDestinationNotAllowed() DaprHttpHttpTypesHttpError{
  return DaprHttpHttpTypesHttpError{kind: DaprHttpHttpTypesHttpErrorKindDestinationNotAllowed}
}

func DaprHttpHttpTypesHttpErrorInvalidUrl() DaprHttpHttpTypesHttpError{
  return DaprHttpHttpTypesHttpError{kind: DaprHttpHttpTypesHttpErrorKindInvalidUrl}
}

func DaprHttpHttpTypesHttpErrorRequestError() DaprHttpHttpTypesHttpError{
  return DaprHttpHttpTypesHttpError{kind: DaprHttpHttpTypesHttpErrorKindRequestError}
}

func DaprHttpHttpTypesHttpErrorRuntimeError() DaprHttpHttpTypesHttpError{
  return DaprHttpHttpTypesHttpError{kind: DaprHttpHttpTypesHttpErrorKindRuntimeError}
}

func DaprHttpHttpTypesHttpErrorTooManyRequests() DaprHttpHttpTypesHttpError{
  return DaprHttpHttpTypesHttpError{kind: DaprHttpHttpTypesHttpErrorKindTooManyRequests}
}

type DaprHttpHttpHandlerRequest = DaprHttpHttpTypesRequest
type DaprHttpHttpHandlerResponse = DaprHttpHttpTypesResponse
// Import functions from dapr:http/http-types
// Export functions from dapr:http/http-handler
var dapr_http_http_handler ExportsDaprHttpHttpHandler = nil
func SetExportsDaprHttpHttpHandler(i ExportsDaprHttpHttpHandler) {
  dapr_http_http_handler = i
}
type ExportsDaprHttpHttpHandler interface {
  HandleHttpRequest(req DaprHttpHttpTypesRequest) DaprHttpHttpTypesResponse 
}
//export exports_dapr_http_http_handler_handle_http_request
func ExportsDaprHttpHttpHandlerHandleHttpRequest(req *C.dapr_http_http_handler_request_t, ret *C.dapr_http_http_handler_response_t) {
  defer C.dapr_http_http_handler_request_free(req)
  var lift_req DaprHttpHttpTypesRequest
  var lift_req_val DaprHttpHttpTypesRequest
  var lift_req_val_Method DaprHttpHttpTypesMethod
  if req.method == 0 {
    lift_req_val_Method = DaprHttpHttpTypesMethodGet()
  }
  if req.method == 1 {
    lift_req_val_Method = DaprHttpHttpTypesMethodPost()
  }
  if req.method == 2 {
    lift_req_val_Method = DaprHttpHttpTypesMethodPut()
  }
  if req.method == 3 {
    lift_req_val_Method = DaprHttpHttpTypesMethodDelete()
  }
  if req.method == 4 {
    lift_req_val_Method = DaprHttpHttpTypesMethodPatch()
  }
  if req.method == 5 {
    lift_req_val_Method = DaprHttpHttpTypesMethodHead()
  }
  if req.method == 6 {
    lift_req_val_Method = DaprHttpHttpTypesMethodOptions()
  }
  lift_req_val.Method = lift_req_val_Method
  var lift_req_val_Uri string
  var lift_req_val_Uri_val string
  lift_req_val_Uri_val = C.GoStringN(req.uri.ptr, C.int(req.uri.len))
  lift_req_val_Uri = lift_req_val_Uri_val
  lift_req_val.Uri = lift_req_val_Uri
  var lift_req_val_Headers []DaprHttpHttpHandlerTuple2StringStringT
  lift_req_val_Headers = make([]DaprHttpHttpHandlerTuple2StringStringT, req.headers.len)
  if req.headers.len > 0 {
    for lift_req_val_Headers_i := 0; lift_req_val_Headers_i < int(req.headers.len); lift_req_val_Headers_i++ {
      var empty_lift_req_val_Headers C.http_tuple2_string_string_t
      lift_req_val_Headers_ptr := *(*C.http_tuple2_string_string_t)(unsafe.Pointer(uintptr(unsafe.Pointer(req.headers.ptr)) +
      uintptr(lift_req_val_Headers_i)*unsafe.Sizeof(empty_lift_req_val_Headers)))
      var list_lift_req_val_Headers DaprHttpHttpHandlerTuple2StringStringT
      var list_lift_req_val_Headers_F0 string
      list_lift_req_val_Headers_F0 = C.GoStringN(lift_req_val_Headers_ptr.f0.ptr, C.int(lift_req_val_Headers_ptr.f0.len))
      list_lift_req_val_Headers.F0 = list_lift_req_val_Headers_F0
      var list_lift_req_val_Headers_F1 string
      list_lift_req_val_Headers_F1 = C.GoStringN(lift_req_val_Headers_ptr.f1.ptr, C.int(lift_req_val_Headers_ptr.f1.len))
      list_lift_req_val_Headers.F1 = list_lift_req_val_Headers_F1
      lift_req_val_Headers[lift_req_val_Headers_i] = list_lift_req_val_Headers
    }
  }
  lift_req_val.Headers = lift_req_val_Headers
  var lift_req_val_Params []DaprHttpHttpHandlerTuple2StringStringT
  lift_req_val_Params = make([]DaprHttpHttpHandlerTuple2StringStringT, req.params.len)
  if req.params.len > 0 {
    for lift_req_val_Params_i := 0; lift_req_val_Params_i < int(req.params.len); lift_req_val_Params_i++ {
      var empty_lift_req_val_Params C.http_tuple2_string_string_t
      lift_req_val_Params_ptr := *(*C.http_tuple2_string_string_t)(unsafe.Pointer(uintptr(unsafe.Pointer(req.params.ptr)) +
      uintptr(lift_req_val_Params_i)*unsafe.Sizeof(empty_lift_req_val_Params)))
      var list_lift_req_val_Params DaprHttpHttpHandlerTuple2StringStringT
      var list_lift_req_val_Params_F0 string
      list_lift_req_val_Params_F0 = C.GoStringN(lift_req_val_Params_ptr.f0.ptr, C.int(lift_req_val_Params_ptr.f0.len))
      list_lift_req_val_Params.F0 = list_lift_req_val_Params_F0
      var list_lift_req_val_Params_F1 string
      list_lift_req_val_Params_F1 = C.GoStringN(lift_req_val_Params_ptr.f1.ptr, C.int(lift_req_val_Params_ptr.f1.len))
      list_lift_req_val_Params.F1 = list_lift_req_val_Params_F1
      lift_req_val_Params[lift_req_val_Params_i] = list_lift_req_val_Params
    }
  }
  lift_req_val.Params = lift_req_val_Params
  var lift_req_val_Body Option[[]uint8]
  if req.body.is_some {
    var lift_req_val_Body_val []uint8
    lift_req_val_Body_val = make([]uint8, req.body.val.len)
    if req.body.val.len > 0 {
      for lift_req_val_Body_val_i := 0; lift_req_val_Body_val_i < int(req.body.val.len); lift_req_val_Body_val_i++ {
        var empty_lift_req_val_Body_val C.uint8_t
        lift_req_val_Body_val_ptr := *(*C.uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(req.body.val.ptr)) +
        uintptr(lift_req_val_Body_val_i)*unsafe.Sizeof(empty_lift_req_val_Body_val)))
        var list_lift_req_val_Body_val uint8
        list_lift_req_val_Body_val = uint8(lift_req_val_Body_val_ptr)
        lift_req_val_Body_val[lift_req_val_Body_val_i] = list_lift_req_val_Body_val
      }
    }
    lift_req_val_Body.Set(lift_req_val_Body_val)
  } else {
    lift_req_val_Body.Unset()
  }
  lift_req_val.Body = lift_req_val_Body
  lift_req = lift_req_val
  result := dapr_http_http_handler.HandleHttpRequest(lift_req)
  var lower_result C.dapr_http_http_types_response_t
  var lower_result_val C.dapr_http_http_types_response_t
  var lower_result_val_status C.uint16_t
  lower_result_val_status_val := C.uint16_t(result.Status)
  lower_result_val_status = lower_result_val_status_val
  lower_result_val.status = lower_result_val_status
  var lower_result_val_headers C.http_option_dapr_http_http_types_headers_t
  if result.Headers.IsSome() {
    var lower_result_val_headers_val C.dapr_http_http_types_headers_t
    if len(result.Headers.Unwrap()) == 0 {
      lower_result_val_headers_val.ptr = nil
      lower_result_val_headers_val.len = 0
    } else {
      var empty_lower_result_val_headers_val C.http_tuple2_string_string_t
      lower_result_val_headers_val.ptr = (*C.http_tuple2_string_string_t)(C.malloc(C.size_t(len(result.Headers.Unwrap())) * C.size_t(unsafe.Sizeof(empty_lower_result_val_headers_val))))
      lower_result_val_headers_val.len = C.size_t(len(result.Headers.Unwrap()))
      for lower_result_val_headers_val_i := range result.Headers.Unwrap() {
        lower_result_val_headers_val_ptr := (*C.http_tuple2_string_string_t)(unsafe.Pointer(uintptr(unsafe.Pointer(lower_result_val_headers_val.ptr)) +
        uintptr(lower_result_val_headers_val_i)*unsafe.Sizeof(empty_lower_result_val_headers_val)))
        var lower_result_val_headers_val_ptr_value C.http_tuple2_string_string_t
        var lower_result_val_headers_val_ptr_value_f0 C.http_string_t
        
        lower_result_val_headers_val_ptr_value_f0.ptr = C.CString(result.Headers.Unwrap()[lower_result_val_headers_val_i].F0)
        lower_result_val_headers_val_ptr_value_f0.len = C.size_t(len(result.Headers.Unwrap()[lower_result_val_headers_val_i].F0))
        lower_result_val_headers_val_ptr_value.f0 = lower_result_val_headers_val_ptr_value_f0
        var lower_result_val_headers_val_ptr_value_f1 C.http_string_t
        
        lower_result_val_headers_val_ptr_value_f1.ptr = C.CString(result.Headers.Unwrap()[lower_result_val_headers_val_i].F1)
        lower_result_val_headers_val_ptr_value_f1.len = C.size_t(len(result.Headers.Unwrap()[lower_result_val_headers_val_i].F1))
        lower_result_val_headers_val_ptr_value.f1 = lower_result_val_headers_val_ptr_value_f1
        *lower_result_val_headers_val_ptr = lower_result_val_headers_val_ptr_value
      }
    }
    lower_result_val_headers.val = lower_result_val_headers_val
    lower_result_val_headers.is_some = true
  }
  lower_result_val.headers = lower_result_val_headers
  var lower_result_val_body C.http_option_dapr_http_http_types_body_t
  if result.Body.IsSome() {
    var lower_result_val_body_val C.dapr_http_http_types_body_t
    if len(result.Body.Unwrap()) == 0 {
      lower_result_val_body_val.ptr = nil
      lower_result_val_body_val.len = 0
    } else {
      var empty_lower_result_val_body_val C.uint8_t
      lower_result_val_body_val.ptr = (*C.uint8_t)(C.malloc(C.size_t(len(result.Body.Unwrap())) * C.size_t(unsafe.Sizeof(empty_lower_result_val_body_val))))
      lower_result_val_body_val.len = C.size_t(len(result.Body.Unwrap()))
      for lower_result_val_body_val_i := range result.Body.Unwrap() {
        lower_result_val_body_val_ptr := (*C.uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(lower_result_val_body_val.ptr)) +
        uintptr(lower_result_val_body_val_i)*unsafe.Sizeof(empty_lower_result_val_body_val)))
        lower_result_val_body_val_ptr_value := C.uint8_t(result.Body.Unwrap()[lower_result_val_body_val_i])
        *lower_result_val_body_val_ptr = lower_result_val_body_val_ptr_value
      }
    }
    lower_result_val_body.val = lower_result_val_body_val
    lower_result_val_body.is_some = true
  }
  lower_result_val.body = lower_result_val_body
  lower_result = lower_result_val
  *ret = lower_result
  
}