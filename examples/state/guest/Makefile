VERSION=v0.0.3

wasm:
	tinygo build -o main.wasm --no-debug -target=wasi main.go

wat:
	wasm2wat main.wasm -o main.wat

image:
	docker build -t docker4zc/dwstate:$(VERSION) .
	docker push docker4zc/dwstate:$(VERSION)