(module
  (type (;0;) (func (param i32 i32 i32 i32) (result i32)))
  (type (;1;) (func (param i32)))
  (type (;2;) (func (param i32 i32)))
  (type (;3;) (func (param i32 i32) (result i32)))
  (type (;4;) (func))
  (type (;5;) (func (param i32) (result i32)))
  (type (;6;) (func (param i32 i32 i32)))
  (type (;7;) (func (result i32)))
  (type (;8;) (func (param i32 i32 i32 i32)))
  (type (;9;) (func (param i32 i32 i32 i32 i32 i32 i32 i32 i32 i32) (result i32)))
  (import "wasi_snapshot_preview1" "fd_write" (func $runtime.fd_write (type 0)))
  (func $exports_dapr_http_http_handler_handle_http_request (type 2) (param i32 i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32)
    global.get $__stack_pointer
    i32.const 624
    i32.sub
    local.tee 2
    global.set $__stack_pointer
    local.get 2
    i32.const 114
    i32.store offset=164
    local.get 2
    i32.const 168
    i32.add
    i32.const 0
    i32.const 456
    memory.fill
    local.get 2
    local.get 0
    i32.store offset=152
    local.get 2
    i64.const 0
    i64.store offset=144
    local.get 2
    i32.const 65988
    i32.load
    local.tee 13
    i32.store offset=160
    i32.const 65988
    local.get 2
    i32.const 160
    i32.add
    i32.store
    block  ;; label = @1
      local.get 0
      i32.eqz
      br_if 0 (;@1;)
      local.get 2
      local.get 0
      i32.load offset=4
      local.tee 3
      i32.store offset=168
      local.get 2
      i32.const 136
      i32.add
      local.get 3
      local.get 0
      i32.const 8
      i32.add
      i32.load
      call $github.com/taction/wit-dapr/dapr/http.C.GoStringN
      local.get 2
      local.get 2
      i32.load offset=136
      local.tee 3
      i32.store offset=324
      local.get 2
      local.get 3
      i32.store offset=308
      local.get 2
      local.get 3
      i32.store offset=172
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              local.get 0
              i32.const 16
              i32.add
              i32.load
              local.tee 9
              i32.const 268435455
              i32.gt_u
              br_if 0 (;@5;)
              local.get 2
              local.get 9
              i32.const 4
              i32.shl
              call $runtime.alloc.llvm.16733369349353020960
              local.tee 7
              i32.store offset=312
              local.get 2
              local.get 7
              i32.store offset=328
              local.get 2
              local.get 7
              i32.store offset=176
              block  ;; label = @6
                local.get 0
                i32.load offset=16
                local.tee 3
                i32.eqz
                br_if 0 (;@6;)
                loop  ;; label = @7
                  local.get 3
                  local.get 5
                  i32.le_s
                  br_if 1 (;@6;)
                  local.get 2
                  local.get 0
                  i32.load offset=12
                  local.tee 3
                  i32.store offset=184
                  local.get 2
                  local.get 3
                  i32.store offset=180
                  local.get 2
                  local.get 3
                  local.get 4
                  i32.add
                  local.tee 3
                  i32.store offset=192
                  local.get 2
                  local.get 3
                  i32.store offset=188
                  local.get 2
                  local.get 3
                  i32.load
                  local.tee 8
                  i32.store offset=204
                  local.get 2
                  local.get 8
                  i32.store offset=196
                  local.get 3
                  i32.load offset=4
                  local.set 10
                  local.get 2
                  local.get 3
                  i32.load offset=8
                  local.tee 6
                  i32.store offset=200
                  local.get 2
                  local.get 6
                  i32.store offset=212
                  local.get 3
                  i32.const 12
                  i32.add
                  i32.load
                  local.set 3
                  local.get 2
                  i32.const 128
                  i32.add
                  local.get 8
                  local.get 10
                  call $github.com/taction/wit-dapr/dapr/http.C.GoStringN
                  local.get 2
                  local.get 2
                  i32.load offset=128
                  local.tee 8
                  i32.store offset=220
                  local.get 2
                  local.get 8
                  i32.store offset=208
                  local.get 2
                  i32.load offset=132
                  local.set 10
                  local.get 2
                  i32.const 120
                  i32.add
                  local.get 6
                  local.get 3
                  call $github.com/taction/wit-dapr/dapr/http.C.GoStringN
                  local.get 2
                  local.get 2
                  i32.load offset=120
                  local.tee 6
                  i32.store offset=224
                  local.get 2
                  local.get 6
                  i32.store offset=216
                  local.get 5
                  local.get 9
                  i32.eq
                  br_if 4 (;@3;)
                  local.get 2
                  i32.load offset=124
                  local.set 11
                  local.get 4
                  local.get 7
                  i32.add
                  local.tee 3
                  local.get 8
                  i32.store
                  local.get 3
                  i32.const 12
                  i32.add
                  local.get 11
                  i32.store
                  local.get 3
                  i32.const 8
                  i32.add
                  local.get 6
                  i32.store
                  local.get 3
                  i32.const 4
                  i32.add
                  local.get 10
                  i32.store
                  local.get 4
                  i32.const 16
                  i32.add
                  local.set 4
                  local.get 5
                  i32.const 1
                  i32.add
                  local.set 5
                  local.get 0
                  i32.load offset=16
                  local.set 3
                  br 0 (;@7;)
                end
                unreachable
              end
              local.get 0
              i32.const 24
              i32.add
              i32.load
              local.tee 9
              i32.const 268435455
              i32.gt_u
              br_if 0 (;@5;)
              local.get 2
              local.get 9
              i32.const 4
              i32.shl
              call $runtime.alloc.llvm.16733369349353020960
              local.tee 7
              i32.store offset=316
              local.get 2
              local.get 7
              i32.store offset=332
              local.get 2
              local.get 7
              i32.store offset=228
              block  ;; label = @6
                local.get 0
                i32.load offset=24
                local.tee 3
                i32.eqz
                br_if 0 (;@6;)
                i32.const 0
                local.set 4
                i32.const 0
                local.set 5
                loop  ;; label = @7
                  local.get 3
                  local.get 5
                  i32.le_s
                  br_if 1 (;@6;)
                  local.get 2
                  local.get 0
                  i32.load offset=20
                  local.tee 3
                  i32.store offset=236
                  local.get 2
                  local.get 3
                  i32.store offset=232
                  local.get 2
                  local.get 3
                  local.get 4
                  i32.add
                  local.tee 3
                  i32.store offset=244
                  local.get 2
                  local.get 3
                  i32.store offset=240
                  local.get 2
                  local.get 3
                  i32.load
                  local.tee 8
                  i32.store offset=256
                  local.get 2
                  local.get 8
                  i32.store offset=248
                  local.get 3
                  i32.load offset=4
                  local.set 10
                  local.get 2
                  local.get 3
                  i32.load offset=8
                  local.tee 6
                  i32.store offset=252
                  local.get 2
                  local.get 6
                  i32.store offset=264
                  local.get 3
                  i32.const 12
                  i32.add
                  i32.load
                  local.set 3
                  local.get 2
                  i32.const 112
                  i32.add
                  local.get 8
                  local.get 10
                  call $github.com/taction/wit-dapr/dapr/http.C.GoStringN
                  local.get 2
                  local.get 2
                  i32.load offset=112
                  local.tee 8
                  i32.store offset=272
                  local.get 2
                  local.get 8
                  i32.store offset=260
                  local.get 2
                  i32.load offset=116
                  local.set 10
                  local.get 2
                  i32.const 104
                  i32.add
                  local.get 6
                  local.get 3
                  call $github.com/taction/wit-dapr/dapr/http.C.GoStringN
                  local.get 2
                  local.get 2
                  i32.load offset=104
                  local.tee 6
                  i32.store offset=276
                  local.get 2
                  local.get 6
                  i32.store offset=268
                  local.get 5
                  local.get 9
                  i32.eq
                  br_if 4 (;@3;)
                  local.get 2
                  i32.load offset=108
                  local.set 11
                  local.get 4
                  local.get 7
                  i32.add
                  local.tee 3
                  local.get 8
                  i32.store
                  local.get 3
                  i32.const 12
                  i32.add
                  local.get 11
                  i32.store
                  local.get 3
                  i32.const 8
                  i32.add
                  local.get 6
                  i32.store
                  local.get 3
                  i32.const 4
                  i32.add
                  local.get 10
                  i32.store
                  local.get 4
                  i32.const 16
                  i32.add
                  local.set 4
                  local.get 5
                  i32.const 1
                  i32.add
                  local.set 5
                  local.get 0
                  i32.load offset=24
                  local.set 3
                  br 0 (;@7;)
                end
                unreachable
              end
              i32.const 0
              local.set 10
              i32.const 0
              local.set 4
              block  ;; label = @6
                local.get 0
                i32.load8_u offset=28
                i32.eqz
                br_if 0 (;@6;)
                local.get 0
                i32.const 36
                i32.add
                i32.load
                local.tee 5
                i32.const 0
                i32.lt_s
                br_if 1 (;@5;)
                local.get 2
                local.get 5
                call $runtime.alloc.llvm.16733369349353020960
                local.tee 4
                i32.store offset=280
                local.get 2
                local.get 4
                i32.store offset=300
                local.get 0
                i32.load offset=36
                local.tee 7
                i32.eqz
                br_if 0 (;@6;)
                i32.const 0
                local.set 3
                local.get 7
                i32.const 0
                local.get 7
                i32.const 0
                i32.gt_s
                select
                local.set 9
                loop  ;; label = @7
                  local.get 3
                  local.get 9
                  i32.eq
                  br_if 1 (;@6;)
                  local.get 2
                  local.get 0
                  i32.load offset=32
                  local.tee 7
                  i32.store offset=288
                  local.get 2
                  local.get 7
                  i32.store offset=284
                  local.get 2
                  local.get 3
                  local.get 7
                  i32.add
                  local.tee 7
                  i32.store offset=296
                  local.get 2
                  local.get 7
                  i32.store offset=292
                  local.get 3
                  local.get 5
                  i32.eq
                  br_if 4 (;@3;)
                  local.get 3
                  local.get 4
                  i32.add
                  local.get 7
                  i32.load8_u
                  i32.store8
                  local.get 3
                  i32.const 1
                  i32.add
                  local.set 3
                  br 0 (;@7;)
                end
                unreachable
              end
              local.get 2
              local.get 4
              i32.store offset=320
              local.get 2
              local.get 4
              i32.store offset=336
              local.get 2
              local.get 4
              i32.store offset=304
              i32.const 16
              call $runtime.alloc.llvm.16733369349353020960
              local.tee 4
              i32.const 65737
              i32.store offset=8
              local.get 4
              i32.const 12
              i32.store offset=4
              local.get 4
              i32.const 65725
              i32.store
              local.get 4
              i32.const 12
              i32.add
              i32.const 10
              i32.store
              local.get 2
              local.get 4
              i32.store offset=484
              local.get 2
              local.get 4
              i32.store offset=500
              local.get 2
              local.get 4
              i32.store offset=468
              local.get 2
              local.get 4
              i32.store offset=452
              local.get 2
              local.get 4
              i32.store offset=428
              local.get 2
              local.get 4
              i32.store offset=420
              local.get 2
              local.get 4
              i32.store offset=404
              local.get 2
              local.get 4
              i32.store offset=396
              local.get 2
              local.get 4
              i32.store offset=392
              local.get 2
              local.get 4
              i32.store offset=384
              local.get 2
              local.get 4
              i32.store offset=376
              local.get 2
              local.get 4
              i32.store offset=372
              local.get 2
              local.get 4
              i32.store offset=368
              local.get 2
              local.get 4
              i32.store offset=340
              i32.const 16
              call $runtime.alloc.llvm.16733369349353020960
              local.tee 5
              i32.const 65747
              i64.load align=1
              i64.store align=1
              i32.const 8
              local.set 0
              local.get 5
              i32.const 8
              i32.add
              i32.const 65755
              i64.load align=1
              i64.store align=1
              local.get 2
              local.get 5
              i32.store offset=568
              local.get 2
              local.get 5
              i32.store offset=592
              local.get 2
              local.get 5
              i32.store offset=560
              local.get 2
              local.get 5
              i32.store offset=544
              local.get 2
              local.get 5
              i32.store offset=536
              local.get 2
              local.get 5
              i32.store offset=532
              local.get 2
              local.get 5
              i32.store offset=388
              local.get 2
              local.get 5
              i32.store offset=380
              local.get 2
              local.get 5
              i32.store offset=364
              local.get 2
              local.get 5
              i32.store offset=360
              local.get 2
              local.get 5
              i32.store offset=356
              local.get 2
              local.get 5
              i32.store offset=352
              local.get 2
              local.get 5
              i32.store offset=348
              local.get 2
              local.get 5
              i32.store offset=344
              local.get 2
              i32.const 96
              i32.add
              local.get 4
              i32.const 1
              call $_github.com/taction/wit-dapr/dapr/http.Option___github.com/taction/wit-dapr/dapr/http.DaprHttpHttpHandlerTuple2StringStringT__.Unwrap___github.com/taction/wit-dapr/dapr/http.DaprHttpHttpHandlerTuple2StringStringT_
              local.get 2
              local.get 2
              i32.load offset=96
              i32.store offset=400
              local.get 2
              i32.load offset=100
              i32.eqz
              if  ;; label = @6
                i32.const 0
                local.set 7
                br 2 (;@4;)
              end
              local.get 2
              i32.const 88
              i32.add
              local.get 4
              i32.const 1
              call $_github.com/taction/wit-dapr/dapr/http.Option___github.com/taction/wit-dapr/dapr/http.DaprHttpHttpHandlerTuple2StringStringT__.Unwrap___github.com/taction/wit-dapr/dapr/http.DaprHttpHttpHandlerTuple2StringStringT_
              local.get 2
              local.get 2
              i32.load offset=88
              i32.store offset=408
              local.get 2
              local.get 2
              i32.load offset=92
              i32.const 4
              i32.shl
              call $malloc
              local.tee 7
              i32.store offset=436
              local.get 2
              local.get 7
              i32.store offset=440
              local.get 2
              local.get 7
              i32.store offset=416
              local.get 2
              local.get 7
              i32.store offset=412
              local.get 2
              i32.const 80
              i32.add
              local.get 4
              i32.const 1
              call $_github.com/taction/wit-dapr/dapr/http.Option___github.com/taction/wit-dapr/dapr/http.DaprHttpHttpHandlerTuple2StringStringT__.Unwrap___github.com/taction/wit-dapr/dapr/http.DaprHttpHttpHandlerTuple2StringStringT_
              local.get 2
              local.get 2
              i32.load offset=80
              i32.store offset=424
              local.get 2
              i32.load offset=84
              local.set 10
              local.get 2
              i32.const 72
              i32.add
              local.get 4
              i32.const 1
              call $_github.com/taction/wit-dapr/dapr/http.Option___github.com/taction/wit-dapr/dapr/http.DaprHttpHttpHandlerTuple2StringStringT__.Unwrap___github.com/taction/wit-dapr/dapr/http.DaprHttpHttpHandlerTuple2StringStringT_
              local.get 2
              local.get 2
              i32.load offset=72
              i32.store offset=432
              i32.const 0
              local.set 3
              local.get 2
              i32.load offset=76
              local.tee 9
              i32.const 0
              local.get 9
              i32.const 0
              i32.gt_s
              select
              local.set 11
              loop  ;; label = @6
                local.get 3
                local.get 11
                i32.eq
                br_if 2 (;@4;)
                local.get 2
                local.get 0
                local.get 7
                i32.add
                i32.const 8
                i32.sub
                local.tee 9
                i32.store offset=444
                local.get 2
                local.get 9
                i32.store offset=448
                local.get 2
                i32.const -64
                i32.sub
                local.get 4
                i32.const 1
                call $_github.com/taction/wit-dapr/dapr/http.Option___github.com/taction/wit-dapr/dapr/http.DaprHttpHttpHandlerTuple2StringStringT__.Unwrap___github.com/taction/wit-dapr/dapr/http.DaprHttpHttpHandlerTuple2StringStringT_
                local.get 2
                local.get 2
                i32.load offset=64
                local.tee 8
                i32.store offset=456
                local.get 3
                local.get 2
                i32.load offset=68
                i32.ge_u
                br_if 3 (;@3;)
                local.get 8
                i32.eqz
                br_if 5 (;@1;)
                local.get 2
                local.get 0
                local.get 8
                i32.add
                local.tee 8
                i32.const 8
                i32.sub
                i32.load
                local.tee 6
                i32.store offset=460
                local.get 2
                local.get 6
                local.get 8
                i32.const 4
                i32.sub
                i32.load
                call $runtime.cgo_CString
                local.tee 8
                i32.store offset=480
                local.get 2
                local.get 8
                i32.store offset=516
                local.get 2
                local.get 8
                i32.store offset=464
                local.get 2
                i32.const 56
                i32.add
                local.get 4
                i32.const 1
                call $_github.com/taction/wit-dapr/dapr/http.Option___github.com/taction/wit-dapr/dapr/http.DaprHttpHttpHandlerTuple2StringStringT__.Unwrap___github.com/taction/wit-dapr/dapr/http.DaprHttpHttpHandlerTuple2StringStringT_
                local.get 2
                local.get 2
                i32.load offset=56
                local.tee 6
                i32.store offset=472
                local.get 3
                local.get 2
                i32.load offset=60
                i32.ge_u
                br_if 3 (;@3;)
                local.get 6
                i32.eqz
                br_if 5 (;@1;)
                local.get 2
                local.get 0
                local.get 6
                i32.add
                local.tee 6
                i32.const 8
                i32.sub
                i32.load
                i32.store offset=476
                local.get 6
                i32.const 4
                i32.sub
                i32.load
                local.set 14
                local.get 2
                i32.const 48
                i32.add
                local.get 4
                i32.const 1
                call $_github.com/taction/wit-dapr/dapr/http.Option___github.com/taction/wit-dapr/dapr/http.DaprHttpHttpHandlerTuple2StringStringT__.Unwrap___github.com/taction/wit-dapr/dapr/http.DaprHttpHttpHandlerTuple2StringStringT_
                local.get 2
                local.get 2
                i32.load offset=48
                local.tee 6
                i32.store offset=488
                local.get 3
                local.get 2
                i32.load offset=52
                i32.ge_u
                br_if 3 (;@3;)
                local.get 2
                local.get 0
                local.get 6
                i32.add
                local.tee 6
                i32.load
                local.tee 12
                i32.store offset=492
                local.get 2
                local.get 12
                local.get 6
                i32.const 4
                i32.add
                i32.load
                call $runtime.cgo_CString
                local.tee 6
                i32.store offset=512
                local.get 2
                local.get 6
                i32.store offset=520
                local.get 2
                local.get 6
                i32.store offset=496
                local.get 2
                i32.const 40
                i32.add
                local.get 4
                i32.const 1
                call $_github.com/taction/wit-dapr/dapr/http.Option___github.com/taction/wit-dapr/dapr/http.DaprHttpHttpHandlerTuple2StringStringT__.Unwrap___github.com/taction/wit-dapr/dapr/http.DaprHttpHttpHandlerTuple2StringStringT_
                local.get 2
                local.get 2
                i32.load offset=40
                local.tee 12
                i32.store offset=504
                local.get 3
                local.get 2
                i32.load offset=44
                i32.ge_u
                br_if 3 (;@3;)
                local.get 2
                local.get 0
                local.get 12
                i32.add
                local.tee 12
                i32.load
                i32.store offset=508
                local.get 9
                i32.const 12
                i32.add
                local.get 12
                i32.const 4
                i32.add
                i32.load
                i32.store
                local.get 9
                local.get 6
                i32.store offset=8
                local.get 9
                local.get 14
                i32.store offset=4
                local.get 9
                local.get 8
                i32.store
                local.get 0
                i32.const 16
                i32.add
                local.set 0
                local.get 3
                i32.const 1
                i32.add
                local.set 3
                br 0 (;@6;)
              end
              unreachable
            end
            call $runtime.slicePanic.llvm.16733369349353020960
            unreachable
          end
          local.get 2
          local.get 7
          i32.store offset=608
          local.get 2
          local.get 7
          i32.store offset=616
          local.get 2
          local.get 7
          i32.store offset=528
          local.get 2
          local.get 7
          i32.store offset=524
          local.get 2
          i32.const 32
          i32.add
          local.get 5
          i32.const 16
          call $_github.com/taction/wit-dapr/dapr/http.Option___github.com/taction/wit-dapr/dapr/http.DaprHttpHttpHandlerTuple2StringStringT__.Unwrap___github.com/taction/wit-dapr/dapr/http.DaprHttpHttpHandlerTuple2StringStringT_
          local.get 2
          local.get 2
          i32.load offset=32
          i32.store offset=540
          local.get 2
          i32.load offset=36
          i32.eqz
          if  ;; label = @4
            i32.const 0
            local.set 0
            i32.const 0
            local.set 4
            br 2 (;@2;)
          end
          local.get 2
          i32.const 24
          i32.add
          local.get 5
          i32.const 16
          call $_github.com/taction/wit-dapr/dapr/http.Option___github.com/taction/wit-dapr/dapr/http.DaprHttpHttpHandlerTuple2StringStringT__.Unwrap___github.com/taction/wit-dapr/dapr/http.DaprHttpHttpHandlerTuple2StringStringT_
          local.get 2
          local.get 2
          i32.load offset=24
          i32.store offset=548
          local.get 2
          local.get 2
          i32.load offset=28
          call $malloc
          local.tee 4
          i32.store offset=576
          local.get 2
          local.get 4
          i32.store offset=580
          local.get 2
          local.get 4
          i32.store offset=556
          local.get 2
          local.get 4
          i32.store offset=552
          local.get 2
          i32.const 16
          i32.add
          local.get 5
          i32.const 16
          call $_github.com/taction/wit-dapr/dapr/http.Option___github.com/taction/wit-dapr/dapr/http.DaprHttpHttpHandlerTuple2StringStringT__.Unwrap___github.com/taction/wit-dapr/dapr/http.DaprHttpHttpHandlerTuple2StringStringT_
          local.get 2
          local.get 2
          i32.load offset=16
          i32.store offset=564
          local.get 2
          i32.load offset=20
          local.set 0
          local.get 2
          i32.const 8
          i32.add
          local.get 5
          i32.const 16
          call $_github.com/taction/wit-dapr/dapr/http.Option___github.com/taction/wit-dapr/dapr/http.DaprHttpHttpHandlerTuple2StringStringT__.Unwrap___github.com/taction/wit-dapr/dapr/http.DaprHttpHttpHandlerTuple2StringStringT_
          local.get 2
          local.get 2
          i32.load offset=8
          i32.store offset=572
          i32.const 0
          local.set 3
          local.get 2
          i32.load offset=12
          local.tee 9
          i32.const 0
          local.get 9
          i32.const 0
          i32.gt_s
          select
          local.set 8
          loop  ;; label = @4
            local.get 3
            local.get 8
            i32.eq
            br_if 2 (;@2;)
            local.get 2
            local.get 3
            local.get 4
            i32.add
            local.tee 9
            i32.store offset=584
            local.get 2
            local.get 9
            i32.store offset=588
            local.get 2
            local.get 5
            i32.const 16
            call $_github.com/taction/wit-dapr/dapr/http.Option___github.com/taction/wit-dapr/dapr/http.DaprHttpHttpHandlerTuple2StringStringT__.Unwrap___github.com/taction/wit-dapr/dapr/http.DaprHttpHttpHandlerTuple2StringStringT_
            local.get 2
            local.get 2
            i32.load
            local.tee 6
            i32.store offset=596
            local.get 3
            local.get 2
            i32.load offset=4
            i32.ge_u
            br_if 1 (;@3;)
            local.get 9
            local.get 3
            local.get 6
            i32.add
            i32.load8_u
            i32.store8
            local.get 3
            i32.const 1
            i32.add
            local.set 3
            br 0 (;@4;)
          end
          unreachable
        end
        call $runtime.lookupPanic.llvm.16733369349353020960
        unreachable
      end
      local.get 2
      local.get 4
      i32.store offset=612
      local.get 2
      local.get 4
      i32.store offset=620
      local.get 2
      local.get 4
      i32.store offset=604
      local.get 2
      local.get 4
      i32.store offset=600
      local.get 1
      i32.eqz
      br_if 0 (;@1;)
      local.get 1
      local.get 0
      i32.store offset=24
      local.get 1
      local.get 4
      i32.store offset=20
      local.get 1
      i32.const 1
      i32.store8 offset=16
      local.get 1
      local.get 10
      i32.store offset=12
      local.get 1
      local.get 7
      i32.store offset=8
      local.get 1
      i32.const 1
      i32.store8 offset=4
      local.get 1
      i32.const 200
      i32.store16
      local.get 2
      i32.const 144
      i32.add
      local.set 3
      loop  ;; label = @2
        local.get 3
        if  ;; label = @3
          local.get 3
          i32.load offset=4
          local.get 3
          i32.load offset=8
          local.tee 1
          i32.const 4
          i32.add
          call $http_string_free
          local.get 1
          i32.const 12
          i32.add
          call $dapr_http_http_types_headers_free
          local.get 1
          i32.const 20
          i32.add
          call $dapr_http_http_types_headers_free
          local.get 1
          i32.const 28
          i32.add
          local.tee 1
          i32.load8_u
          if  ;; label = @4
            local.get 1
            i32.const 4
            i32.add
            local.tee 1
            i32.load offset=4
            if  ;; label = @5
              local.get 1
              i32.load
              call $free
            end
          end
          local.set 3
          br 1 (;@2;)
        end
      end
      i32.const 65988
      local.get 13
      i32.store
      local.get 2
      i32.const 624
      i32.add
      global.set $__stack_pointer
      return
    end
    call $runtime.nilPanic.llvm.16733369349353020960
    unreachable)
  (func $github.com/taction/wit-dapr/dapr/http.C.GoStringN (type 6) (param i32 i32 i32)
    (local i32 i32 i32)
    global.get $__stack_pointer
    i32.const 32
    i32.sub
    local.tee 3
    global.set $__stack_pointer
    local.get 3
    i32.const 16
    i32.add
    i64.const 0
    i64.store
    local.get 3
    i32.const 24
    i32.add
    i64.const 0
    i64.store
    local.get 3
    i64.const 0
    i64.store offset=8
    local.get 3
    i32.const 6
    i32.store offset=4
    i32.const 65988
    i32.load
    local.set 5
    i32.const 65988
    local.get 3
    i32.store
    local.get 3
    local.get 5
    i32.store
    block  ;; label = @1
      local.get 2
      if  ;; label = @2
        local.get 2
        i32.const 0
        i32.lt_s
        br_if 1 (;@1;)
        local.get 3
        local.get 2
        call $runtime.alloc.llvm.16733369349353020960
        local.tee 4
        i32.store offset=12
        local.get 3
        local.get 4
        i32.store offset=16
        local.get 3
        local.get 4
        i32.store offset=8
        local.get 4
        local.get 1
        local.get 2
        memory.copy
      end
      i32.const 65988
      local.get 5
      i32.store
      local.get 0
      local.get 2
      i32.store offset=4
      local.get 0
      local.get 4
      i32.store
      local.get 3
      i32.const 32
      i32.add
      global.set $__stack_pointer
      return
    end
    call $runtime.slicePanic.llvm.16733369349353020960
    unreachable)
  (func $runtime.alloc.llvm.16733369349353020960 (type 5) (param i32) (result i32)
    (local i32 i32 i32 i32 i32 i32 i32)
    local.get 0
    i32.eqz
    if  ;; label = @1
      i32.const 65984
      return
    end
    i32.const 65960
    i32.const 65960
    i64.load
    local.get 0
    i64.extend_i32_u
    i64.add
    i64.store
    i32.const 65968
    i32.const 65968
    i64.load
    i64.const 1
    i64.add
    i64.store
    local.get 0
    i32.const 15
    i32.add
    i32.const 4
    i32.shr_u
    local.set 5
    i32.const 65948
    i32.load
    local.tee 4
    local.set 3
    loop  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              local.get 3
              local.get 4
              i32.ne
              if  ;; label = @6
                local.get 2
                local.set 1
                br 1 (;@5;)
              end
              i32.const 1
              local.set 1
              block  ;; label = @6
                block  ;; label = @7
                  local.get 2
                  i32.const 255
                  i32.and
                  br_table 2 (;@5;) 0 (;@7;) 1 (;@6;)
                end
                i32.const 65988
                i32.load
                drop
                global.get $__stack_pointer
                i32.const 65536
                call $runtime.markRoots
                i32.const 65536
                i32.const 66304
                call $runtime.markRoots
                loop  ;; label = @7
                  i32.const 65985
                  i32.load8_u
                  i32.eqz
                  if  ;; label = @8
                    i32.const 0
                    local.set 2
                    i32.const 0
                    local.set 4
                    i32.const 0
                    local.set 1
                    loop  ;; label = @9
                      block  ;; label = @10
                        block  ;; label = @11
                          i32.const 65952
                          i32.load
                          local.get 1
                          i32.gt_u
                          if  ;; label = @12
                            block  ;; label = @13
                              block  ;; label = @14
                                block  ;; label = @15
                                  block  ;; label = @16
                                    local.get 1
                                    call $_runtime.gcBlock_.state
                                    i32.const 255
                                    i32.and
                                    br_table 3 (;@13;) 0 (;@16;) 1 (;@15;) 2 (;@14;) 6 (;@10;)
                                  end
                                  local.get 1
                                  call $_runtime.gcBlock_.markFree
                                  i32.const 65976
                                  i32.const 65976
                                  i64.load
                                  i64.const 1
                                  i64.add
                                  i64.store
                                  br 4 (;@11;)
                                end
                                local.get 4
                                i32.const 1
                                i32.and
                                i32.const 0
                                local.set 4
                                i32.eqz
                                br_if 4 (;@10;)
                                local.get 1
                                call $_runtime.gcBlock_.markFree
                                br 3 (;@11;)
                              end
                              i32.const 0
                              local.set 4
                              i32.const 65944
                              i32.load
                              local.get 1
                              i32.const 2
                              i32.shr_u
                              i32.add
                              local.tee 6
                              local.get 6
                              i32.load8_u
                              i32.const 2
                              local.get 1
                              i32.const 1
                              i32.shl
                              i32.const 6
                              i32.and
                              i32.shl
                              i32.const -1
                              i32.xor
                              i32.and
                              i32.store8
                              br 3 (;@10;)
                            end
                            local.get 2
                            i32.const 16
                            i32.add
                            local.set 2
                            br 2 (;@10;)
                          end
                          i32.const 2
                          local.set 1
                          local.get 2
                          i32.const 65944
                          i32.load
                          i32.const 66304
                          i32.sub
                          i32.const 3
                          i32.div_u
                          i32.ge_u
                          br_if 6 (;@5;)
                          call $runtime.growHeap
                          drop
                          br 6 (;@5;)
                        end
                        local.get 2
                        i32.const 16
                        i32.add
                        local.set 2
                        i32.const 1
                        local.set 4
                      end
                      local.get 1
                      i32.const 1
                      i32.add
                      local.set 1
                      br 0 (;@9;)
                    end
                    unreachable
                  end
                  i32.const 0
                  local.set 1
                  i32.const 65985
                  i32.const 0
                  i32.store8
                  i32.const 65952
                  i32.load
                  local.set 2
                  loop  ;; label = @8
                    local.get 1
                    local.get 2
                    i32.ge_u
                    br_if 1 (;@7;)
                    local.get 1
                    call $_runtime.gcBlock_.state
                    i32.const 255
                    i32.and
                    i32.const 3
                    i32.eq
                    if  ;; label = @9
                      local.get 1
                      call $runtime.startMark
                      i32.const 65952
                      i32.load
                      local.set 2
                    end
                    local.get 1
                    i32.const 1
                    i32.add
                    local.set 1
                    br 0 (;@8;)
                  end
                  unreachable
                end
                unreachable
              end
              local.get 2
              local.set 1
              call $runtime.growHeap
              i32.const 1
              i32.and
              i32.eqz
              br_if 1 (;@4;)
            end
            i32.const 65952
            i32.load
            local.get 3
            i32.eq
            if  ;; label = @5
              i32.const 0
              local.set 3
              br 2 (;@3;)
            end
            local.get 3
            call $_runtime.gcBlock_.state
            i32.const 255
            i32.and
            if  ;; label = @5
              local.get 3
              i32.const 1
              i32.add
              local.set 3
              br 2 (;@3;)
            end
            local.get 3
            i32.const 1
            i32.add
            local.set 2
            local.get 5
            local.get 7
            i32.const 1
            i32.add
            local.tee 7
            i32.ne
            if  ;; label = @5
              local.get 2
              local.set 3
              br 3 (;@2;)
            end
            i32.const 65948
            local.get 2
            i32.store
            local.get 2
            local.get 5
            i32.sub
            local.tee 2
            i32.const 1
            call $_runtime.gcBlock_.setState
            local.get 3
            local.get 5
            i32.sub
            i32.const 2
            i32.add
            local.set 1
            loop  ;; label = @5
              i32.const 65948
              i32.load
              local.get 1
              i32.ne
              if  ;; label = @6
                local.get 1
                i32.const 2
                call $_runtime.gcBlock_.setState
                local.get 1
                i32.const 1
                i32.add
                local.set 1
                br 1 (;@5;)
              end
            end
            local.get 2
            i32.const 4
            i32.shl
            i32.const 66304
            i32.add
            local.tee 2
            i32.const 0
            local.get 0
            memory.fill
            local.get 2
            return
          end
          i32.const 65624
          i32.const 13
          call $runtime.runtimePanicAt.llvm.16733369349353020960
          unreachable
        end
        i32.const 0
        local.set 7
      end
      i32.const 65948
      i32.load
      local.set 4
      local.get 1
      local.set 2
      br 0 (;@1;)
    end
    unreachable)
  (func $_github.com/taction/wit-dapr/dapr/http.Option___github.com/taction/wit-dapr/dapr/http.DaprHttpHttpHandlerTuple2StringStringT__.Unwrap___github.com/taction/wit-dapr/dapr/http.DaprHttpHttpHandlerTuple2StringStringT_ (type 6) (param i32 i32 i32)
    local.get 0
    local.get 1
    i32.store
    local.get 0
    local.get 2
    i32.store offset=4)
  (func $malloc (type 5) (param i32) (result i32)
    (local i32 i32 i32)
    global.get $__stack_pointer
    i32.const 32
    i32.sub
    local.tee 1
    global.set $__stack_pointer
    local.get 1
    i32.const 2
    i32.store offset=20
    i32.const 65988
    i32.load
    local.set 3
    i32.const 65988
    local.get 1
    i32.const 16
    i32.add
    i32.store
    local.get 1
    local.get 3
    i32.store offset=16
    block  ;; label = @1
      local.get 0
      if  ;; label = @2
        local.get 0
        i32.const 0
        i32.lt_s
        br_if 1 (;@1;)
        local.get 1
        local.get 0
        call $runtime.alloc.llvm.16733369349353020960
        local.tee 2
        i32.store offset=24
        local.get 1
        local.get 2
        i32.store offset=28
        local.get 1
        local.get 0
        i32.store offset=8
        local.get 1
        local.get 0
        i32.store offset=4
        local.get 1
        local.get 2
        i32.store
        local.get 1
        local.get 2
        i32.store offset=12
        local.get 1
        i32.const 12
        i32.add
        local.get 1
        call $runtime.hashmapBinarySet.llvm.16733369349353020960
      end
      i32.const 65988
      local.get 3
      i32.store
      local.get 1
      i32.const 32
      i32.add
      global.set $__stack_pointer
      local.get 2
      return
    end
    call $runtime.slicePanic.llvm.16733369349353020960
    unreachable)
  (func $runtime.cgo_CString (type 3) (param i32 i32) (result i32)
    (local i32 i32 i32)
    global.get $__stack_pointer
    i32.const 16
    i32.sub
    local.tee 2
    global.set $__stack_pointer
    local.get 2
    i64.const 1
    i64.store offset=4 align=4
    i32.const 65988
    i32.load
    local.set 3
    i32.const 65988
    local.get 2
    i32.store
    local.get 2
    local.get 3
    i32.store
    local.get 1
    i32.const 1
    i32.add
    call $malloc
    local.tee 4
    local.get 0
    local.get 1
    memory.copy
    local.get 1
    local.get 4
    i32.add
    i32.const 0
    i32.store8
    i32.const 65988
    local.get 3
    i32.store
    local.get 2
    i32.const 16
    i32.add
    global.set $__stack_pointer
    local.get 4)
  (func $runtime.slicePanic.llvm.16733369349353020960 (type 4)
    i32.const 65707
    i32.const 18
    call $runtime.runtimePanicAt.llvm.16733369349353020960
    unreachable)
  (func $runtime.lookupPanic.llvm.16733369349353020960 (type 4)
    i32.const 65689
    i32.const 18
    call $runtime.runtimePanicAt.llvm.16733369349353020960
    unreachable)
  (func $runtime.nilPanic.llvm.16733369349353020960 (type 4)
    i32.const 65666
    i32.const 23
    call $runtime.runtimePanicAt.llvm.16733369349353020960
    unreachable)
  (func $runtime._panic.llvm.16733369349353020960 (type 1) (param i32)
    i32.const 65637
    i32.const 7
    call $runtime.printstring.llvm.16733369349353020960
    local.get 0
    i32.load
    local.get 0
    i32.load offset=4
    call $runtime.printstring.llvm.16733369349353020960
    i32.const 10
    call $runtime.putchar.llvm.16733369349353020960
    unreachable)
  (func $runtime.memequal (type 0) (param i32 i32 i32 i32) (result i32)
    (local i32 i32)
    i32.const 0
    local.set 3
    block (result i32)  ;; label = @1
      loop  ;; label = @2
        local.get 2
        local.get 2
        local.get 3
        i32.eq
        br_if 1 (;@1;)
        drop
        local.get 1
        local.get 3
        i32.add
        local.set 4
        local.get 0
        local.get 3
        i32.add
        local.get 3
        i32.const 1
        i32.add
        local.set 3
        i32.load8_u
        local.get 4
        i32.load8_u
        i32.eq
        br_if 0 (;@2;)
      end
      local.get 3
      i32.const 1
      i32.sub
    end
    local.get 2
    i32.ge_u)
  (func $runtime.hash32.llvm.16733369349353020960 (type 0) (param i32 i32 i32 i32) (result i32)
    local.get 1
    i32.const -962287725
    i32.mul
    local.get 2
    i32.xor
    i32.const -1130422988
    i32.xor
    local.set 2
    loop  ;; label = @1
      local.get 1
      i32.const 4
      i32.lt_s
      i32.eqz
      if  ;; label = @2
        local.get 0
        i32.load align=1
        local.get 2
        i32.add
        i32.const -962287725
        i32.mul
        local.tee 2
        i32.const 16
        i32.shr_u
        local.get 2
        i32.xor
        local.set 2
        local.get 1
        i32.const 4
        i32.sub
        local.set 1
        local.get 0
        i32.const 4
        i32.add
        local.set 0
        br 1 (;@1;)
      end
    end
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            local.get 1
            i32.const 1
            i32.sub
            br_table 2 (;@2;) 1 (;@3;) 0 (;@4;) 3 (;@1;)
          end
          local.get 0
          i32.load8_u offset=2
          i32.const 16
          i32.shl
          local.get 2
          i32.add
          local.set 2
        end
        local.get 0
        i32.load8_u offset=1
        i32.const 8
        i32.shl
        local.get 2
        i32.add
        local.set 2
      end
      local.get 2
      local.get 0
      i32.load8_u
      i32.add
      i32.const -962287725
      i32.mul
      local.tee 0
      i32.const 24
      i32.shr_u
      local.get 0
      i32.xor
      local.set 2
    end
    local.get 2)
  (func $runtime.runtimePanicAt.llvm.16733369349353020960 (type 2) (param i32 i32)
    i32.const 65644
    i32.const 22
    call $runtime.printstring.llvm.16733369349353020960
    local.get 0
    local.get 1
    call $runtime.printstring.llvm.16733369349353020960
    i32.const 10
    call $runtime.putchar.llvm.16733369349353020960
    unreachable)
  (func $runtime.printstring.llvm.16733369349353020960 (type 2) (param i32 i32)
    local.get 1
    i32.const 0
    local.get 1
    i32.const 0
    i32.gt_s
    select
    local.set 1
    loop  ;; label = @1
      local.get 1
      if  ;; label = @2
        local.get 0
        i32.load8_u
        call $runtime.putchar.llvm.16733369349353020960
        local.get 1
        i32.const 1
        i32.sub
        local.set 1
        local.get 0
        i32.const 1
        i32.add
        local.set 0
        br 1 (;@1;)
      end
    end)
  (func $runtime.putchar.llvm.16733369349353020960 (type 1) (param i32)
    (local i32 i32)
    i32.const 65820
    i32.load
    local.tee 1
    i32.const 119
    i32.le_u
    if  ;; label = @1
      i32.const 65820
      local.get 1
      i32.const 1
      i32.add
      local.tee 2
      i32.store
      local.get 1
      i32.const 65824
      i32.add
      local.get 0
      i32.store8
      local.get 0
      i32.const 255
      i32.and
      i32.const 10
      i32.ne
      local.get 1
      i32.const 119
      i32.ne
      i32.and
      i32.eqz
      if  ;; label = @2
        i32.const 65772
        local.get 2
        i32.store
        i32.const 1
        i32.const 65768
        i32.const 1
        i32.const 65992
        call $runtime.fd_write
        drop
        i32.const 65820
        i32.const 0
        i32.store
      end
      return
    end
    call $runtime.lookupPanic.llvm.16733369349353020960
    unreachable)
  (func $runtime.markRoots (type 2) (param i32 i32)
    (local i32)
    loop  ;; label = @1
      local.get 0
      local.get 1
      i32.ge_u
      i32.eqz
      if  ;; label = @2
        block  ;; label = @3
          local.get 0
          i32.load
          local.tee 2
          i32.const 66304
          i32.lt_u
          br_if 0 (;@3;)
          local.get 2
          i32.const 65944
          i32.load
          i32.ge_u
          br_if 0 (;@3;)
          local.get 2
          i32.const 66304
          i32.sub
          i32.const 4
          i32.shr_u
          local.tee 2
          call $_runtime.gcBlock_.state
          i32.const 255
          i32.and
          i32.eqz
          br_if 0 (;@3;)
          local.get 2
          call $_runtime.gcBlock_.findHead
          local.tee 2
          call $_runtime.gcBlock_.state
          i32.const 255
          i32.and
          i32.const 3
          i32.eq
          br_if 0 (;@3;)
          local.get 2
          call $runtime.startMark
        end
        local.get 0
        i32.const 4
        i32.add
        local.set 0
        br 1 (;@1;)
      end
    end)
  (func $_runtime.gcBlock_.state (type 5) (param i32) (result i32)
    i32.const 65944
    i32.load
    local.get 0
    i32.const 2
    i32.shr_u
    i32.add
    i32.load8_u
    local.get 0
    i32.const 1
    i32.shl
    i32.const 6
    i32.and
    i32.shr_u
    i32.const 3
    i32.and)
  (func $_runtime.gcBlock_.markFree (type 1) (param i32)
    (local i32)
    i32.const 65944
    i32.load
    local.get 0
    i32.const 2
    i32.shr_u
    i32.add
    local.tee 1
    local.get 1
    i32.load8_u
    i32.const 3
    local.get 0
    i32.const 1
    i32.shl
    i32.const 6
    i32.and
    i32.shl
    i32.const -1
    i32.xor
    i32.and
    i32.store8)
  (func $runtime.growHeap (type 7) (result i32)
    (local i32 i32 i32)
    memory.size
    memory.grow
    i32.const -1
    i32.ne
    local.tee 1
    if  ;; label = @1
      memory.size
      local.set 0
      i32.const 65816
      i32.load
      local.set 2
      i32.const 65816
      local.get 0
      i32.const 16
      i32.shl
      i32.store
      i32.const 65944
      i32.load
      local.set 0
      call $runtime.calculateHeapAddresses
      i32.const 65944
      i32.load
      local.get 0
      local.get 2
      local.get 0
      i32.sub
      memory.copy
    end
    local.get 1)
  (func $runtime.startMark (type 1) (param i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32)
    global.get $__stack_pointer
    i32.const -64
    i32.add
    local.tee 4
    global.set $__stack_pointer
    local.get 4
    i32.const 4
    i32.add
    i32.const 0
    i32.const 60
    memory.fill
    local.get 4
    local.get 0
    i32.store
    local.get 0
    i32.const 3
    call $_runtime.gcBlock_.setState
    i32.const 1
    local.set 1
    block  ;; label = @1
      loop  ;; label = @2
        local.get 1
        i32.const 0
        i32.gt_s
        if  ;; label = @3
          local.get 1
          i32.const 1
          i32.sub
          local.tee 1
          i32.const 15
          i32.gt_u
          br_if 2 (;@1;)
          local.get 4
          local.get 1
          i32.const 2
          i32.shl
          i32.add
          i32.load
          local.tee 3
          i32.const 4
          i32.shl
          local.set 0
          block  ;; label = @4
            block  ;; label = @5
              local.get 3
              call $_runtime.gcBlock_.state
              i32.const 255
              i32.and
              i32.const 1
              i32.sub
              br_table 0 (;@5;) 1 (;@4;) 0 (;@5;) 1 (;@4;)
            end
            local.get 3
            i32.const 1
            i32.add
            local.set 3
          end
          local.get 0
          i32.const 66304
          i32.add
          local.set 6
          local.get 3
          i32.const 4
          i32.shl
          local.tee 5
          local.get 0
          i32.sub
          local.set 2
          local.get 5
          i32.const 66304
          i32.add
          local.set 5
          i32.const 65944
          i32.load
          local.set 7
          loop  ;; label = @4
            block  ;; label = @5
              local.get 2
              local.set 0
              local.get 5
              local.get 7
              i32.ge_u
              br_if 0 (;@5;)
              local.get 0
              i32.const 16
              i32.add
              local.set 2
              local.get 5
              i32.const 16
              i32.add
              local.set 5
              local.get 3
              call $_runtime.gcBlock_.state
              local.get 3
              i32.const 1
              i32.add
              local.set 3
              i32.const 255
              i32.and
              i32.const 2
              i32.eq
              br_if 1 (;@4;)
            end
          end
          loop  ;; label = @4
            local.get 0
            i32.eqz
            br_if 2 (;@2;)
            block  ;; label = @5
              local.get 6
              i32.load
              local.tee 2
              i32.const 66304
              i32.lt_u
              br_if 0 (;@5;)
              local.get 2
              i32.const 65944
              i32.load
              i32.ge_u
              br_if 0 (;@5;)
              local.get 2
              i32.const 66304
              i32.sub
              i32.const 4
              i32.shr_u
              local.tee 2
              call $_runtime.gcBlock_.state
              i32.const 255
              i32.and
              i32.eqz
              br_if 0 (;@5;)
              local.get 2
              call $_runtime.gcBlock_.findHead
              local.tee 2
              call $_runtime.gcBlock_.state
              i32.const 255
              i32.and
              i32.const 3
              i32.eq
              br_if 0 (;@5;)
              local.get 2
              i32.const 3
              call $_runtime.gcBlock_.setState
              local.get 1
              i32.const 16
              i32.eq
              if  ;; label = @6
                i32.const 65985
                i32.const 1
                i32.store8
                i32.const 16
                local.set 1
                br 1 (;@5;)
              end
              local.get 1
              i32.const 15
              i32.gt_u
              br_if 4 (;@1;)
              local.get 4
              local.get 1
              i32.const 2
              i32.shl
              i32.add
              local.get 2
              i32.store
              local.get 1
              i32.const 1
              i32.add
              local.set 1
            end
            local.get 0
            i32.const 4
            i32.sub
            local.set 0
            local.get 6
            i32.const 4
            i32.add
            local.set 6
            br 0 (;@4;)
          end
          unreachable
        end
      end
      local.get 4
      i32.const -64
      i32.sub
      global.set $__stack_pointer
      return
    end
    call $runtime.lookupPanic.llvm.16733369349353020960
    unreachable)
  (func $_runtime.gcBlock_.setState (type 2) (param i32 i32)
    (local i32)
    i32.const 65944
    i32.load
    local.get 0
    i32.const 2
    i32.shr_u
    i32.add
    local.tee 2
    local.get 2
    i32.load8_u
    local.get 1
    local.get 0
    i32.const 1
    i32.shl
    i32.const 6
    i32.and
    i32.shl
    i32.or
    i32.store8)
  (func $runtime.calculateHeapAddresses (type 4)
    (local i32)
    i32.const 65944
    i32.const 65816
    i32.load
    local.tee 0
    local.get 0
    i32.const 66240
    i32.sub
    i32.const 65
    i32.div_u
    i32.sub
    local.tee 0
    i32.store
    i32.const 65952
    local.get 0
    i32.const 66304
    i32.sub
    i32.const 4
    i32.shr_u
    i32.store)
  (func $_runtime.gcBlock_.findHead (type 5) (param i32) (result i32)
    (local i32)
    loop  ;; label = @1
      local.get 0
      call $_runtime.gcBlock_.state
      local.get 0
      i32.const 1
      i32.sub
      local.set 0
      i32.const 255
      i32.and
      i32.const 2
      i32.eq
      br_if 0 (;@1;)
    end
    local.get 0
    i32.const 1
    i32.add)
  (func $runtime.hashmapBinarySet.llvm.16733369349353020960 (type 2) (param i32 i32)
    i32.const 65776
    local.get 0
    local.get 1
    local.get 0
    i32.const 65788
    i32.load
    i32.const 65780
    i32.load
    local.get 0
    call $runtime.hash32.llvm.16733369349353020960
    call $runtime.hashmapSet.llvm.16733369349353020960)
  (func $free (type 1) (param i32)
    (local i32)
    global.get $__stack_pointer
    i32.const 16
    i32.sub
    local.tee 1
    global.set $__stack_pointer
    block  ;; label = @1
      local.get 0
      if  ;; label = @2
        local.get 1
        local.get 0
        i32.store offset=12
        local.get 1
        i32.const 12
        i32.add
        local.get 1
        call $runtime.hashmapBinaryGet.llvm.16733369349353020960
        i32.const 1
        i32.and
        i32.eqz
        br_if 1 (;@1;)
        local.get 1
        local.get 0
        i32.store
        local.get 1
        call $runtime.hashmapBinaryDelete.llvm.16733369349353020960
      end
      local.get 1
      i32.const 16
      i32.add
      global.set $__stack_pointer
      return
    end
    i32.const 65584
    call $runtime._panic.llvm.16733369349353020960
    unreachable)
  (func $runtime.hashmapBinaryGet.llvm.16733369349353020960 (type 3) (param i32 i32) (result i32)
    i32.const 65776
    local.get 0
    local.get 1
    local.get 0
    i32.const 65788
    i32.load
    i32.const 65780
    i32.load
    local.get 0
    call $runtime.hash32.llvm.16733369349353020960
    call $runtime.hashmapGet.llvm.16733369349353020960)
  (func $runtime.hashmapBinaryDelete.llvm.16733369349353020960 (type 1) (param i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32 i32 i32)
    global.get $__stack_pointer
    i32.const 32
    i32.sub
    local.tee 1
    global.set $__stack_pointer
    local.get 1
    i32.const 24
    i32.add
    i64.const 0
    i64.store
    local.get 1
    i64.const 0
    i64.store offset=16
    local.get 1
    i32.const 6
    i32.store offset=4
    i32.const 65988
    i32.load
    local.set 4
    i32.const 65988
    local.get 1
    i32.store
    local.get 1
    local.get 4
    i32.store
    local.get 0
    i32.const 65788
    i32.load
    i32.const 65780
    i32.load
    i32.const 0
    call $runtime.hash32.llvm.16733369349353020960
    local.tee 2
    i32.const 24
    i32.shr_u
    local.tee 3
    i32.const 1
    local.get 3
    select
    local.set 7
    i32.const 65776
    local.get 2
    call $runtime.hashmapBucketAddrForHash.llvm.16733369349353020960
    local.set 2
    block  ;; label = @1
      loop  ;; label = @2
        local.get 1
        local.get 2
        i32.store offset=8
        local.get 1
        local.get 2
        i32.store offset=12
        local.get 2
        i32.eqz
        br_if 1 (;@1;)
        i32.const 0
        local.set 3
        block  ;; label = @3
          loop  ;; label = @4
            local.get 3
            i32.const 8
            i32.ne
            if  ;; label = @5
              local.get 1
              i32.const 65788
              i32.load
              local.tee 8
              local.get 3
              i32.mul
              local.get 2
              i32.add
              i32.const 12
              i32.add
              local.tee 5
              i32.store offset=16
              block  ;; label = @6
                local.get 2
                local.get 3
                i32.add
                local.tee 9
                i32.load8_u
                local.get 7
                i32.ne
                br_if 0 (;@6;)
                local.get 1
                i32.const 65800
                i32.load
                local.tee 10
                i32.store offset=20
                local.get 1
                i32.const 65804
                i32.load
                local.tee 6
                i32.store offset=24
                local.get 6
                i32.eqz
                br_if 3 (;@3;)
                local.get 0
                local.get 5
                local.get 8
                local.get 10
                local.get 6
                call_indirect (type 0)
                i32.const 1
                i32.and
                i32.eqz
                br_if 0 (;@6;)
                local.get 9
                i32.const 0
                i32.store8
                local.get 5
                i32.const 0
                i32.const 65788
                i32.load
                memory.fill
                local.get 1
                i32.const 65792
                i32.load
                local.tee 0
                local.get 3
                i32.mul
                i32.const 65788
                i32.load
                i32.const 3
                i32.shl
                i32.add
                local.get 2
                i32.add
                i32.const 12
                i32.add
                local.tee 2
                i32.store offset=28
                local.get 2
                i32.const 0
                local.get 0
                memory.fill
                i32.const 65784
                i32.const 65784
                i32.load
                i32.const 1
                i32.sub
                i32.store
                br 5 (;@1;)
              end
              local.get 3
              i32.const 1
              i32.add
              local.set 3
              br 1 (;@4;)
            end
          end
          local.get 2
          i32.load offset=8
          local.set 2
          br 1 (;@2;)
        end
      end
      call $runtime.nilPanic.llvm.16733369349353020960
      unreachable
    end
    i32.const 65988
    local.get 4
    i32.store
    local.get 1
    i32.const 32
    i32.add
    global.set $__stack_pointer)
  (func $calloc (type 3) (param i32 i32) (result i32)
    (local i32 i32)
    global.get $__stack_pointer
    i32.const 16
    i32.sub
    local.tee 2
    global.set $__stack_pointer
    i32.const 65988
    i32.load
    local.set 3
    i32.const 65988
    local.get 2
    i32.store
    local.get 0
    local.get 1
    i32.mul
    call $malloc
    i32.const 65988
    local.get 3
    i32.store
    local.get 2
    i32.const 16
    i32.add
    global.set $__stack_pointer)
  (func $realloc (type 3) (param i32 i32) (result i32)
    (local i32 i32 i32 i32)
    global.get $__stack_pointer
    i32.const 32
    i32.sub
    local.tee 2
    global.set $__stack_pointer
    local.get 2
    i32.const 2
    i32.store offset=20
    i32.const 65988
    i32.load
    local.set 4
    i32.const 65988
    local.get 2
    i32.const 16
    i32.add
    i32.store
    local.get 2
    local.get 4
    i32.store offset=16
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 1
          i32.eqz
          if  ;; label = @4
            local.get 0
            call $free
            br 1 (;@3;)
          end
          local.get 1
          i32.const 0
          i32.lt_s
          br_if 1 (;@2;)
          local.get 2
          local.get 1
          call $runtime.alloc.llvm.16733369349353020960
          local.tee 3
          i32.store offset=24
          local.get 2
          local.get 3
          i32.store offset=28
          local.get 0
          if  ;; label = @4
            local.get 2
            local.get 0
            i32.store offset=12
            local.get 2
            i32.const 12
            i32.add
            local.get 2
            call $runtime.hashmapBinaryGet.llvm.16733369349353020960
            i32.const 1
            i32.and
            i32.eqz
            br_if 3 (;@1;)
            local.get 3
            local.get 2
            i32.load
            local.get 2
            i32.load offset=4
            local.tee 5
            local.get 1
            local.get 1
            local.get 5
            i32.gt_u
            select
            memory.copy
            local.get 2
            local.get 0
            i32.store
            local.get 2
            call $runtime.hashmapBinaryDelete.llvm.16733369349353020960
          end
          local.get 2
          local.get 1
          i32.store offset=8
          local.get 2
          local.get 1
          i32.store offset=4
          local.get 2
          local.get 3
          i32.store
          local.get 2
          local.get 3
          i32.store offset=12
          local.get 2
          i32.const 12
          i32.add
          local.get 2
          call $runtime.hashmapBinarySet.llvm.16733369349353020960
        end
        i32.const 65988
        local.get 4
        i32.store
        local.get 2
        i32.const 32
        i32.add
        global.set $__stack_pointer
        local.get 3
        return
      end
      call $runtime.slicePanic.llvm.16733369349353020960
      unreachable
    end
    i32.const 65616
    call $runtime._panic.llvm.16733369349353020960
    unreachable)
  (func $_start (type 4)
    (local i32 i32)
    i32.const 65816
    memory.size
    i32.const 16
    i32.shl
    local.tee 0
    i32.store
    call $runtime.calculateHeapAddresses
    i32.const 65944
    i32.load
    local.tee 1
    i32.const 0
    local.get 0
    local.get 1
    i32.sub
    memory.fill
    i32.const 65816
    memory.size
    i32.const 16
    i32.shl
    i32.store)
  (func $runtime.hashmapGet.llvm.16733369349353020960 (type 0) (param i32 i32 i32 i32) (result i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32 i32)
    global.get $__stack_pointer
    i32.const 48
    i32.sub
    local.tee 4
    global.set $__stack_pointer
    local.get 4
    i32.const 32
    i32.add
    i64.const 0
    i64.store
    local.get 4
    i32.const 40
    i32.add
    i32.const 0
    i32.store
    local.get 4
    i64.const 0
    i64.store offset=24
    local.get 4
    i32.const 7
    i32.store offset=12
    i32.const 65988
    i32.load
    local.set 6
    i32.const 65988
    local.get 4
    i32.const 8
    i32.add
    i32.store
    local.get 4
    local.get 6
    i32.store offset=8
    local.get 4
    local.get 0
    local.get 3
    call $runtime.hashmapBucketAddrForHash.llvm.16733369349353020960
    local.tee 5
    i32.store offset=16
    local.get 3
    i32.const 24
    i32.shr_u
    local.tee 3
    i32.const 1
    local.get 3
    select
    local.set 9
    block  ;; label = @1
      block  ;; label = @2
        loop  ;; label = @3
          block  ;; label = @4
            local.get 4
            local.get 5
            i32.store offset=20
            local.get 5
            i32.eqz
            br_if 0 (;@4;)
            i32.const 0
            local.set 3
            loop  ;; label = @5
              local.get 3
              i32.const 8
              i32.ne
              if  ;; label = @6
                local.get 4
                local.get 0
                i32.load offset=12
                local.tee 7
                local.get 3
                i32.mul
                local.get 5
                i32.add
                i32.const 12
                i32.add
                local.tee 10
                i32.store offset=24
                local.get 4
                local.get 0
                i32.load offset=16
                local.get 3
                i32.mul
                local.get 7
                i32.const 3
                i32.shl
                i32.add
                local.get 5
                i32.add
                i32.const 12
                i32.add
                local.tee 11
                i32.store offset=28
                block  ;; label = @7
                  local.get 3
                  local.get 5
                  i32.add
                  i32.load8_u
                  local.get 9
                  i32.ne
                  br_if 0 (;@7;)
                  local.get 4
                  local.get 0
                  i32.load offset=24
                  local.tee 12
                  i32.store offset=32
                  local.get 4
                  local.get 0
                  i32.load offset=28
                  local.tee 8
                  i32.store offset=36
                  local.get 8
                  i32.eqz
                  br_if 5 (;@2;)
                  local.get 1
                  local.get 10
                  local.get 7
                  local.get 12
                  local.get 8
                  call_indirect (type 0)
                  i32.const 1
                  i32.and
                  i32.eqz
                  br_if 0 (;@7;)
                  local.get 2
                  local.get 11
                  local.get 0
                  i32.load offset=16
                  memory.copy
                  br 6 (;@1;)
                end
                local.get 3
                i32.const 1
                i32.add
                local.set 3
                br 1 (;@5;)
              end
            end
            local.get 4
            local.get 5
            i32.load offset=8
            local.tee 5
            i32.store offset=40
            br 1 (;@3;)
          end
        end
        local.get 2
        i32.const 0
        local.get 0
        i32.load offset=16
        memory.fill
        br 1 (;@1;)
      end
      call $runtime.nilPanic.llvm.16733369349353020960
      unreachable
    end
    i32.const 65988
    local.get 6
    i32.store
    local.get 4
    i32.const 48
    i32.add
    global.set $__stack_pointer
    local.get 5
    i32.const 0
    i32.ne)
  (func $runtime.hashmapBucketAddrForHash.llvm.16733369349353020960 (type 3) (param i32 i32) (result i32)
    local.get 0
    i32.load
    local.get 0
    i32.load offset=16
    local.get 0
    i32.load offset=12
    i32.add
    i32.const 3
    i32.shl
    i32.const 12
    i32.add
    i32.const -1
    i32.const -1
    local.get 0
    i32.load8_u offset=20
    local.tee 0
    i32.shl
    i32.const -1
    i32.xor
    local.get 0
    i32.const 31
    i32.gt_u
    select
    local.get 1
    i32.and
    i32.mul
    i32.add)
  (func $runtime.hashmapSet.llvm.16733369349353020960 (type 8) (param i32 i32 i32 i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32)
    global.get $__stack_pointer
    i32.const 272
    i32.sub
    local.tee 4
    global.set $__stack_pointer
    local.get 4
    i32.const 55
    i32.store offset=44
    local.get 4
    i32.const 48
    i32.add
    i32.const 0
    i32.const 220
    memory.fill
    local.get 4
    i32.const 65988
    i32.load
    local.tee 15
    i32.store offset=40
    i32.const 65988
    local.get 4
    i32.const 40
    i32.add
    i32.store
    block  ;; label = @1
      block  ;; label = @2
        local.get 0
        i32.eqz
        br_if 0 (;@2;)
        block  ;; label = @3
          local.get 0
          i32.load8_u offset=20
          local.tee 5
          i32.const 29
          i32.gt_u
          br_if 0 (;@3;)
          local.get 0
          i32.load offset=8
          i32.const 6
          local.get 5
          i32.shl
          i32.le_u
          br_if 0 (;@3;)
          local.get 4
          i64.const 0
          i64.store offset=16
          local.get 4
          local.get 0
          i32.load offset=36
          local.tee 3
          i32.store offset=64
          local.get 4
          local.get 0
          i32.load offset=32
          local.tee 7
          i32.store offset=60
          local.get 4
          local.get 0
          i32.load offset=28
          local.tee 6
          i32.store offset=56
          local.get 4
          local.get 0
          i32.load offset=24
          local.tee 8
          i32.store offset=52
          local.get 4
          local.get 0
          i32.load
          i32.store offset=48
          local.get 4
          local.get 3
          i32.store offset=36
          local.get 4
          local.get 7
          i32.store offset=32
          local.get 4
          local.get 6
          i32.store offset=28
          local.get 4
          local.get 8
          i32.store offset=24
          local.get 4
          local.get 0
          i32.load offset=16
          i32.store offset=16
          local.get 4
          local.get 0
          i32.load offset=12
          i32.store offset=12
          i32.const 65764
          i32.const 65764
          i32.load
          local.tee 3
          i32.const 7
          i32.shl
          local.get 3
          i32.xor
          local.tee 3
          i32.const 1
          i32.shr_u
          local.get 3
          i32.xor
          local.tee 3
          i32.const 9
          i32.shl
          local.get 3
          i32.xor
          local.tee 3
          i32.store
          local.get 4
          i32.const 0
          i32.store offset=8
          local.get 4
          local.get 3
          i32.store offset=4
          local.get 4
          local.get 5
          i32.const 1
          i32.add
          local.tee 3
          i32.store8 offset=20
          local.get 4
          local.get 0
          i32.load offset=16
          local.get 0
          i32.load offset=12
          i32.add
          i32.const 3
          i32.shl
          i32.const 12
          i32.add
          local.get 3
          i32.shl
          call $runtime.alloc.llvm.16733369349353020960
          local.tee 3
          i32.store
          local.get 4
          local.get 3
          i32.store offset=68
          local.get 4
          local.get 0
          i32.load offset=12
          call $runtime.alloc.llvm.16733369349353020960
          local.tee 8
          i32.store offset=72
          local.get 4
          local.get 0
          i32.load offset=16
          call $runtime.alloc.llvm.16733369349353020960
          local.tee 12
          i32.store offset=76
          i32.const 0
          local.set 7
          i32.const 0
          local.set 3
          i32.const 0
          local.set 5
          i32.const 0
          local.set 6
          loop  ;; label = @4
            local.get 4
            local.get 5
            i32.store offset=80
            local.get 5
            i32.eqz
            if  ;; label = @5
              local.get 4
              local.get 0
              i32.load
              local.tee 5
              i32.store offset=84
              i32.const 1
              local.get 0
              i32.load8_u offset=20
              local.tee 9
              i32.shl
              i32.const 0
              local.get 9
              i32.const 31
              i32.le_u
              select
              local.set 13
            end
            local.get 4
            local.get 5
            i32.store offset=100
            local.get 4
            local.get 5
            i32.store offset=120
            block  ;; label = @5
              loop  ;; label = @6
                local.get 4
                local.get 3
                i32.store offset=88
                local.get 6
                i32.const 255
                i32.and
                i32.const 8
                i32.ge_u
                if  ;; label = @7
                  local.get 3
                  i32.eqz
                  br_if 5 (;@2;)
                  local.get 4
                  local.get 3
                  i32.load offset=8
                  local.tee 3
                  i32.store offset=92
                  i32.const 0
                  local.set 6
                end
                local.get 4
                local.get 3
                i32.store offset=96
                local.get 3
                i32.eqz
                if  ;; label = @7
                  local.get 7
                  local.get 13
                  i32.ge_u
                  br_if 2 (;@5;)
                  local.get 4
                  local.get 5
                  local.get 0
                  i32.load offset=16
                  local.get 0
                  i32.load offset=12
                  i32.add
                  i32.const 3
                  i32.shl
                  i32.const 12
                  i32.add
                  local.get 7
                  i32.mul
                  i32.add
                  local.tee 3
                  i32.store offset=104
                  local.get 7
                  i32.const 1
                  i32.add
                  local.set 7
                end
                local.get 4
                local.get 3
                i32.store offset=112
                local.get 4
                local.get 3
                i32.store offset=128
                local.get 4
                local.get 3
                i32.store offset=108
                local.get 3
                i32.eqz
                br_if 4 (;@2;)
                local.get 3
                local.get 6
                i32.const 255
                i32.and
                local.tee 10
                i32.add
                i32.load8_u
                i32.eqz
                if  ;; label = @7
                  local.get 6
                  i32.const 1
                  i32.add
                  local.set 6
                  br 1 (;@6;)
                end
                local.get 4
                local.get 0
                i32.load offset=12
                local.tee 9
                local.get 10
                i32.mul
                local.get 3
                i32.add
                i32.const 12
                i32.add
                local.tee 11
                i32.store offset=116
                local.get 8
                local.get 11
                local.get 9
                memory.copy
                local.get 4
                local.get 0
                i32.load
                local.tee 11
                i32.store offset=124
                block  ;; label = @7
                  local.get 5
                  local.get 11
                  i32.eq
                  if  ;; label = @8
                    local.get 4
                    local.get 10
                    local.get 0
                    i32.load offset=16
                    local.tee 10
                    i32.mul
                    local.get 9
                    i32.const 3
                    i32.shl
                    i32.add
                    local.get 3
                    i32.add
                    i32.const 12
                    i32.add
                    local.tee 9
                    i32.store offset=132
                    local.get 12
                    local.get 9
                    local.get 10
                    memory.copy
                    local.get 6
                    i32.const 1
                    i32.add
                    local.set 6
                    br 1 (;@7;)
                  end
                  local.get 4
                  local.get 0
                  i32.load offset=32
                  local.tee 11
                  i32.store offset=136
                  local.get 4
                  local.get 0
                  i32.load offset=36
                  local.tee 10
                  i32.store offset=140
                  local.get 10
                  i32.eqz
                  br_if 5 (;@2;)
                  local.get 6
                  i32.const 1
                  i32.add
                  local.set 6
                  local.get 0
                  local.get 8
                  local.get 12
                  local.get 8
                  local.get 9
                  local.get 0
                  i32.load offset=4
                  local.get 11
                  local.get 10
                  call_indirect (type 0)
                  call $runtime.hashmapGet.llvm.16733369349353020960
                  i32.const 1
                  i32.and
                  i32.eqz
                  br_if 1 (;@6;)
                end
              end
              local.get 4
              local.get 4
              i32.load offset=32
              local.tee 10
              i32.store offset=144
              local.get 4
              local.get 4
              i32.load offset=36
              local.tee 9
              i32.store offset=148
              local.get 9
              i32.eqz
              br_if 3 (;@2;)
              local.get 4
              local.get 8
              local.get 12
              local.get 8
              local.get 4
              i32.load offset=12
              local.get 4
              i32.load offset=4
              local.get 10
              local.get 9
              call_indirect (type 0)
              call $runtime.hashmapSet.llvm.16733369349353020960
              br 1 (;@4;)
            end
          end
          local.get 0
          local.get 4
          i32.load
          local.tee 3
          i32.store
          local.get 0
          local.get 4
          i64.load offset=4 align=4
          i64.store offset=4 align=4
          local.get 0
          local.get 4
          i64.load offset=12 align=4
          i64.store offset=12 align=4
          local.get 0
          local.get 4
          i32.load8_u offset=20
          i32.store8 offset=20
          local.get 0
          local.get 4
          i32.load offset=24
          local.tee 5
          i32.store offset=24
          local.get 0
          local.get 4
          i32.load offset=28
          local.tee 7
          i32.store offset=28
          local.get 0
          local.get 4
          i32.load offset=32
          local.tee 6
          i32.store offset=32
          local.get 0
          local.get 4
          i32.load offset=36
          local.tee 8
          i32.store offset=36
          local.get 4
          local.get 3
          i32.store offset=152
          local.get 4
          local.get 5
          i32.store offset=156
          local.get 4
          local.get 7
          i32.store offset=160
          local.get 4
          local.get 6
          i32.store offset=164
          local.get 4
          local.get 8
          i32.store offset=168
          local.get 4
          local.get 0
          i32.load offset=32
          local.tee 5
          i32.store offset=172
          local.get 4
          local.get 0
          i32.load offset=36
          local.tee 3
          i32.store offset=176
          local.get 3
          i32.eqz
          br_if 1 (;@2;)
          local.get 1
          local.get 0
          i32.load offset=12
          local.get 0
          i32.load offset=4
          local.get 5
          local.get 3
          call_indirect (type 0)
          local.set 3
        end
        local.get 4
        local.get 0
        local.get 3
        call $runtime.hashmapBucketAddrForHash.llvm.16733369349353020960
        local.tee 7
        i32.store offset=180
        local.get 3
        i32.const 24
        i32.shr_u
        local.tee 3
        i32.const 1
        local.get 3
        select
        local.set 9
        i32.const 0
        local.set 3
        i32.const 0
        local.set 6
        i32.const 0
        local.set 12
        i32.const 0
        local.set 8
        loop  ;; label = @3
          block  ;; label = @4
            local.get 4
            local.get 3
            i32.store offset=196
            local.get 4
            local.get 7
            local.tee 5
            i32.store offset=200
            local.get 4
            local.get 6
            i32.store offset=192
            local.get 4
            local.get 12
            i32.store offset=188
            local.get 4
            local.get 8
            i32.store offset=184
            local.get 5
            i32.eqz
            br_if 0 (;@4;)
            i32.const 0
            local.set 3
            loop  ;; label = @5
              block  ;; label = @6
                local.get 4
                local.get 12
                i32.store offset=208
                local.get 4
                local.get 6
                i32.store offset=212
                local.get 4
                local.get 8
                i32.store offset=204
                local.get 3
                i32.const 8
                i32.eq
                br_if 0 (;@6;)
                local.get 4
                local.get 0
                i32.load offset=12
                local.tee 7
                local.get 3
                i32.mul
                local.get 5
                i32.add
                i32.const 12
                i32.add
                local.tee 13
                i32.store offset=216
                local.get 4
                local.get 0
                i32.load offset=16
                local.get 3
                i32.mul
                local.get 7
                i32.const 3
                i32.shl
                i32.add
                local.get 5
                i32.add
                i32.const 12
                i32.add
                local.tee 10
                i32.store offset=220
                local.get 4
                local.get 6
                local.get 13
                local.get 3
                local.get 5
                i32.add
                local.tee 11
                i32.load8_u
                local.get 6
                i32.or
                local.tee 14
                select
                local.tee 6
                i32.store offset=232
                local.get 4
                local.get 8
                local.get 11
                local.get 14
                select
                local.tee 8
                i32.store offset=224
                local.get 4
                local.get 12
                local.get 10
                local.get 14
                select
                local.tee 12
                i32.store offset=228
                block  ;; label = @7
                  local.get 11
                  i32.load8_u
                  local.get 9
                  i32.ne
                  br_if 0 (;@7;)
                  local.get 4
                  local.get 0
                  i32.load offset=24
                  local.tee 14
                  i32.store offset=236
                  local.get 4
                  local.get 0
                  i32.load offset=28
                  local.tee 11
                  i32.store offset=240
                  local.get 11
                  i32.eqz
                  br_if 5 (;@2;)
                  local.get 1
                  local.get 13
                  local.get 7
                  local.get 14
                  local.get 11
                  call_indirect (type 0)
                  i32.const 1
                  i32.and
                  i32.eqz
                  br_if 0 (;@7;)
                  local.get 10
                  local.get 2
                  local.get 0
                  i32.load offset=16
                  memory.copy
                  br 6 (;@1;)
                end
                local.get 3
                i32.const 1
                i32.add
                local.set 3
                br 1 (;@5;)
              end
            end
            local.get 4
            local.get 5
            i32.load offset=8
            local.tee 7
            i32.store offset=244
            local.get 5
            local.set 3
            br 1 (;@3;)
          end
        end
        local.get 6
        i32.eqz
        if  ;; label = @3
          local.get 0
          i32.load offset=16
          local.get 0
          i32.load offset=12
          i32.add
          i32.const 3
          i32.shl
          i32.const 12
          i32.add
          call $runtime.alloc.llvm.16733369349353020960
          local.set 5
          local.get 0
          local.get 0
          i32.load offset=8
          i32.const 1
          i32.add
          i32.store offset=8
          local.get 4
          local.get 5
          i32.store offset=252
          local.get 4
          local.get 5
          i32.store offset=264
          local.get 4
          local.get 5
          i32.store offset=248
          local.get 4
          local.get 5
          i32.const 12
          i32.add
          local.tee 7
          i32.store offset=256
          local.get 4
          local.get 7
          local.get 0
          i32.load offset=12
          local.tee 6
          i32.const 3
          i32.shl
          i32.add
          local.tee 8
          i32.store offset=260
          local.get 7
          local.get 1
          local.get 6
          memory.copy
          local.get 8
          local.get 2
          local.get 0
          i32.load offset=16
          memory.copy
          local.get 5
          local.get 9
          i32.store8
          local.get 3
          i32.eqz
          br_if 1 (;@2;)
          local.get 3
          local.get 5
          i32.store offset=8
          br 2 (;@1;)
        end
        local.get 0
        local.get 0
        i32.load offset=8
        i32.const 1
        i32.add
        i32.store offset=8
        local.get 6
        local.get 1
        local.get 0
        i32.load offset=12
        memory.copy
        local.get 12
        local.get 2
        local.get 0
        i32.load offset=16
        memory.copy
        local.get 8
        i32.eqz
        br_if 0 (;@2;)
        local.get 8
        local.get 9
        i32.store8
        br 1 (;@1;)
      end
      call $runtime.nilPanic.llvm.16733369349353020960
      unreachable
    end
    i32.const 65988
    local.get 15
    i32.store
    local.get 4
    i32.const 272
    i32.add
    global.set $__stack_pointer)
  (func $__wasm_export_exports_dapr_http_http_handler_handle_http_request_post_return (type 1) (param i32)
    (local i32 i32 i32 i32)
    block  ;; label = @1
      local.get 0
      i32.const 4
      i32.add
      i32.load8_u
      i32.const 1
      i32.ne
      br_if 0 (;@1;)
      local.get 0
      i32.const 12
      i32.add
      i32.load
      local.tee 3
      i32.const 0
      local.get 3
      i32.const 0
      i32.gt_s
      select
      local.set 2
      local.get 0
      i32.const 8
      i32.add
      i32.load
      local.tee 4
      local.set 1
      loop  ;; label = @2
        local.get 2
        i32.eqz
        if  ;; label = @3
          local.get 3
          i32.const 0
          i32.le_s
          br_if 2 (;@1;)
          local.get 4
          call $free
          br 2 (;@1;)
        end
        local.get 1
        i32.const 4
        i32.add
        i32.load
        i32.const 0
        i32.gt_s
        if  ;; label = @3
          local.get 1
          i32.load
          call $free
        end
        local.get 1
        i32.const 12
        i32.add
        i32.load
        i32.const 0
        i32.gt_s
        if  ;; label = @3
          local.get 1
          i32.const 8
          i32.add
          i32.load
          call $free
        end
        local.get 2
        i32.const 1
        i32.sub
        local.set 2
        local.get 1
        i32.const 16
        i32.add
        local.set 1
        br 0 (;@2;)
      end
      unreachable
    end
    block  ;; label = @1
      local.get 0
      i32.const 16
      i32.add
      i32.load8_u
      i32.const 1
      i32.ne
      br_if 0 (;@1;)
      local.get 0
      i32.const 24
      i32.add
      i32.load
      i32.const 0
      i32.le_s
      br_if 0 (;@1;)
      local.get 0
      i32.const 20
      i32.add
      i32.load
      call $free
    end)
  (func $cabi_realloc (type 0) (param i32 i32 i32 i32) (result i32)
    block  ;; label = @1
      local.get 3
      i32.eqz
      br_if 0 (;@1;)
      local.get 0
      local.get 3
      call $realloc
      local.tee 2
      br_if 0 (;@1;)
      unreachable
    end
    local.get 2)
  (func $http_string_free (type 1) (param i32)
    local.get 0
    i32.load offset=4
    if  ;; label = @1
      local.get 0
      i32.load
      call $free
    end
    local.get 0
    i64.const 0
    i64.store align=4)
  (func $dapr_http_http_types_headers_free (type 1) (param i32)
    (local i32 i32 i32)
    loop  ;; label = @1
      local.get 0
      i32.load offset=4
      local.tee 1
      local.get 2
      i32.le_u
      if  ;; label = @2
        local.get 1
        if  ;; label = @3
          local.get 0
          i32.load
          call $free
        end
      else
        local.get 0
        i32.load
        local.get 3
        i32.add
        local.tee 1
        call $http_string_free
        local.get 1
        i32.const 8
        i32.add
        call $http_string_free
        local.get 3
        i32.const 16
        i32.add
        local.set 3
        local.get 2
        i32.const 1
        i32.add
        local.set 2
        br 1 (;@1;)
      end
    end)
  (func $__wasm_export_exports_dapr_http_http_handler_handle_http_request (type 9) (param i32 i32 i32 i32 i32 i32 i32 i32 i32 i32) (result i32)
    (local i32)
    global.get $__stack_pointer
    i32.const 80
    i32.sub
    local.tee 10
    global.set $__stack_pointer
    local.get 10
    i32.const 76
    i32.add
    local.get 9
    i32.store
    local.get 10
    i32.const 72
    i32.add
    local.get 8
    i32.store
    local.get 10
    i32.const -64
    i32.sub
    local.get 6
    i32.store
    local.get 10
    i32.const 56
    i32.add
    local.get 4
    i32.store
    local.get 10
    i32.const 48
    i32.add
    local.get 2
    i32.store
    local.get 10
    local.get 5
    i32.store offset=60
    local.get 10
    local.get 3
    i32.store offset=52
    local.get 10
    local.get 1
    i32.store offset=44
    local.get 10
    local.get 0
    i32.store8 offset=40
    local.get 10
    local.get 7
    i32.const 1
    i32.eq
    i32.store8 offset=68
    local.get 10
    i32.const 40
    i32.add
    local.get 10
    i32.const 8
    i32.add
    call $exports_dapr_http_http_handler_handle_http_request
    i32.const 66276
    local.get 10
    i32.load16_u offset=8
    i32.store16
    block  ;; label = @1
      local.get 10
      i32.load8_u offset=12
      if  ;; label = @2
        i32.const 66280
        i32.const 1
        i32.store8
        i32.const 66284
        local.get 10
        i32.const 16
        i32.add
        i64.load
        i64.store align=4
        br 1 (;@1;)
      end
      i32.const 66280
      i32.const 0
      i32.store8
    end
    block  ;; label = @1
      local.get 10
      i32.load8_u offset=24
      if  ;; label = @2
        i32.const 66292
        i32.const 1
        i32.store8
        i32.const 66296
        local.get 10
        i32.const 28
        i32.add
        i64.load align=4
        i64.store align=4
        br 1 (;@1;)
      end
      i32.const 66292
      i32.const 0
      i32.store8
    end
    local.get 10
    i32.const 80
    i32.add
    global.set $__stack_pointer
    i32.const 66276)
  (table (;0;) 3 3 funcref)
  (memory (;0;) 2)
  (global $__stack_pointer (mut i32) (i32.const 65536))
  (export "memory" (memory 0))
  (export "exports_dapr_http_http_handler_handle_http_request" (func $exports_dapr_http_http_handler_handle_http_request))
  (export "malloc" (func $malloc))
  (export "free" (func $free))
  (export "calloc" (func $calloc))
  (export "realloc" (func $realloc))
  (export "_start" (func $_start))
  (export "cabi_post_dapr:http/http-handler#handle-http-request" (func $__wasm_export_exports_dapr_http_http_handler_handle_http_request_post_return))
  (export "cabi_realloc" (func $cabi_realloc))
  (export "dapr:http/http-handler#handle-http-request" (func $__wasm_export_exports_dapr_http_http_handler_handle_http_request))
  (elem (;0;) (i32.const 1) func $runtime.memequal $runtime.hash32.llvm.16733369349353020960)
  (data $.rodata (i32.const 65536) "Option is None\00\00\00\00\01\00\0e\00\00\00free: invalid pointer\00\00\00\18\00\01\00\15\00\00\00realloc: invalid pointer8\00\01\00\18\00\00\00out of memorypanic: panic: runtime error: nil pointer dereferenceindex out of rangeslice out of rangeContent-Typetext/plainHello from WASM!")
  (data $.data (i32.const 65764) "\c1\82\01\00 \01\01\00\00\00\00\00\cc\01\01\00\c1\82\01\00\00\00\00\00\04\00\00\00\0c\00\00\00\01\00\00\00\00\00\00\00\01\00\00\00\00\00\00\00\02"))
