package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {

	health := map[string]any{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	err := app.writeJSON(w, http.StatusOK, health, nil)
	if err != nil {
		app.errorLog.Println(err)
		app.serverError(w, err)
	}
}

func (app *application) createURLHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Create URL"))
}

func (app *application) showURLHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		app.notFound(w, r)
		return
	}

	w.Write([]byte(fmt.Sprintf("Show url id %d", id)))
}

func (app *application) Index(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello World"))
}

func (app *application) ShortenURL(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello World"))
}

func (app *application) ExpandURL(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello World"))
}
