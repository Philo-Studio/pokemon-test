package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type PokemonListResponse struct {
	Count    int                       `json:"count"`
	Next     *string                   `json:"next"`
	Previous *string                   `json:"previous"`
	Results  []PokemonListResponseItem `json:"results"`
}

type PokemonListResponseItemArray []PokemonListResponseItem

func (p *PokemonListResponseItemArray) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), p)
}

type PokemonListResponseItem struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (p *PokemonListResponseItem) PokedexID() (int, error) {
	parts := strings.Split(p.URL, "/")
	id, err := strconv.Atoi(parts[len(parts)-2])
	if err != nil {
		return 0, fmt.Errorf("error parsing pokedex ID: %w", err)
	}
	return id, nil
}

type Pokemon struct {
	PokedexID int    `json:"pokedex_id"`
	Name      string `json:"name"`
	URL       string `json:"url"`
}

//	{
//	  // "abilities": [...],
//	  // "base_experience": 194,
//	  // "cries": {...},
//	  // "forms": [...],
//	  // "game_indices": [...],
//	  "height": 19,
//	  // "held_items": [...],
//	  "id": 59,
//	  "is_default": true,
//	  "location_area_encounters": "https://pokeapi.co/api/v2/pokemon/59/encounters",
//	  // "moves": [...],
//	  "name": "arcanine",
//	  "order": 98,
//	  // "past_abilities": [...],
//	  // "past_types": [...],
//	  "species": { "name": "arcanine", "url": "https://pokeapi.co/api/v2/pokemon-species/59/" },
//	  "sprites": {
//	    "front_default": "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/59.png"
//	    // ...
//	  },
//	  // "stats": [...],
//	  // "types": [...],
//	  "weight": 1550
//	}
type PokemonDetails struct {
	ID      int `json:"id"`
	Species struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Height  int `json:"height"`
	Weight  int `json:"weight"`
	Sprites struct {
		FrontDefault string `json:"front_default"`
	} `json:"sprites"`
	// Theres a lot more fields here, but we're not using them for now
}
