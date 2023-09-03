VERSION=v0.0.3

wasm:
	tinygo build -o main.wasm --no-debug -target=wasi main.go

wasm2wat:
	wasm2wat main.wasm -o main.wat

image:
	docker build -t docker4zc/dwhttp:$(VERSION) .
	docker push docker4zc/dwhttp:$(VERSION)