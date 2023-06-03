package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {

	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFound)

	staticFiles := http.FileServer(http.Dir("./ui/static/"))

	router.HandlerFunc(http.MethodGet, "/app/:shorten", app.ExpandURL)
	router.HandlerFunc(http.MethodGet, "/", app.Index)
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", staticFiles))

	return router
}
