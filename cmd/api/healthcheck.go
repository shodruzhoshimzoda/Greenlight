package main

import (
	"net/http"
)



func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {

	data := envelope{
		"status":"available",
		"system_info": map[string]string{
			"environment":app.cfg.env,
		"version":version,
		},
		
	}


	err := app.writeJSON(w, 200, data, nil) 
	
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	


}