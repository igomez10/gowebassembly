build-wasm:
	GOOS=js GOARCH=wasm go build -o dist/main.wasm .

build:
	go build .

test:
	go test ./...

