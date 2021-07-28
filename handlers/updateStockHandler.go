package handlers

import (
	"net/http"
)

func UpdateStockHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update stock"))
}