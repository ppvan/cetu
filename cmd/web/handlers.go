package main

import (
	"errors"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"github.com/ppvan/cetu/internal/models"
)

func (app *application) Index(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/pages/index.html",
		"./ui/html/layouts/base.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
	}

	ts.ExecuteTemplate(w, "base", nil)
}

func (app *application) ShortenURL(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	originalUrl := r.PostForm.Get("url")

	url, err := app.GenerateShortURL()
	if err != nil {
		app.serverError(w, err)
		return
	}

	_, err = app.urlModel.Insert(url, originalUrl)

	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func (app *application) ExpandURL(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	shortenURL := params.ByName("url")

	url, err := app.urlModel.Get(shortenURL)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w, r)
			return
		}

		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, url.Original, http.StatusMovedPermanently)
}
