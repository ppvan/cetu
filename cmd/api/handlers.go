package main

import (
	"net/http"
)

func (app *application) Index(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello World"))
}

func (app *application) ShortenURL(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello World"))
}

func (app *application) ExpandURL(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello World"))
}
