// Generated by `wit-bindgen` 0.7.0. DO NOT EDIT!
package state

// #include "state.h"
import "C"

import "unsafe"

type DaprStateStateTypesError = string
type DaprStateStateTypesKey = string
type DaprStateStateTypesTuple2StringStringT = DaprStateStateInterfaceTuple2StringStringT
type DaprStateStateInterfaceTuple2StringStringT struct {
  F0 string
  F1 string
}

type DaprStateStateTypesMetadata = DaprStateStateTypesTuple2StringStringT
type DaprStateStateTypesConsistency = string
type DaprStateStateTypesGetStateOptions struct {
  Consistency string
}

type DaprStateStateTypesGetRequest struct {
  Key string
  Metadata Option[[]DaprStateStateTypesTuple2StringStringT]
  Options Option[DaprStateStateTypesGetStateOptions]
}

type DaprStateStateTypesData = uint8
type DaprStateStateTypesGetResponse struct {
  Data []uint8
  Etag Option[string]
  Metadata Option[[]DaprStateStateTypesTuple2StringStringT]
  ContentType Option[string]
}

type DaprStateStateTypesSetStateOptions struct {
  Concurrency string
  Consistency string
}

type DaprStateStateTypesSetRequest struct {
  Key string
  Value []uint8
  Etag Option[string]
  Metadata Option[[]DaprStateStateTypesTuple2StringStringT]
  Options DaprStateStateTypesSetStateOptions
  ContentType Option[string]
}

type DaprStateStateTypesDeleteRequest struct {
  Key string
  Etag Option[string]
  Metadata Option[[]DaprStateStateTypesTuple2StringStringT]
  Options DaprStateStateTypesSetStateOptions
}

type DaprStateStateInterfaceGetRequest = DaprStateStateTypesGetRequest
type DaprStateStateInterfaceGetResponse = DaprStateStateTypesGetResponse
type DaprStateStateInterfaceSetRequest = DaprStateStateTypesSetRequest
type DaprStateStateInterfaceDeleteRequest = DaprStateStateTypesDeleteRequest
type DaprStateStateInterfaceError = string
// Import functions from dapr:state/state-types
// Export functions from dapr:state/state-interface
var dapr_state_state_interface ExportsDaprStateStateInterface = nil
func SetExportsDaprStateStateInterface(i ExportsDaprStateStateInterface) {
  dapr_state_state_interface = i
}
type ExportsDaprStateStateInterface interface {
  Get(req DaprStateStateTypesGetRequest) Result[DaprStateStateTypesGetResponse, string] 
  Set(req DaprStateStateTypesSetRequest) Result[uint32, string] 
  Delete(req DaprStateStateTypesDeleteRequest) Result[uint32, string] 
}
//export exports_dapr_state_state_interface_get
func ExportsDaprStateStateInterfaceGet(req *C.dapr_state_state_interface_get_request_t, ret *C.state_result_dapr_state_state_interface_get_response_dapr_state_state_interface_error_t) {
  defer C.dapr_state_state_interface_get_request_free(req)
  var lift_req DaprStateStateTypesGetRequest
  var lift_req_val DaprStateStateTypesGetRequest
  var lift_req_val_Key string
  var lift_req_val_Key_val string
  lift_req_val_Key_val = C.GoStringN(req.key.ptr, C.int(req.key.len))
  lift_req_val_Key = lift_req_val_Key_val
  lift_req_val.Key = lift_req_val_Key
  var lift_req_val_Metadata Option[[]DaprStateStateInterfaceTuple2StringStringT]
  if req.metadata.is_some {
    var lift_req_val_Metadata_val []DaprStateStateInterfaceTuple2StringStringT
    lift_req_val_Metadata_val = make([]DaprStateStateInterfaceTuple2StringStringT, req.metadata.val.len)
    if req.metadata.val.len > 0 {
      for lift_req_val_Metadata_val_i := 0; lift_req_val_Metadata_val_i < int(req.metadata.val.len); lift_req_val_Metadata_val_i++ {
        var empty_lift_req_val_Metadata_val C.state_tuple2_string_string_t
        lift_req_val_Metadata_val_ptr := *(*C.state_tuple2_string_string_t)(unsafe.Pointer(uintptr(unsafe.Pointer(req.metadata.val.ptr)) +
        uintptr(lift_req_val_Metadata_val_i)*unsafe.Sizeof(empty_lift_req_val_Metadata_val)))
        var list_lift_req_val_Metadata_val DaprStateStateInterfaceTuple2StringStringT
        var list_lift_req_val_Metadata_val_F0 string
        list_lift_req_val_Metadata_val_F0 = C.GoStringN(lift_req_val_Metadata_val_ptr.f0.ptr, C.int(lift_req_val_Metadata_val_ptr.f0.len))
        list_lift_req_val_Metadata_val.F0 = list_lift_req_val_Metadata_val_F0
        var list_lift_req_val_Metadata_val_F1 string
        list_lift_req_val_Metadata_val_F1 = C.GoStringN(lift_req_val_Metadata_val_ptr.f1.ptr, C.int(lift_req_val_Metadata_val_ptr.f1.len))
        list_lift_req_val_Metadata_val.F1 = list_lift_req_val_Metadata_val_F1
        lift_req_val_Metadata_val[lift_req_val_Metadata_val_i] = list_lift_req_val_Metadata_val
      }
    }
    lift_req_val_Metadata.Set(lift_req_val_Metadata_val)
  } else {
    lift_req_val_Metadata.Unset()
  }
  lift_req_val.Metadata = lift_req_val_Metadata
  var lift_req_val_Options Option[DaprStateStateTypesGetStateOptions]
  if req.options.is_some {
    var lift_req_val_Options_val DaprStateStateTypesGetStateOptions
    var lift_req_val_Options_val_Consistency string
    var lift_req_val_Options_val_Consistency_val string
    lift_req_val_Options_val_Consistency_val = C.GoStringN(req.options.val.consistency.ptr, C.int(req.options.val.consistency.len))
    lift_req_val_Options_val_Consistency = lift_req_val_Options_val_Consistency_val
    lift_req_val_Options_val.Consistency = lift_req_val_Options_val_Consistency
    lift_req_val_Options.Set(lift_req_val_Options_val)
  } else {
    lift_req_val_Options.Unset()
  }
  lift_req_val.Options = lift_req_val_Options
  lift_req = lift_req_val
  result := dapr_state_state_interface.Get(lift_req)
  var lower_result C.state_result_dapr_state_state_interface_get_response_dapr_state_state_interface_error_t
  lower_result.is_err = result.IsErr()
  if result.IsOk() {
    lower_result_ptr := (*C.dapr_state_state_interface_get_response_t)(unsafe.Pointer(&lower_result.val))
    var lower_result_val C.dapr_state_state_types_get_response_t
    var lower_result_val_val C.dapr_state_state_types_get_response_t
    var lower_result_val_val_data C.dapr_state_state_types_data_t
    if len(result.Unwrap().Data) == 0 {
      lower_result_val_val_data.ptr = nil
      lower_result_val_val_data.len = 0
    } else {
      var empty_lower_result_val_val_data C.uint8_t
      lower_result_val_val_data.ptr = (*C.uint8_t)(C.malloc(C.size_t(len(result.Unwrap().Data)) * C.size_t(unsafe.Sizeof(empty_lower_result_val_val_data))))
      lower_result_val_val_data.len = C.size_t(len(result.Unwrap().Data))
      for lower_result_val_val_data_i := range result.Unwrap().Data {
        lower_result_val_val_data_ptr := (*C.uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(lower_result_val_val_data.ptr)) +
        uintptr(lower_result_val_val_data_i)*unsafe.Sizeof(empty_lower_result_val_val_data)))
        lower_result_val_val_data_ptr_value := C.uint8_t(result.Unwrap().Data[lower_result_val_val_data_i])
        *lower_result_val_val_data_ptr = lower_result_val_val_data_ptr_value
      }
    }
    lower_result_val_val.data = lower_result_val_val_data
    var lower_result_val_val_etag C.dapr_state_state_types_etag_t
    if result.Unwrap().Etag.IsSome() {
      var lower_result_val_val_etag_val C.state_string_t
      
      lower_result_val_val_etag_val.ptr = C.CString(result.Unwrap().Etag.Unwrap())
      lower_result_val_val_etag_val.len = C.size_t(len(result.Unwrap().Etag.Unwrap()))
      lower_result_val_val_etag.val = lower_result_val_val_etag_val
      lower_result_val_val_etag.is_some = true
    }
    lower_result_val_val.etag = lower_result_val_val_etag
    var lower_result_val_val_metadata C.state_option_dapr_state_state_types_metadata_t
    if result.Unwrap().Metadata.IsSome() {
      var lower_result_val_val_metadata_val C.dapr_state_state_types_metadata_t
      if len(result.Unwrap().Metadata.Unwrap()) == 0 {
        lower_result_val_val_metadata_val.ptr = nil
        lower_result_val_val_metadata_val.len = 0
      } else {
        var empty_lower_result_val_val_metadata_val C.state_tuple2_string_string_t
        lower_result_val_val_metadata_val.ptr = (*C.state_tuple2_string_string_t)(C.malloc(C.size_t(len(result.Unwrap().Metadata.Unwrap())) * C.size_t(unsafe.Sizeof(empty_lower_result_val_val_metadata_val))))
        lower_result_val_val_metadata_val.len = C.size_t(len(result.Unwrap().Metadata.Unwrap()))
        for lower_result_val_val_metadata_val_i := range result.Unwrap().Metadata.Unwrap() {
          lower_result_val_val_metadata_val_ptr := (*C.state_tuple2_string_string_t)(unsafe.Pointer(uintptr(unsafe.Pointer(lower_result_val_val_metadata_val.ptr)) +
          uintptr(lower_result_val_val_metadata_val_i)*unsafe.Sizeof(empty_lower_result_val_val_metadata_val)))
          var lower_result_val_val_metadata_val_ptr_value C.state_tuple2_string_string_t
          var lower_result_val_val_metadata_val_ptr_value_f0 C.state_string_t
          
          lower_result_val_val_metadata_val_ptr_value_f0.ptr = C.CString(result.Unwrap().Metadata.Unwrap()[lower_result_val_val_metadata_val_i].F0)
          lower_result_val_val_metadata_val_ptr_value_f0.len = C.size_t(len(result.Unwrap().Metadata.Unwrap()[lower_result_val_val_metadata_val_i].F0))
          lower_result_val_val_metadata_val_ptr_value.f0 = lower_result_val_val_metadata_val_ptr_value_f0
          var lower_result_val_val_metadata_val_ptr_value_f1 C.state_string_t
          
          lower_result_val_val_metadata_val_ptr_value_f1.ptr = C.CString(result.Unwrap().Metadata.Unwrap()[lower_result_val_val_metadata_val_i].F1)
          lower_result_val_val_metadata_val_ptr_value_f1.len = C.size_t(len(result.Unwrap().Metadata.Unwrap()[lower_result_val_val_metadata_val_i].F1))
          lower_result_val_val_metadata_val_ptr_value.f1 = lower_result_val_val_metadata_val_ptr_value_f1
          *lower_result_val_val_metadata_val_ptr = lower_result_val_val_metadata_val_ptr_value
        }
      }
      lower_result_val_val_metadata.val = lower_result_val_val_metadata_val
      lower_result_val_val_metadata.is_some = true
    }
    lower_result_val_val.metadata = lower_result_val_val_metadata
    var lower_result_val_val_content_type C.dapr_state_state_types_content_type_t
    if result.Unwrap().ContentType.IsSome() {
      var lower_result_val_val_content_type_val C.state_string_t
      
      lower_result_val_val_content_type_val.ptr = C.CString(result.Unwrap().ContentType.Unwrap())
      lower_result_val_val_content_type_val.len = C.size_t(len(result.Unwrap().ContentType.Unwrap()))
      lower_result_val_val_content_type.val = lower_result_val_val_content_type_val
      lower_result_val_val_content_type.is_some = true
    }
    lower_result_val_val.content_type = lower_result_val_val_content_type
    lower_result_val = lower_result_val_val
    *lower_result_ptr = lower_result_val
  } else {
    lower_result_ptr := (*C.dapr_state_state_interface_error_t)(unsafe.Pointer(&lower_result.val))
    var lower_result_val C.dapr_state_state_types_error_t
    var lower_result_val_val C.state_string_t
    var lower_result_val_val_val C.state_string_t
    
    lower_result_val_val_val.ptr = C.CString(result.UnwrapErr())
    lower_result_val_val_val.len = C.size_t(len(result.UnwrapErr()))
    lower_result_val_val = lower_result_val_val_val
    lower_result_val = lower_result_val_val
    *lower_result_ptr = lower_result_val
  }
  *ret = lower_result
  
}
//export exports_dapr_state_state_interface_set
func ExportsDaprStateStateInterfaceSet(req *C.dapr_state_state_interface_set_request_t, ret *C.state_result_u32_dapr_state_state_interface_error_t) {
  defer C.dapr_state_state_interface_set_request_free(req)
  var lift_req DaprStateStateTypesSetRequest
  var lift_req_val DaprStateStateTypesSetRequest
  var lift_req_val_Key string
  var lift_req_val_Key_val string
  lift_req_val_Key_val = C.GoStringN(req.key.ptr, C.int(req.key.len))
  lift_req_val_Key = lift_req_val_Key_val
  lift_req_val.Key = lift_req_val_Key
  var lift_req_val_Value []uint8
  lift_req_val_Value = make([]uint8, req.value.len)
  if req.value.len > 0 {
    for lift_req_val_Value_i := 0; lift_req_val_Value_i < int(req.value.len); lift_req_val_Value_i++ {
      var empty_lift_req_val_Value C.uint8_t
      lift_req_val_Value_ptr := *(*C.uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(req.value.ptr)) +
      uintptr(lift_req_val_Value_i)*unsafe.Sizeof(empty_lift_req_val_Value)))
      var list_lift_req_val_Value uint8
      list_lift_req_val_Value = uint8(lift_req_val_Value_ptr)
      lift_req_val_Value[lift_req_val_Value_i] = list_lift_req_val_Value
    }
  }
  lift_req_val.Value = lift_req_val_Value
  var lift_req_val_Etag Option[string]
  if req.etag.is_some {
    var lift_req_val_Etag_val string
    lift_req_val_Etag_val = C.GoStringN(req.etag.val.ptr, C.int(req.etag.val.len))
    lift_req_val_Etag.Set(lift_req_val_Etag_val)
  } else {
    lift_req_val_Etag.Unset()
  }
  lift_req_val.Etag = lift_req_val_Etag
  var lift_req_val_Metadata Option[[]DaprStateStateInterfaceTuple2StringStringT]
  if req.metadata.is_some {
    var lift_req_val_Metadata_val []DaprStateStateInterfaceTuple2StringStringT
    lift_req_val_Metadata_val = make([]DaprStateStateInterfaceTuple2StringStringT, req.metadata.val.len)
    if req.metadata.val.len > 0 {
      for lift_req_val_Metadata_val_i := 0; lift_req_val_Metadata_val_i < int(req.metadata.val.len); lift_req_val_Metadata_val_i++ {
        var empty_lift_req_val_Metadata_val C.state_tuple2_string_string_t
        lift_req_val_Metadata_val_ptr := *(*C.state_tuple2_string_string_t)(unsafe.Pointer(uintptr(unsafe.Pointer(req.metadata.val.ptr)) +
        uintptr(lift_req_val_Metadata_val_i)*unsafe.Sizeof(empty_lift_req_val_Metadata_val)))
        var list_lift_req_val_Metadata_val DaprStateStateInterfaceTuple2StringStringT
        var list_lift_req_val_Metadata_val_F0 string
        list_lift_req_val_Metadata_val_F0 = C.GoStringN(lift_req_val_Metadata_val_ptr.f0.ptr, C.int(lift_req_val_Metadata_val_ptr.f0.len))
        list_lift_req_val_Metadata_val.F0 = list_lift_req_val_Metadata_val_F0
        var list_lift_req_val_Metadata_val_F1 string
        list_lift_req_val_Metadata_val_F1 = C.GoStringN(lift_req_val_Metadata_val_ptr.f1.ptr, C.int(lift_req_val_Metadata_val_ptr.f1.len))
        list_lift_req_val_Metadata_val.F1 = list_lift_req_val_Metadata_val_F1
        lift_req_val_Metadata_val[lift_req_val_Metadata_val_i] = list_lift_req_val_Metadata_val
      }
    }
    lift_req_val_Metadata.Set(lift_req_val_Metadata_val)
  } else {
    lift_req_val_Metadata.Unset()
  }
  lift_req_val.Metadata = lift_req_val_Metadata
  var lift_req_val_Options DaprStateStateTypesSetStateOptions
  var lift_req_val_Options_Concurrency string
  lift_req_val_Options_Concurrency = C.GoStringN(req.options.concurrency.ptr, C.int(req.options.concurrency.len))
  lift_req_val_Options.Concurrency = lift_req_val_Options_Concurrency
  var lift_req_val_Options_Consistency string
  var lift_req_val_Options_Consistency_val string
  lift_req_val_Options_Consistency_val = C.GoStringN(req.options.consistency.ptr, C.int(req.options.consistency.len))
  lift_req_val_Options_Consistency = lift_req_val_Options_Consistency_val
  lift_req_val_Options.Consistency = lift_req_val_Options_Consistency
  lift_req_val.Options = lift_req_val_Options
  var lift_req_val_ContentType Option[string]
  if req.content_type.is_some {
    var lift_req_val_ContentType_val string
    lift_req_val_ContentType_val = C.GoStringN(req.content_type.val.ptr, C.int(req.content_type.val.len))
    lift_req_val_ContentType.Set(lift_req_val_ContentType_val)
  } else {
    lift_req_val_ContentType.Unset()
  }
  lift_req_val.ContentType = lift_req_val_ContentType
  lift_req = lift_req_val
  result := dapr_state_state_interface.Set(lift_req)
  var lower_result C.state_result_u32_dapr_state_state_interface_error_t
  lower_result.is_err = result.IsErr()
  if result.IsOk() {
    lower_result_ptr := (*C.uint32_t)(unsafe.Pointer(&lower_result.val))
    lower_result_val := C.uint32_t(result.Unwrap())
    *lower_result_ptr = lower_result_val
  } else {
    lower_result_ptr := (*C.dapr_state_state_interface_error_t)(unsafe.Pointer(&lower_result.val))
    var lower_result_val C.dapr_state_state_types_error_t
    var lower_result_val_val C.state_string_t
    var lower_result_val_val_val C.state_string_t
    
    lower_result_val_val_val.ptr = C.CString(result.UnwrapErr())
    lower_result_val_val_val.len = C.size_t(len(result.UnwrapErr()))
    lower_result_val_val = lower_result_val_val_val
    lower_result_val = lower_result_val_val
    *lower_result_ptr = lower_result_val
  }
  *ret = lower_result
  
}
//export exports_dapr_state_state_interface_delete
func ExportsDaprStateStateInterfaceDelete(req *C.dapr_state_state_interface_delete_request_t, ret *C.state_result_u32_dapr_state_state_interface_error_t) {
  defer C.dapr_state_state_interface_delete_request_free(req)
  var lift_req DaprStateStateTypesDeleteRequest
  var lift_req_val DaprStateStateTypesDeleteRequest
  var lift_req_val_Key string
  var lift_req_val_Key_val string
  lift_req_val_Key_val = C.GoStringN(req.key.ptr, C.int(req.key.len))
  lift_req_val_Key = lift_req_val_Key_val
  lift_req_val.Key = lift_req_val_Key
  var lift_req_val_Etag Option[string]
  if req.etag.is_some {
    var lift_req_val_Etag_val string
    lift_req_val_Etag_val = C.GoStringN(req.etag.val.ptr, C.int(req.etag.val.len))
    lift_req_val_Etag.Set(lift_req_val_Etag_val)
  } else {
    lift_req_val_Etag.Unset()
  }
  lift_req_val.Etag = lift_req_val_Etag
  var lift_req_val_Metadata Option[[]DaprStateStateInterfaceTuple2StringStringT]
  if req.metadata.is_some {
    var lift_req_val_Metadata_val []DaprStateStateInterfaceTuple2StringStringT
    lift_req_val_Metadata_val = make([]DaprStateStateInterfaceTuple2StringStringT, req.metadata.val.len)
    if req.metadata.val.len > 0 {
      for lift_req_val_Metadata_val_i := 0; lift_req_val_Metadata_val_i < int(req.metadata.val.len); lift_req_val_Metadata_val_i++ {
        var empty_lift_req_val_Metadata_val C.state_tuple2_string_string_t
        lift_req_val_Metadata_val_ptr := *(*C.state_tuple2_string_string_t)(unsafe.Pointer(uintptr(unsafe.Pointer(req.metadata.val.ptr)) +
        uintptr(lift_req_val_Metadata_val_i)*unsafe.Sizeof(empty_lift_req_val_Metadata_val)))
        var list_lift_req_val_Metadata_val DaprStateStateInterfaceTuple2StringStringT
        var list_lift_req_val_Metadata_val_F0 string
        list_lift_req_val_Metadata_val_F0 = C.GoStringN(lift_req_val_Metadata_val_ptr.f0.ptr, C.int(lift_req_val_Metadata_val_ptr.f0.len))
        list_lift_req_val_Metadata_val.F0 = list_lift_req_val_Metadata_val_F0
        var list_lift_req_val_Metadata_val_F1 string
        list_lift_req_val_Metadata_val_F1 = C.GoStringN(lift_req_val_Metadata_val_ptr.f1.ptr, C.int(lift_req_val_Metadata_val_ptr.f1.len))
        list_lift_req_val_Metadata_val.F1 = list_lift_req_val_Metadata_val_F1
        lift_req_val_Metadata_val[lift_req_val_Metadata_val_i] = list_lift_req_val_Metadata_val
      }
    }
    lift_req_val_Metadata.Set(lift_req_val_Metadata_val)
  } else {
    lift_req_val_Metadata.Unset()
  }
  lift_req_val.Metadata = lift_req_val_Metadata
  var lift_req_val_Options DaprStateStateTypesSetStateOptions
  var lift_req_val_Options_Concurrency string
  lift_req_val_Options_Concurrency = C.GoStringN(req.options.concurrency.ptr, C.int(req.options.concurrency.len))
  lift_req_val_Options.Concurrency = lift_req_val_Options_Concurrency
  var lift_req_val_Options_Consistency string
  var lift_req_val_Options_Consistency_val string
  lift_req_val_Options_Consistency_val = C.GoStringN(req.options.consistency.ptr, C.int(req.options.consistency.len))
  lift_req_val_Options_Consistency = lift_req_val_Options_Consistency_val
  lift_req_val_Options.Consistency = lift_req_val_Options_Consistency
  lift_req_val.Options = lift_req_val_Options
  lift_req = lift_req_val
  result := dapr_state_state_interface.Delete(lift_req)
  var lower_result C.state_result_u32_dapr_state_state_interface_error_t
  lower_result.is_err = result.IsErr()
  if result.IsOk() {
    lower_result_ptr := (*C.uint32_t)(unsafe.Pointer(&lower_result.val))
    lower_result_val := C.uint32_t(result.Unwrap())
    *lower_result_ptr = lower_result_val
  } else {
    lower_result_ptr := (*C.dapr_state_state_interface_error_t)(unsafe.Pointer(&lower_result.val))
    var lower_result_val C.dapr_state_state_types_error_t
    var lower_result_val_val C.state_string_t
    var lower_result_val_val_val C.state_string_t
    
    lower_result_val_val_val.ptr = C.CString(result.UnwrapErr())
    lower_result_val_val_val.len = C.size_t(len(result.UnwrapErr()))
    lower_result_val_val = lower_result_val_val_val
    lower_result_val = lower_result_val_val
    *lower_result_ptr = lower_result_val
  }
  *ret = lower_result
  
}
