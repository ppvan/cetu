package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type envelope map[string]interface{}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {

	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	js = append(js, '\n')

	for key, val := range headers {
		w.Header()[key] = val
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {

	max_bytes := 1_048_576 // 1MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(max_bytes))

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {

		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		// JSON syntax error
		case errors.As(err, &syntaxError):
			return errors.New("body contains badly-formed JSON (at character " + strconv.Itoa(int(syntaxError.Offset)) + ")")

		// JSON type error. For example, when the JSON value is a string but the target is an int
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		// Body is empty
		case err.Error() == "EOF":
			return errors.New("request body must not be empty")

		// This means the dst is null -> programmer errors -> panic
		case errors.As(err, &invalidUnmarshalError):
			panic(err)

		// Limit the size of the request body to 1MB
		case err.Error() == "http: request body too large":
			return errors.New("request body must not be larger than 1MB")

		// Unknown field in JSON
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("body contains unknown key %s", fieldName)
		default:
			return err
		}
	}

	return nil
}

func (app *application) readIDParam(r *http.Request) (int64, error) {

	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid ID parameter")
	}

	return id, nil
}

// Generate a short URL from the last inserted ID
// The ID is fetched from the database
func (app *application) GenerateShortURL() (string, error) {

	id, err := app.urlModel.GetLastInsertId()
	if err != nil {
		return "", err
	}

	slug := base62Encode(id)
	// url := fmt.Sprintf("%s/%s", app.config.BaseURL(), slug)

	return slug, nil
}
