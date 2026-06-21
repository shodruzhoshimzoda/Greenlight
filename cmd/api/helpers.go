package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Read id parametr from request
func (app *application) readIdParams(r *http.Request) (int64, error) {

	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)  

	if err != nil || id < 1 {
		return 0, errors.New("invalid id parametrs")
	}

	return id, nil


	
}


// This helper method help us to write JSON
func (app *application) writeJSON(w http.ResponseWriter, status int, data any, header http.Header) error {


	js, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		return  err
	}

	

	for key, value := range header {
		w.Header()[key] = value
	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}