package main

import (
	"fmt"
	"net/http"
)

// logError - method for logging error
func (app *application) logError(r *http.Request, error error) {
	var (
		method = r.Method
		url    = r.URL.RequestURI()
	)

	app.logger.Error(error.Error(), "method", method, "url",url)
}


// errorResponse - helper for sending JSON formated error
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {

	data := envelope{"error":message}

	
	if err := app.writeJSON(w, status, data, nil); err!=nil{
		app.logError(r, err)
		w.WriteHeader(500)
		return
	}

}
// The notFoundResponse() method will be used to send a 404 Not Found status code and
// JSON response to the client.
func (app *application) serverError(w http.ResponseWriter, r *http.Request, error error) {
	app.logError(r, error)

	message := "The server encountered error and could not proccess your request"

	app.errorResponse(w, r, http.StatusInternalServerError, message)
}


func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	
	message := "the requested rssource could not be found"

	app.errorResponse(w, r, http.StatusNotFound, message)
	
}


func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("The %s method is not supported for this response", r.Method)

	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}


