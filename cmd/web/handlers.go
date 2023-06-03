package main

import (
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ppvan/cetu/internal/models"
)

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	numbers := 100000000001
	w.Write([]byte("Shorten URL: " + Base62Encode(uint64(numbers))))
}

func (app *application) ExpandURL(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	shortenURL := params.ByName("shorten")

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
