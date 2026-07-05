package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/shodruzhoshimzoda/Greenlight/internal/data"
)




func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Create a new movie")
} 

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request)  {


	id, err := app.readIdParams(r)

	if err != nil {
		app.notFoundResponse(w,r)
		return
	}


	data := data.Movie{
		ID: id,
		CreatedAt: time.Now(),
		Title: "Casablanca",
		Year: 2006,
		Genres: []string{"roman","drama","tragedy"},
		Runtime: 102,
	}


	err = app.writeJSON(w, 201, envelope{"movies":data}, nil)
	if err != nil {
		app.serverError(w,r, err)
	}

	


} 	