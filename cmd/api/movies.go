package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/shodruzhoshimzoda/Greenlight/internal/data"
	"github.com/shodruzhoshimzoda/Greenlight/internal/validator"
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

	v := validator.New()
	v.Check(input.Title != "", "title", "must be provided")
	v.Check(len(input.Title) <= 500, "title", "must not be more than 500 bytes")
	v.Check(input.Year != 0, "year", "must be provided")
	v.Check(input.Year >= 1888, "year", "must be greater than or equal to 2018")
	v.Check(input.Year <= int32(time.Now().Year()), "year", "must be greater than or equal to current year")
	v.Check(input.Runtime != 0, "runtime", "must be provided")
	v.Check(input.Runtime > 0, "runtime", "must not be positive number")
	v.Check(input.Genres != nil, "genres", "must be provided")
	v.Check(len(input.Genres) > 1, "genres", "genres must contain at least 1 item")
	v.Check(len(input.Genres) < 5, "genres", "genres must not contain at least 5 item")
	v.Check(validator.Unique(input.Genres), "genres", "genres must not contain duplicate values")

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
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
