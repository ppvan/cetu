BINARY_NAME=cetu

build:
	./ui/tailwindcss -i ui/static/css/main.css -o ui/static/css/main.min.css --minify
	go build -o ./bin/${BINARY_NAME} ./cmd/...

tidy:
	go mod tidy

run:
	go run ./cmd/...
clean:
	rm -rf ui/static/css/main.min.css
	rm -rf bin/${BINARY_NAME}
watch:
	./ui/tailwindcss -i ui/static/css/main.css -o ui/static/css/main.min.css --watch