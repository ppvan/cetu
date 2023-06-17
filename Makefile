BINARY_NAME=cetu

build:
	go build -o ./bin/${BINARY_NAME} ./cmd/...

tidy:
	go mod tidy

run:
	go run ./cmd/...
clean:
	rm -rf bin/${BINARY_NAME}