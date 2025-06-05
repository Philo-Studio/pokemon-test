package models

import "time"

type MyPokemon struct {
	ID        string     `json:"id"`
	PokedexID int        `json:"pokedex_id"`
	Nickname  string     `json:"nickname"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
