// Generated by `wit-bindgen` 0.9.0. DO NOT EDIT!
#include "http.h"

__attribute__((__weak__, __export_name__("cabi_post_dapr:http/http-handler#handle-http-request")))
void __wasm_export_exports_dapr_http_http_handler_handle_http_request_post_return(int32_t arg0) {
  switch ((int32_t) (int32_t) (*((uint8_t*) (arg0 + 4)))) {
    case 0: {
      break;
    }
    case 1: {
      int32_t ptr = *((int32_t*) (arg0 + 8));
      int32_t len = *((int32_t*) (arg0 + 12));
      for (int32_t i = 0; i < len; i++) {
        int32_t base = ptr + i * 16;
        (void) base;
        if ((*((int32_t*) (base + 4))) > 0) {
          free((void*) (*((int32_t*) (base + 0))));
        }
        if ((*((int32_t*) (base + 12))) > 0) {
          free((void*) (*((int32_t*) (base + 8))));
        }
      }
      if (len > 0) {
        free((void*) (ptr));
      }
      break;
    }
  }
  switch ((int32_t) (int32_t) (*((uint8_t*) (arg0 + 16)))) {
    case 0: {
      break;
    }
    case 1: {
      int32_t ptr0 = *((int32_t*) (arg0 + 20));
      int32_t len1 = *((int32_t*) (arg0 + 24));
      for (int32_t i2 = 0; i2 < len1; i2++) {
        int32_t base = ptr0 + i2 * 1;
        (void) base;
      }
      if (len1 > 0) {
        free((void*) (ptr0));
      }
      break;
    }
  }
}

__attribute__((__weak__, __export_name__("cabi_realloc")))
void *cabi_realloc(void *ptr, size_t old_size, size_t align, size_t new_size) {
  if (new_size == 0) return (void*) align;
  void *ret = realloc(ptr, new_size);
  if (!ret) abort();
  return ret;
}

// Helper Functions

void dapr_http_http_types_body_free(dapr_http_http_types_body_t *ptr) {
  if (ptr->len > 0) {
    free(ptr->ptr);
  }
}

void http_tuple2_string_string_free(http_tuple2_string_string_t *ptr) {
  http_string_free(&ptr->f0);
  http_string_free(&ptr->f1);
}

void dapr_http_http_types_headers_free(dapr_http_http_types_headers_t *ptr) {
  for (size_t i = 0; i < ptr->len; i++) {
    http_tuple2_string_string_free(&ptr->ptr[i]);
  }
  if (ptr->len > 0) {
    free(ptr->ptr);
  }
}

void dapr_http_http_types_params_free(dapr_http_http_types_params_t *ptr) {
  for (size_t i = 0; i < ptr->len; i++) {
    http_tuple2_string_string_free(&ptr->ptr[i]);
  }
  if (ptr->len > 0) {
    free(ptr->ptr);
  }
}

void dapr_http_http_types_uri_free(dapr_http_http_types_uri_t *ptr) {
  http_string_free(ptr);
}

void http_option_dapr_http_http_types_body_free(http_option_dapr_http_http_types_body_t *ptr) {
  if (ptr->is_some) {
    dapr_http_http_types_body_free(&ptr->val);
  }
}

void dapr_http_http_types_request_free(dapr_http_http_types_request_t *ptr) {
  dapr_http_http_types_uri_free(&ptr->uri);
  dapr_http_http_types_headers_free(&ptr->headers);
  dapr_http_http_types_params_free(&ptr->params);
  http_option_dapr_http_http_types_body_free(&ptr->body);
}

void http_option_dapr_http_http_types_headers_free(http_option_dapr_http_http_types_headers_t *ptr) {
  if (ptr->is_some) {
    dapr_http_http_types_headers_free(&ptr->val);
  }
}

void dapr_http_http_types_response_free(dapr_http_http_types_response_t *ptr) {
  http_option_dapr_http_http_types_headers_free(&ptr->headers);
  http_option_dapr_http_http_types_body_free(&ptr->body);
}

void dapr_http_http_handler_request_free(dapr_http_http_handler_request_t *ptr) {
  dapr_http_http_types_request_free(ptr);
}

void dapr_http_http_handler_response_free(dapr_http_http_handler_response_t *ptr) {
  dapr_http_http_types_response_free(ptr);
}

void http_string_set(http_string_t *ret, const char*s) {
  ret->ptr = (char*) s;
  ret->len = strlen(s);
}

void http_string_dup(http_string_t *ret, const char*s) {
  ret->len = strlen(s);
  ret->ptr = cabi_realloc(NULL, 0, 1, ret->len * 1);
  memcpy(ret->ptr, s, ret->len * 1);
}

void http_string_free(http_string_t *ret) {
  if (ret->len > 0) {
    free(ret->ptr);
  }
  ret->ptr = NULL;
  ret->len = 0;
}

// Component Adapters

__attribute__((__aligned__(4)))
static uint8_t RET_AREA[28];

__attribute__((__export_name__("dapr:http/http-handler#handle-http-request")))
int32_t __wasm_export_exports_dapr_http_http_handler_handle_http_request(int32_t arg, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, int32_t arg4, int32_t arg5, int32_t arg6, int32_t arg7, int32_t arg8) {
  http_option_dapr_http_http_types_body_t option;
  switch (arg6) {
    case 0: {
      option.is_some = false;
      break;
    }
    case 1: {
      option.is_some = true;
      option.val = (dapr_http_http_types_body_t) { (uint8_t*)(arg7), (size_t)(arg8) };
      break;
    }
  }
  dapr_http_http_handler_request_t arg9 = (dapr_http_http_types_request_t) {
    arg,
    (http_string_t) { (char*)(arg0), (size_t)(arg1) },
    (dapr_http_http_types_headers_t) { (http_tuple2_string_string_t*)(arg2), (size_t)(arg3) },
    (dapr_http_http_types_params_t) { (http_tuple2_string_string_t*)(arg4), (size_t)(arg5) },
    option,
  };
  dapr_http_http_handler_response_t ret;
  exports_dapr_http_http_handler_handle_http_request(&arg9, &ret);
  int32_t ptr = (int32_t) &RET_AREA;
  *((int16_t*)(ptr + 0)) = (int32_t) ((ret).status);
  if (((ret).headers).is_some) {
    const dapr_http_http_types_headers_t *payload10 = &((ret).headers).val;
    *((int8_t*)(ptr + 4)) = 1;
    *((int32_t*)(ptr + 12)) = (int32_t) (*payload10).len;
    *((int32_t*)(ptr + 8)) = (int32_t) (*payload10).ptr;
  } else {
    *((int8_t*)(ptr + 4)) = 0;
  }
  if (((ret).body).is_some) {
    const dapr_http_http_types_body_t *payload12 = &((ret).body).val;
    *((int8_t*)(ptr + 16)) = 1;
    *((int32_t*)(ptr + 24)) = (int32_t) (*payload12).len;
    *((int32_t*)(ptr + 20)) = (int32_t) (*payload12).ptr;
  } else {
    *((int8_t*)(ptr + 16)) = 0;
  }
  return ptr;
}

extern void __component_type_object_force_link_http(void);
void __component_type_object_force_link_http_public_use_in_this_compilation_unit(void) {
  __component_type_object_force_link_http();
}