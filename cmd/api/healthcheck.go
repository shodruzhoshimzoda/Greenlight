package main

import (
	"net/http"
)



func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"status":"available",
		"environment":app.cfg.env,
		"version":version,
	}


	err := app.writeJSON(w, 200, data, nil) 
	
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "the server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}

	


}