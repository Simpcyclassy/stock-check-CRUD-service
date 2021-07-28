package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/simpcyclassy/stock-check-CRUD-service/handlers"
)
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/create", handlers.CreateStockHandler)
	r.HandleFunc("/read", handlers.ReadStockHandler)
	r.HandleFunc("/update", handlers.UpdateStockHandler)
	r.HandleFunc("/delete", handlers.DeleteStockHandler)

	http.Handle("/", r)
	log.Info().Msg("Starting server at port 8009")
	err := http.ListenAndServe(":8009", nil)
	if err != nil {
		log.Fatal().Err(err).Msg("Could not start server")
	}
}