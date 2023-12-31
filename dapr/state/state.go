// Generated by `wit-bindgen` 0.7.0. DO NOT EDIT!
package state

// #include "state.h"
import "C"

import "unsafe"

type DaprStateStateTypesError = string
type DaprStateStateTypesStoreName = string
type DaprStateStateTypesKey = string
type DaprStateStateTypesTuple2StringStringT = DaprStateStateInterfaceTuple2StringStringT
type DaprStateStateInterfaceTuple2StringStringT struct {
  F0 string
  F1 string
}

type DaprStateStateTypesMetadata = DaprStateStateTypesTuple2StringStringT
type DaprStateStateTypesConsistencyKind int

const (
DaprStateStateTypesConsistencyKindUnspecified DaprStateStateTypesConsistencyKind = iota
DaprStateStateTypesConsistencyKindEventual
DaprStateStateTypesConsistencyKindStrong
)

type DaprStateStateTypesConsistency struct {
  kind DaprStateStateTypesConsistencyKind
}

func (n DaprStateStateTypesConsistency) Kind() DaprStateStateTypesConsistencyKind {
  return n.kind
}

func DaprStateStateTypesConsistencyUnspecified() DaprStateStateTypesConsistency{
  return DaprStateStateTypesConsistency{kind: DaprStateStateTypesConsistencyKindUnspecified}
}

func DaprStateStateTypesConsistencyEventual() DaprStateStateTypesConsistency{
  return DaprStateStateTypesConsistency{kind: DaprStateStateTypesConsistencyKindEventual}
}

func DaprStateStateTypesConsistencyStrong() DaprStateStateTypesConsistency{
  return DaprStateStateTypesConsistency{kind: DaprStateStateTypesConsistencyKindStrong}
}

type DaprStateStateTypesGetStateOptions struct {
  Consistency DaprStateStateTypesConsistency
}

type DaprStateStateTypesGetRequest struct {
  Key string
  Metadata Option[[]DaprStateStateTypesTuple2StringStringT]
  Options DaprStateStateTypesGetStateOptions
}

type DaprStateStateTypesData = uint8
type DaprStateStateTypesGetResponse struct {
  Data []uint8
  Etag Option[string]
  Metadata Option[[]DaprStateStateTypesTuple2StringStringT]
  ContentType Option[string]
}

type DaprStateStateTypesConcurrencyKind int

const (
DaprStateStateTypesConcurrencyKindUnspecified DaprStateStateTypesConcurrencyKind = iota
DaprStateStateTypesConcurrencyKindFirstWrite
DaprStateStateTypesConcurrencyKindLastWrite
)

type DaprStateStateTypesConcurrency struct {
  kind DaprStateStateTypesConcurrencyKind
}

func (n DaprStateStateTypesConcurrency) Kind() DaprStateStateTypesConcurrencyKind {
  return n.kind
}

func DaprStateStateTypesConcurrencyUnspecified() DaprStateStateTypesConcurrency{
  return DaprStateStateTypesConcurrency{kind: DaprStateStateTypesConcurrencyKindUnspecified}
}

func DaprStateStateTypesConcurrencyFirstWrite() DaprStateStateTypesConcurrency{
  return DaprStateStateTypesConcurrency{kind: DaprStateStateTypesConcurrencyKindFirstWrite}
}

func DaprStateStateTypesConcurrencyLastWrite() DaprStateStateTypesConcurrency{
  return DaprStateStateTypesConcurrency{kind: DaprStateStateTypesConcurrencyKindLastWrite}
}

type DaprStateStateTypesSetStateOptions struct {
  Concurrency DaprStateStateTypesConcurrency
  Consistency DaprStateStateTypesConsistency
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

type DaprStateStateInterfaceStoreName = string
type DaprStateStateInterfaceGetRequest = DaprStateStateTypesGetRequest
type DaprStateStateInterfaceGetResponse = DaprStateStateTypesGetResponse
type DaprStateStateInterfaceSetRequest = DaprStateStateTypesSetRequest
type DaprStateStateInterfaceDeleteRequest = DaprStateStateTypesDeleteRequest
type DaprStateStateInterfaceError = string
// Import functions from dapr:state/state-types
// Import functions from dapr:state/state-interface
func DaprStateStateInterfaceGet(name string, req DaprStateStateTypesGetRequest) Result[DaprStateStateTypesGetResponse, string] {
  var lower_name C.dapr_state_state_types_store_name_t
  var lower_name_val C.state_string_t
  var lower_name_val_val C.state_string_t
  
  lower_name_val_val.ptr = C.CString(name)
  lower_name_val_val.len = C.size_t(len(name))
  lower_name_val = lower_name_val_val
  lower_name = lower_name_val
  defer C.dapr_state_state_interface_store_name_free(&lower_name)
  var lower_req C.dapr_state_state_types_get_request_t
  var lower_req_val C.dapr_state_state_types_get_request_t
  var lower_req_val_key C.state_string_t
  var lower_req_val_key_val C.state_string_t
  
  lower_req_val_key_val.ptr = C.CString(req.Key)
  lower_req_val_key_val.len = C.size_t(len(req.Key))
  lower_req_val_key = lower_req_val_key_val
  lower_req_val.key = lower_req_val_key
  var lower_req_val_metadata C.state_option_dapr_state_state_types_metadata_t
  if req.Metadata.IsSome() {
    var lower_req_val_metadata_val C.dapr_state_state_types_metadata_t
    if len(req.Metadata.Unwrap()) == 0 {
      lower_req_val_metadata_val.ptr = nil
      lower_req_val_metadata_val.len = 0
    } else {
      var empty_lower_req_val_metadata_val C.state_tuple2_string_string_t
      lower_req_val_metadata_val.ptr = (*C.state_tuple2_string_string_t)(C.malloc(C.size_t(len(req.Metadata.Unwrap())) * C.size_t(unsafe.Sizeof(empty_lower_req_val_metadata_val))))
      lower_req_val_metadata_val.len = C.size_t(len(req.Metadata.Unwrap()))
      for lower_req_val_metadata_val_i := range req.Metadata.Unwrap() {
        lower_req_val_metadata_val_ptr := (*C.state_tuple2_string_string_t)(unsafe.Pointer(uintptr(unsafe.Pointer(lower_req_val_metadata_val.ptr)) +
        uintptr(lower_req_val_metadata_val_i)*unsafe.Sizeof(empty_lower_req_val_metadata_val)))
        var lower_req_val_metadata_val_ptr_value C.state_tuple2_string_string_t
        var lower_req_val_metadata_val_ptr_value_f0 C.state_string_t
        
        lower_req_val_metadata_val_ptr_value_f0.ptr = C.CString(req.Metadata.Unwrap()[lower_req_val_metadata_val_i].F0)
        lower_req_val_metadata_val_ptr_value_f0.len = C.size_t(len(req.Metadata.Unwrap()[lower_req_val_metadata_val_i].F0))
        lower_req_val_metadata_val_ptr_value.f0 = lower_req_val_metadata_val_ptr_value_f0
        var lower_req_val_metadata_val_ptr_value_f1 C.state_string_t
        
        lower_req_val_metadata_val_ptr_value_f1.ptr = C.CString(req.Metadata.Unwrap()[lower_req_val_metadata_val_i].F1)
        lower_req_val_metadata_val_ptr_value_f1.len = C.size_t(len(req.Metadata.Unwrap()[lower_req_val_metadata_val_i].F1))
        lower_req_val_metadata_val_ptr_value.f1 = lower_req_val_metadata_val_ptr_value_f1
        *lower_req_val_metadata_val_ptr = lower_req_val_metadata_val_ptr_value
      }
    }
    lower_req_val_metadata.val = lower_req_val_metadata_val
    lower_req_val_metadata.is_some = true
  }
  lower_req_val.metadata = lower_req_val_metadata
  var lower_req_val_options C.dapr_state_state_types_get_state_options_t
  var lower_req_val_options_consistency C.dapr_state_state_types_consistency_t
  if req.Options.Consistency.Kind() == DaprStateStateTypesConsistencyKindUnspecified {
    lower_req_val_options_consistency = 0
  }
  if req.Options.Consistency.Kind() == DaprStateStateTypesConsistencyKindEventual {
    lower_req_val_options_consistency = 1
  }
  if req.Options.Consistency.Kind() == DaprStateStateTypesConsistencyKindStrong {
    lower_req_val_options_consistency = 2
  }
  lower_req_val_options.consistency = lower_req_val_options_consistency
  lower_req_val.options = lower_req_val_options
  lower_req = lower_req_val
  defer C.dapr_state_state_interface_get_request_free(&lower_req)
  var ret C.state_result_dapr_state_state_interface_get_response_dapr_state_state_interface_error_t
  C.dapr_state_state_interface_get(&lower_name, &lower_req, &ret)
  var lift_ret Result[DaprStateStateTypesGetResponse, string]
  if ret.is_err {
    lift_ret_ptr := *(*C.dapr_state_state_interface_error_t)(unsafe.Pointer(&ret.val))
    var lift_ret_val string
    var lift_ret_val_val string
    var lift_ret_val_val_val string
    lift_ret_val_val_val = C.GoStringN(lift_ret_ptr.ptr, C.int(lift_ret_ptr.len))
    lift_ret_val_val = lift_ret_val_val_val
    lift_ret_val = lift_ret_val_val
    lift_ret.SetErr(lift_ret_val)
  } else {
    lift_ret_ptr := *(*C.dapr_state_state_interface_get_response_t)(unsafe.Pointer(&ret.val))
    var lift_ret_val DaprStateStateTypesGetResponse
    var lift_ret_val_val DaprStateStateTypesGetResponse
    var lift_ret_val_val_Data []uint8
    lift_ret_val_val_Data = make([]uint8, lift_ret_ptr.data.len)
    if lift_ret_ptr.data.len > 0 {
      for lift_ret_val_val_Data_i := 0; lift_ret_val_val_Data_i < int(lift_ret_ptr.data.len); lift_ret_val_val_Data_i++ {
        var empty_lift_ret_val_val_Data C.uint8_t
        lift_ret_val_val_Data_ptr := *(*C.uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(lift_ret_ptr.data.ptr)) +
        uintptr(lift_ret_val_val_Data_i)*unsafe.Sizeof(empty_lift_ret_val_val_Data)))
        var list_lift_ret_val_val_Data uint8
        list_lift_ret_val_val_Data = uint8(lift_ret_val_val_Data_ptr)
        lift_ret_val_val_Data[lift_ret_val_val_Data_i] = list_lift_ret_val_val_Data
      }
    }
    lift_ret_val_val.Data = lift_ret_val_val_Data
    var lift_ret_val_val_Etag Option[string]
    if lift_ret_ptr.etag.is_some {
      var lift_ret_val_val_Etag_val string
      lift_ret_val_val_Etag_val = C.GoStringN(lift_ret_ptr.etag.val.ptr, C.int(lift_ret_ptr.etag.val.len))
      lift_ret_val_val_Etag.Set(lift_ret_val_val_Etag_val)
    } else {
      lift_ret_val_val_Etag.Unset()
    }
    lift_ret_val_val.Etag = lift_ret_val_val_Etag
    var lift_ret_val_val_Metadata Option[[]DaprStateStateInterfaceTuple2StringStringT]
    if lift_ret_ptr.metadata.is_some {
      var lift_ret_val_val_Metadata_val []DaprStateStateInterfaceTuple2StringStringT
      lift_ret_val_val_Metadata_val = make([]DaprStateStateInterfaceTuple2StringStringT, lift_ret_ptr.metadata.val.len)
      if lift_ret_ptr.metadata.val.len > 0 {
        for lift_ret_val_val_Metadata_val_i := 0; lift_ret_val_val_Metadata_val_i < int(lift_ret_ptr.metadata.val.len); lift_ret_val_val_Metadata_val_i++ {
          var empty_lift_ret_val_val_Metadata_val C.state_tuple2_string_string_t
          lift_ret_val_val_Metadata_val_ptr := *(*C.state_tuple2_string_string_t)(unsafe.Pointer(uintptr(unsafe.Pointer(lift_ret_ptr.metadata.val.ptr)) +
          uintptr(lift_ret_val_val_Metadata_val_i)*unsafe.Sizeof(empty_lift_ret_val_val_Metadata_val)))
          var list_lift_ret_val_val_Metadata_val DaprStateStateInterfaceTuple2StringStringT
          var list_lift_ret_val_val_Metadata_val_F0 string
          list_lift_ret_val_val_Metadata_val_F0 = C.GoStringN(lift_ret_val_val_Metadata_val_ptr.f0.ptr, C.int(lift_ret_val_val_Metadata_val_ptr.f0.len))
          list_lift_ret_val_val_Metadata_val.F0 = list_lift_ret_val_val_Metadata_val_F0
          var list_lift_ret_val_val_Metadata_val_F1 string
          list_lift_ret_val_val_Metadata_val_F1 = C.GoStringN(lift_ret_val_val_Metadata_val_ptr.f1.ptr, C.int(lift_ret_val_val_Metadata_val_ptr.f1.len))
          list_lift_ret_val_val_Metadata_val.F1 = list_lift_ret_val_val_Metadata_val_F1
          lift_ret_val_val_Metadata_val[lift_ret_val_val_Metadata_val_i] = list_lift_ret_val_val_Metadata_val
        }
      }
      lift_ret_val_val_Metadata.Set(lift_ret_val_val_Metadata_val)
    } else {
      lift_ret_val_val_Metadata.Unset()
    }
    lift_ret_val_val.Metadata = lift_ret_val_val_Metadata
    var lift_ret_val_val_ContentType Option[string]
    if lift_ret_ptr.content_type.is_some {
      var lift_ret_val_val_ContentType_val string
      lift_ret_val_val_ContentType_val = C.GoStringN(lift_ret_ptr.content_type.val.ptr, C.int(lift_ret_ptr.content_type.val.len))
      lift_ret_val_val_ContentType.Set(lift_ret_val_val_ContentType_val)
    } else {
      lift_ret_val_val_ContentType.Unset()
    }
    lift_ret_val_val.ContentType = lift_ret_val_val_ContentType
    lift_ret_val = lift_ret_val_val
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func DaprStateStateInterfaceSet(name string, req DaprStateStateTypesSetRequest) Result[uint32, string] {
  var lower_name C.dapr_state_state_types_store_name_t
  var lower_name_val C.state_string_t
  var lower_name_val_val C.state_string_t
  
  lower_name_val_val.ptr = C.CString(name)
  lower_name_val_val.len = C.size_t(len(name))
  lower_name_val = lower_name_val_val
  lower_name = lower_name_val
  defer C.dapr_state_state_interface_store_name_free(&lower_name)
  var lower_req C.dapr_state_state_types_set_request_t
  var lower_req_val C.dapr_state_state_types_set_request_t
  var lower_req_val_key C.state_string_t
  var lower_req_val_key_val C.state_string_t
  
  lower_req_val_key_val.ptr = C.CString(req.Key)
  lower_req_val_key_val.len = C.size_t(len(req.Key))
  lower_req_val_key = lower_req_val_key_val
  lower_req_val.key = lower_req_val_key
  var lower_req_val_value C.dapr_state_state_types_data_t
  if len(req.Value) == 0 {
    lower_req_val_value.ptr = nil
    lower_req_val_value.len = 0
  } else {
    var empty_lower_req_val_value C.uint8_t
    lower_req_val_value.ptr = (*C.uint8_t)(C.malloc(C.size_t(len(req.Value)) * C.size_t(unsafe.Sizeof(empty_lower_req_val_value))))
    lower_req_val_value.len = C.size_t(len(req.Value))
    for lower_req_val_value_i := range req.Value {
      lower_req_val_value_ptr := (*C.uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(lower_req_val_value.ptr)) +
      uintptr(lower_req_val_value_i)*unsafe.Sizeof(empty_lower_req_val_value)))
      lower_req_val_value_ptr_value := C.uint8_t(req.Value[lower_req_val_value_i])
      *lower_req_val_value_ptr = lower_req_val_value_ptr_value
    }
  }
  lower_req_val.value = lower_req_val_value
  var lower_req_val_etag C.dapr_state_state_types_etag_t
  if req.Etag.IsSome() {
    var lower_req_val_etag_val C.state_string_t
    
    lower_req_val_etag_val.ptr = C.CString(req.Etag.Unwrap())
    lower_req_val_etag_val.len = C.size_t(len(req.Etag.Unwrap()))
    lower_req_val_etag.val = lower_req_val_etag_val
    lower_req_val_etag.is_some = true
  }
  lower_req_val.etag = lower_req_val_etag
  var lower_req_val_metadata C.state_option_dapr_state_state_types_metadata_t
  if req.Metadata.IsSome() {
    var lower_req_val_metadata_val C.dapr_state_state_types_metadata_t
    if len(req.Metadata.Unwrap()) == 0 {
      lower_req_val_metadata_val.ptr = nil
      lower_req_val_metadata_val.len = 0
    } else {
      var empty_lower_req_val_metadata_val C.state_tuple2_string_string_t
      lower_req_val_metadata_val.ptr = (*C.state_tuple2_string_string_t)(C.malloc(C.size_t(len(req.Metadata.Unwrap())) * C.size_t(unsafe.Sizeof(empty_lower_req_val_metadata_val))))
      lower_req_val_metadata_val.len = C.size_t(len(req.Metadata.Unwrap()))
      for lower_req_val_metadata_val_i := range req.Metadata.Unwrap() {
        lower_req_val_metadata_val_ptr := (*C.state_tuple2_string_string_t)(unsafe.Pointer(uintptr(unsafe.Pointer(lower_req_val_metadata_val.ptr)) +
        uintptr(lower_req_val_metadata_val_i)*unsafe.Sizeof(empty_lower_req_val_metadata_val)))
        var lower_req_val_metadata_val_ptr_value C.state_tuple2_string_string_t
        var lower_req_val_metadata_val_ptr_value_f0 C.state_string_t
        
        lower_req_val_metadata_val_ptr_value_f0.ptr = C.CString(req.Metadata.Unwrap()[lower_req_val_metadata_val_i].F0)
        lower_req_val_metadata_val_ptr_value_f0.len = C.size_t(len(req.Metadata.Unwrap()[lower_req_val_metadata_val_i].F0))
        lower_req_val_metadata_val_ptr_value.f0 = lower_req_val_metadata_val_ptr_value_f0
        var lower_req_val_metadata_val_ptr_value_f1 C.state_string_t
        
        lower_req_val_metadata_val_ptr_value_f1.ptr = C.CString(req.Metadata.Unwrap()[lower_req_val_metadata_val_i].F1)
        lower_req_val_metadata_val_ptr_value_f1.len = C.size_t(len(req.Metadata.Unwrap()[lower_req_val_metadata_val_i].F1))
        lower_req_val_metadata_val_ptr_value.f1 = lower_req_val_metadata_val_ptr_value_f1
        *lower_req_val_metadata_val_ptr = lower_req_val_metadata_val_ptr_value
      }
    }
    lower_req_val_metadata.val = lower_req_val_metadata_val
    lower_req_val_metadata.is_some = true
  }
  lower_req_val.metadata = lower_req_val_metadata
  var lower_req_val_options C.dapr_state_state_types_set_state_options_t
  var lower_req_val_options_concurrency C.dapr_state_state_types_concurrency_t
  if req.Options.Concurrency.Kind() == DaprStateStateTypesConcurrencyKindUnspecified {
    lower_req_val_options_concurrency = 0
  }
  if req.Options.Concurrency.Kind() == DaprStateStateTypesConcurrencyKindFirstWrite {
    lower_req_val_options_concurrency = 1
  }
  if req.Options.Concurrency.Kind() == DaprStateStateTypesConcurrencyKindLastWrite {
    lower_req_val_options_concurrency = 2
  }
  lower_req_val_options.concurrency = lower_req_val_options_concurrency
  var lower_req_val_options_consistency C.dapr_state_state_types_consistency_t
  if req.Options.Consistency.Kind() == DaprStateStateTypesConsistencyKindUnspecified {
    lower_req_val_options_consistency = 0
  }
  if req.Options.Consistency.Kind() == DaprStateStateTypesConsistencyKindEventual {
    lower_req_val_options_consistency = 1
  }
  if req.Options.Consistency.Kind() == DaprStateStateTypesConsistencyKindStrong {
    lower_req_val_options_consistency = 2
  }
  lower_req_val_options.consistency = lower_req_val_options_consistency
  lower_req_val.options = lower_req_val_options
  var lower_req_val_content_type C.dapr_state_state_types_content_type_t
  if req.ContentType.IsSome() {
    var lower_req_val_content_type_val C.state_string_t
    
    lower_req_val_content_type_val.ptr = C.CString(req.ContentType.Unwrap())
    lower_req_val_content_type_val.len = C.size_t(len(req.ContentType.Unwrap()))
    lower_req_val_content_type.val = lower_req_val_content_type_val
    lower_req_val_content_type.is_some = true
  }
  lower_req_val.content_type = lower_req_val_content_type
  lower_req = lower_req_val
  defer C.dapr_state_state_interface_set_request_free(&lower_req)
  var ret C.state_result_u32_dapr_state_state_interface_error_t
  C.dapr_state_state_interface_set(&lower_name, &lower_req, &ret)
  var lift_ret Result[uint32, string]
  if ret.is_err {
    lift_ret_ptr := *(*C.dapr_state_state_interface_error_t)(unsafe.Pointer(&ret.val))
    var lift_ret_val string
    var lift_ret_val_val string
    var lift_ret_val_val_val string
    lift_ret_val_val_val = C.GoStringN(lift_ret_ptr.ptr, C.int(lift_ret_ptr.len))
    lift_ret_val_val = lift_ret_val_val_val
    lift_ret_val = lift_ret_val_val
    lift_ret.SetErr(lift_ret_val)
  } else {
    lift_ret_ptr := *(*C.uint32_t)(unsafe.Pointer(&ret.val))
    var lift_ret_val uint32
    lift_ret_val = uint32(lift_ret_ptr)
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

func DaprStateStateInterfaceDelete(name string, req DaprStateStateTypesDeleteRequest) Result[uint32, string] {
  var lower_name C.dapr_state_state_types_store_name_t
  var lower_name_val C.state_string_t
  var lower_name_val_val C.state_string_t
  
  lower_name_val_val.ptr = C.CString(name)
  lower_name_val_val.len = C.size_t(len(name))
  lower_name_val = lower_name_val_val
  lower_name = lower_name_val
  defer C.dapr_state_state_interface_store_name_free(&lower_name)
  var lower_req C.dapr_state_state_types_delete_request_t
  var lower_req_val C.dapr_state_state_types_delete_request_t
  var lower_req_val_key C.state_string_t
  var lower_req_val_key_val C.state_string_t
  
  lower_req_val_key_val.ptr = C.CString(req.Key)
  lower_req_val_key_val.len = C.size_t(len(req.Key))
  lower_req_val_key = lower_req_val_key_val
  lower_req_val.key = lower_req_val_key
  var lower_req_val_etag C.dapr_state_state_types_etag_t
  if req.Etag.IsSome() {
    var lower_req_val_etag_val C.state_string_t
    
    lower_req_val_etag_val.ptr = C.CString(req.Etag.Unwrap())
    lower_req_val_etag_val.len = C.size_t(len(req.Etag.Unwrap()))
    lower_req_val_etag.val = lower_req_val_etag_val
    lower_req_val_etag.is_some = true
  }
  lower_req_val.etag = lower_req_val_etag
  var lower_req_val_metadata C.state_option_dapr_state_state_types_metadata_t
  if req.Metadata.IsSome() {
    var lower_req_val_metadata_val C.dapr_state_state_types_metadata_t
    if len(req.Metadata.Unwrap()) == 0 {
      lower_req_val_metadata_val.ptr = nil
      lower_req_val_metadata_val.len = 0
    } else {
      var empty_lower_req_val_metadata_val C.state_tuple2_string_string_t
      lower_req_val_metadata_val.ptr = (*C.state_tuple2_string_string_t)(C.malloc(C.size_t(len(req.Metadata.Unwrap())) * C.size_t(unsafe.Sizeof(empty_lower_req_val_metadata_val))))
      lower_req_val_metadata_val.len = C.size_t(len(req.Metadata.Unwrap()))
      for lower_req_val_metadata_val_i := range req.Metadata.Unwrap() {
        lower_req_val_metadata_val_ptr := (*C.state_tuple2_string_string_t)(unsafe.Pointer(uintptr(unsafe.Pointer(lower_req_val_metadata_val.ptr)) +
        uintptr(lower_req_val_metadata_val_i)*unsafe.Sizeof(empty_lower_req_val_metadata_val)))
        var lower_req_val_metadata_val_ptr_value C.state_tuple2_string_string_t
        var lower_req_val_metadata_val_ptr_value_f0 C.state_string_t
        
        lower_req_val_metadata_val_ptr_value_f0.ptr = C.CString(req.Metadata.Unwrap()[lower_req_val_metadata_val_i].F0)
        lower_req_val_metadata_val_ptr_value_f0.len = C.size_t(len(req.Metadata.Unwrap()[lower_req_val_metadata_val_i].F0))
        lower_req_val_metadata_val_ptr_value.f0 = lower_req_val_metadata_val_ptr_value_f0
        var lower_req_val_metadata_val_ptr_value_f1 C.state_string_t
        
        lower_req_val_metadata_val_ptr_value_f1.ptr = C.CString(req.Metadata.Unwrap()[lower_req_val_metadata_val_i].F1)
        lower_req_val_metadata_val_ptr_value_f1.len = C.size_t(len(req.Metadata.Unwrap()[lower_req_val_metadata_val_i].F1))
        lower_req_val_metadata_val_ptr_value.f1 = lower_req_val_metadata_val_ptr_value_f1
        *lower_req_val_metadata_val_ptr = lower_req_val_metadata_val_ptr_value
      }
    }
    lower_req_val_metadata.val = lower_req_val_metadata_val
    lower_req_val_metadata.is_some = true
  }
  lower_req_val.metadata = lower_req_val_metadata
  var lower_req_val_options C.dapr_state_state_types_set_state_options_t
  var lower_req_val_options_concurrency C.dapr_state_state_types_concurrency_t
  if req.Options.Concurrency.Kind() == DaprStateStateTypesConcurrencyKindUnspecified {
    lower_req_val_options_concurrency = 0
  }
  if req.Options.Concurrency.Kind() == DaprStateStateTypesConcurrencyKindFirstWrite {
    lower_req_val_options_concurrency = 1
  }
  if req.Options.Concurrency.Kind() == DaprStateStateTypesConcurrencyKindLastWrite {
    lower_req_val_options_concurrency = 2
  }
  lower_req_val_options.concurrency = lower_req_val_options_concurrency
  var lower_req_val_options_consistency C.dapr_state_state_types_consistency_t
  if req.Options.Consistency.Kind() == DaprStateStateTypesConsistencyKindUnspecified {
    lower_req_val_options_consistency = 0
  }
  if req.Options.Consistency.Kind() == DaprStateStateTypesConsistencyKindEventual {
    lower_req_val_options_consistency = 1
  }
  if req.Options.Consistency.Kind() == DaprStateStateTypesConsistencyKindStrong {
    lower_req_val_options_consistency = 2
  }
  lower_req_val_options.consistency = lower_req_val_options_consistency
  lower_req_val.options = lower_req_val_options
  lower_req = lower_req_val
  defer C.dapr_state_state_interface_delete_request_free(&lower_req)
  var ret C.state_result_u32_dapr_state_state_interface_error_t
  C.dapr_state_state_interface_delete(&lower_name, &lower_req, &ret)
  var lift_ret Result[uint32, string]
  if ret.is_err {
    lift_ret_ptr := *(*C.dapr_state_state_interface_error_t)(unsafe.Pointer(&ret.val))
    var lift_ret_val string
    var lift_ret_val_val string
    var lift_ret_val_val_val string
    lift_ret_val_val_val = C.GoStringN(lift_ret_ptr.ptr, C.int(lift_ret_ptr.len))
    lift_ret_val_val = lift_ret_val_val_val
    lift_ret_val = lift_ret_val_val
    lift_ret.SetErr(lift_ret_val)
  } else {
    lift_ret_ptr := *(*C.uint32_t)(unsafe.Pointer(&ret.val))
    var lift_ret_val uint32
    lift_ret_val = uint32(lift_ret_ptr)
    lift_ret.Set(lift_ret_val)
  }
  return lift_ret
}

