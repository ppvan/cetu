package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {

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

var BASE_CHARS = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
var BASE int64 = int64(len(BASE_CHARS))
var BASE_COUNTER int64 = 100_000_000_000
var MAX_LENGTH = 7

// base62Encode encodes a number to a base62 string.
// Param number: the number to encode.
// Return: the base62 string.
func base62Encode(number int64) (str string) {
	result := make([]byte, MAX_LENGTH)
	i := 0
	number += BASE_COUNTER
	for number > 0 {
		result[i] = BASE_CHARS[number%BASE]
		number = number / BASE
		i++
	}
	return string(reverseBytes(result[:i]))
}

// ReverseBytes reverses a byte array.
// It also change the original array.
// Examples: []byte("Hello, World!") -> []byte("!dlroW ,olleH")
func reverseBytes(bytes []byte) []byte {
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}

	return bytes
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	// Get the stack trace for the error and output it to the error log.
	// Pop 2 step to log the caller of the function that called serverError.
	trace := fmt.Sprintf("%s\n%s\n", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal Server Error"))
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter, r *http.Request) {
	app.clientError(w, http.StatusNotFound)
}