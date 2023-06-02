package main

import (
	"net/http"
)

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	numbers := 100000000001
	w.Write([]byte("Shorten URL: " + Base62Encode(uint64(numbers))))
}

func ExpandURL(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Expand URL: Hello, World!"))
}
