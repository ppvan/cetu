package main

import (
	"net/http"
)

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Shorten URL: Hello !"))
}

func ExpandURL(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Expand URL: Hello, World!"))
}
