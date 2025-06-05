package routes

import (
	"net/http"

	"github.com/Philo-Ultra/poke-list-api/pkg/handlers"
	"github.com/gorilla/mux"
)

func AddRoutes(r *mux.Router, handler *handlers.Handler) {
	r.HandleFunc("/pokemon/mine", http.HandlerFunc(handler.MyPokemon)).Methods("GET")
	r.HandleFunc("/pokemon/all", http.HandlerFunc(handler.AllPokemon)).Methods("GET")
}
