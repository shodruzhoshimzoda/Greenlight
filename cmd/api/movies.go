package main

import (
	"fmt"
	"net/http"

)




func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Create a new movie")
} 

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request)  {


	id, err := app.readIdParams(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}


	fmt.Fprintf(w, "Show the details for movie %d", id)


} 	