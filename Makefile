build:
	./ui/tailwindcss -i ui/static/css/main.css -o ui/static/css/main.min.css --minify
	go build ./cmd/...

run:
	go run ./cmd/...

watch:
	./ui/tailwindcss -i ui/static/css/main.css -o ui/static/css/main.min.css --watch