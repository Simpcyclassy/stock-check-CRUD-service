package handlers

import (
	"net/http"
)

func CreateStockHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("404 Method not found."))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("create stock"))
}