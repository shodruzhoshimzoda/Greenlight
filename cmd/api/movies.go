package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/shodruzhoshimzoda/Greenlight/internal/data"
)




func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request)  {
	var input struct {
		Title 			string		`json:"title"`
		Year			int32 		`json:"year"`
		Runtime 		int			`json:"runtime"`
		Genres			[]string	`json:"genres"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		app.errorResponse(w, r, http.StatusBadRequest,  err.Error())
		return
	}

	fmt.Fprintf(w, "%+v", input)

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