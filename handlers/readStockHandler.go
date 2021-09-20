package handlers

import (
	"io/ioutil"
	"net/http"
)

func ReadStockHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 Method not found."))
		return
	}
	// check the request body
	//     return if request body empty/ wrong type
	_, err := ioutil.ReadAll(r.Body)
	if err != nil {

	}
	// unmarshall the body to struct (import from models)
	//     this time it is only the ID we have
	//     return error if cannot unmarshal
	// validate the data
	//     return error if invalid
	// connect to the database
	//     return error if you can't connect
	// find database entry
	//     return the database entry or error if the read was unsuccessful

	w.WriteHeader(200)
	w.Write([]byte("read stock"))
}
