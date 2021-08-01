package handlers

import (
	"net/http"
)

func DeleteStockHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Method not found."))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("delete stock"))
}