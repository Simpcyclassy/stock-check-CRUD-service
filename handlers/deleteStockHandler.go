package handlers

import (
	"net/http"
)

func DeleteStockHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete stock"))
}