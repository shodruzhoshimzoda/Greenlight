package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/shodruzhoshimzoda/Greenlight/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}

	if err := app.readJSON(w, r, &input); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	fmt.Fprintf(w, "%+v+\n", input)
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIdParams(r)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Year:      2006,
		Genres:    []string{"roman", "drama", "tragedy"},
		Runtime:   102,
	}

	err = app.writeJSON(w, 201, envelope{"movies": movie}, nil)
	if err != nil {
		app.serverError(w, r, err)
	}

}
