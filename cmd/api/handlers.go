package main

import (
	"net/http"
	"time"

	"github.com/ppvan/cetu/internal/models"
	"github.com/ppvan/cetu/internal/validators"
)

func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {

	status, err := models.DatabaseStatusCheck(app.db)
	if err != nil {
		app.errorLog.Println(err)
	}

	health := envelope{
		"status": "available",
		"system": map[string]any{
			"environment": app.config.env,
			"version":     version,
		},
		"database": status,
	}

	err = app.writeJSON(w, http.StatusOK, health, nil)
	if err != nil {
		app.errorLog.Println(err)
		app.serverError(w, r, err)
	}
}

func (app *application) createURLHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		OriginalURL string `json:"originalUrl"`
		CustomAlias string `json:"customAlias"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	v := validators.New()
	v.Check(input.OriginalURL != "", "originalUrl", "must be provided")
	v.Check(v.IsBlank(input.OriginalURL), "originalUrl", "must not be blank")
	v.Check(v.IsURL(input.OriginalURL), "originalUrl", "must be a valid URL")

	if !v.Valid() {
		app.failedValidation(w, r, v.Errors)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"url": input}, nil)
}

func (app *application) showURLHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		app.notFound(w, r)
		return
	}

	url := models.URL{
		ID:          id,
		OriginalURL: "https://www.google.com",
		ShortURL:    "https://cetu.com/abc123",
		Clicks:      0,
		ExpiredTime: time.Now(),
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"url": url}, nil)
	if err != nil {
		app.serverError(w, r, err)
		app.errorLog.Println(err)
	}
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
