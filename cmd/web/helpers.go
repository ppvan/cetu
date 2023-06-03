package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

var BASE_CHARS = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
var BASE uint64 = uint64(len(BASE_CHARS))
var BASE_COUNTER uint64 = 100_000_000_000
var MAX_LENGTH = 7

// Base62Encode encodes a number to a base62 string.
// Param number: the number to encode.
// Return: the base62 string.
func Base62Encode(number uint64) (str string) {
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
