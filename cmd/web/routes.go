package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {

	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFound)

	router.HandlerFunc(http.MethodGet, "/app/:shorten", app.ExpandURL)
	router.HandlerFunc(http.MethodGet, "/", ShortenURL)

	return router
}
