package handlers

import (
	"net/http"
)

func UpdateStockHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("404 Method not found."))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("update stock"))
}