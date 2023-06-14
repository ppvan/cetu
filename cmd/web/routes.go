package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {

	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFound)

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	staticFiles := http.FileServer(http.Dir("./ui/static/"))

	router.HandlerFunc(http.MethodGet, "/z/:url", app.ExpandURL)
	router.HandlerFunc(http.MethodGet, "/", app.Index)
	router.HandlerFunc(http.MethodPost, "/", app.ShortenURL)
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", staticFiles))

	return standard.Then(router)
}
