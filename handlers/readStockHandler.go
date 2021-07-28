package handlers

import (
	"net/http"
)

func ReadStockHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("read stock"))
}