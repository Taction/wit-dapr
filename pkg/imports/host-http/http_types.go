package host_http

// inspired from https://github.com/moznion/go-optional

type optionKind int

const (
	none optionKind = iota
	some
)

type Option[T any] struct {
	kind optionKind
	val  T
}

// IsNone returns true if the option is None.
func (o Option[T]) IsNone() bool {
	return o.kind == none
}

// IsSome returns true if the option is Some.
func (o Option[T]) IsSome() bool {
	return o.kind == some
}

// Unwrap returns the value if the option is Some.
func (o Option[T]) Unwrap() T {
	if o.kind != some {
		panic("Option is None")
	}
	return o.val
}

// Set sets the value and returns it.
func (o *Option[T]) Set(val T) T {
	o.kind = some
	o.val = val
	return val
}

// Unset sets the value to None.
func (o *Option[T]) Unset() {
	o.kind = none
}

// Some is a constructor for Option[T] which represents Some.
func Some[T any](v T) Option[T] {
	return Option[T]{
		kind: some,
		val:  v,
	}
}

// None is a constructor for Option[T] which represents None.
func None[T any]() Option[T] {
	return Option[T]{
		kind: none,
	}
}

type ResultKind int

const (
	Ok ResultKind = iota
	Err
)

type Result[T any, E any] struct {
	Kind ResultKind
	Val  T
	Err  E
}

func (r Result[T, E]) IsOk() bool {
	return r.Kind == Ok
}

func (r Result[T, E]) IsErr() bool {
	return r.Kind == Err
}

func (r Result[T, E]) Unwrap() T {
	if r.Kind != Ok {
		panic("Result is Err")
	}
	return r.Val
}

func (r Result[T, E]) UnwrapErr() E {
	if r.Kind != Err {
		panic("Result is Ok")
	}
	return r.Err
}

func (r *Result[T, E]) Set(val T) T {
	r.Kind = Ok
	r.Val = val
	return val
}

func (r *Result[T, E]) SetErr(err E) E {
	r.Kind = Err
	r.Err = err
	return err
}

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

func DaprHttpHttpTypesMethodGet() DaprHttpHttpTypesMethod {
	return DaprHttpHttpTypesMethod{kind: DaprHttpHttpTypesMethodKindGet}
}

func DaprHttpHttpTypesMethodPost() DaprHttpHttpTypesMethod {
	return DaprHttpHttpTypesMethod{kind: DaprHttpHttpTypesMethodKindPost}
}

func DaprHttpHttpTypesMethodPut() DaprHttpHttpTypesMethod {
	return DaprHttpHttpTypesMethod{kind: DaprHttpHttpTypesMethodKindPut}
}

func DaprHttpHttpTypesMethodDelete() DaprHttpHttpTypesMethod {
	return DaprHttpHttpTypesMethod{kind: DaprHttpHttpTypesMethodKindDelete}
}

func DaprHttpHttpTypesMethodPatch() DaprHttpHttpTypesMethod {
	return DaprHttpHttpTypesMethod{kind: DaprHttpHttpTypesMethodKindPatch}
}

func DaprHttpHttpTypesMethodHead() DaprHttpHttpTypesMethod {
	return DaprHttpHttpTypesMethod{kind: DaprHttpHttpTypesMethodKindHead}
}

func DaprHttpHttpTypesMethodOptions() DaprHttpHttpTypesMethod {
	return DaprHttpHttpTypesMethod{kind: DaprHttpHttpTypesMethodKindOptions}
}

type DaprHttpHttpTypesRequest struct {
	Method  DaprHttpHttpTypesMethod
	Uri     string
	Headers []DaprHttpHttpTypesTuple2StringStringT
	Params  []DaprHttpHttpTypesTuple2StringStringT
	Body    Option[[]uint8]
}

type DaprHttpHttpTypesResponse struct {
	Status  uint16
	Headers Option[[]DaprHttpHttpTypesTuple2StringStringT]
	Body    Option[[]uint8]
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

func DaprHttpHttpTypesHttpErrorSuccess() DaprHttpHttpTypesHttpError {
	return DaprHttpHttpTypesHttpError{kind: DaprHttpHttpTypesHttpErrorKindSuccess}
}

func DaprHttpHttpTypesHttpErrorDestinationNotAllowed() DaprHttpHttpTypesHttpError {
	return DaprHttpHttpTypesHttpError{kind: DaprHttpHttpTypesHttpErrorKindDestinationNotAllowed}
}

func DaprHttpHttpTypesHttpErrorInvalidUrl() DaprHttpHttpTypesHttpError {
	return DaprHttpHttpTypesHttpError{kind: DaprHttpHttpTypesHttpErrorKindInvalidUrl}
}

func DaprHttpHttpTypesHttpErrorRequestError() DaprHttpHttpTypesHttpError {
	return DaprHttpHttpTypesHttpError{kind: DaprHttpHttpTypesHttpErrorKindRequestError}
}

func DaprHttpHttpTypesHttpErrorRuntimeError() DaprHttpHttpTypesHttpError {
	return DaprHttpHttpTypesHttpError{kind: DaprHttpHttpTypesHttpErrorKindRuntimeError}
}

func DaprHttpHttpTypesHttpErrorTooManyRequests() DaprHttpHttpTypesHttpError {
	return DaprHttpHttpTypesHttpError{kind: DaprHttpHttpTypesHttpErrorKindTooManyRequests}
}

type DaprHttpHttpHandlerRequest = DaprHttpHttpTypesRequest
type DaprHttpHttpHandlerResponse = DaprHttpHttpTypesResponse
