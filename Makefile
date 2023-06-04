build:
	./ui/tailwindcss -i ./ui/main.css -o ./ui/main.min.css --minify
	go build ./cmd/...

run:
	go run ./cmd/...

watch:
	./ui/tailwindcss -i ./ui/main.css -o ./ui/main.min.css --watch