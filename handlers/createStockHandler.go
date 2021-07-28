package handlers

import (
	"net/http"
)

func CreateStockHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create stock"))
}