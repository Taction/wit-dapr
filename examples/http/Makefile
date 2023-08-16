

wasm:
	tinygo build -o main.wasm -scheduler=none --no-debug -target=wasi main.go

wasm2wat:
	wasm2wat main.wasm -o main.wat