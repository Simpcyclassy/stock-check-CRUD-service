package handlers

import (
	"io/ioutil"
	"net/http"
)

func CreateStockHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
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
	// insert the new data
	//     return confirmation that the new entry was successful or error if it was unsuccessful
	// ARCHITECT QUESTION: do you want to also return the entry, as you would with a read?
	//     Then you don't need to read immedietly after but this may mean making more DB requests than necessary.
	w.WriteHeader(200)
	w.Write([]byte("create stock"))
}
