package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	// Get the stack trace for the error and output it to the error log.
	// Pop 2 step to log the caller of the function that called serverError.
	trace := fmt.Sprintf("%s\n%s\n", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
}

func (app *application) methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("Method %s not allowed", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func (app *application) notFound(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotFound, "Resource not found")
}

func (app *application) badRequest(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	response := envelope{
		"error": message,
	}

	err := app.writeJSON(w, status, response, nil)
	if err != nil {
		app.errorLog.Println(err)
	}
}
