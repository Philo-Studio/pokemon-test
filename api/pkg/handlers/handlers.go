package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Philo-Ultra/poke-list-api/pkg/pokeapi"
	"github.com/Philo-Ultra/poke-list-api/pkg/service"
)

type Handler struct {
	svc service.Service
}

func NewHandler(svc service.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) MyPokemon(w http.ResponseWriter, r *http.Request) {
	myPokemon, err := h.svc.MyPokemon()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(myPokemon)
}

func (h *Handler) AllPokemon(w http.ResponseWriter, r *http.Request) {
	pokemon, err := pokeapi.GetAllPokemon()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pokemon)
}
