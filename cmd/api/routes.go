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

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthCheck)
	router.HandlerFunc(http.MethodGet, "/v1/urls/:id", app.showURLHandler)
	router.HandlerFunc(http.MethodGet, "/v1/urls", app.createURLHandler)
	// router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthCheck)

	return standard.Then(router)
}
