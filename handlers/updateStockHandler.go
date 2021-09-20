package handlers

import (
	"io/ioutil"
	"net/http"
)

func UpdateStockHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
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
	//     return error if cannot unmarshal
	// validate the data
	//     return error if invalid
	// connect to the database
	//     return error if you can't connect
	// find database entry (using ID)
	//     return error if you cannot find it
	// update the database entry
	//     return confirmation that the entry was successfully updated or error if it was unsuccessful
	// ARCHITECT QUESTION: do you want to also return the entry, as you would with a read?
	//     Then you don't need to read immedietly after but this may mean making more DB requests than necessary.

	w.WriteHeader(200)
	w.Write([]byte("update stock"))
}
