package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Philo-Ultra/poke-list-api/pkg/models"
)

const (
	pokemonListEndpoint    = "https://pokeapi.co/api/v2/pokemon?limit=1000&offset=0"
	pokemonDetailsEndpoint = "https://pokeapi.co/api/v2/pokemon/%d"
)

func GetAllPokemon() ([]models.Pokemon, error) {
	// Create the request
	req, err := http.NewRequest(http.MethodGet, pokemonListEndpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	// Make the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Check the status in the response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: unexpected status code: %d", resp.StatusCode)
	}

	// Read the response body in
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var pokemonListResponse models.PokemonListResponse
	if err := json.Unmarshal(body, &pokemonListResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	// Extract the pokedex IDs from the results
	pokemonList := []models.Pokemon{}
	for _, p := range pokemonListResponse.Results {
		pokedexID, err := p.PokedexID()
		if err != nil {
			return nil, fmt.Errorf("error parsing pokedex ID: %w", err)
		}
		pokemonList = append(pokemonList, models.Pokemon{
			PokedexID: pokedexID,
			Name:      p.Name,
			URL:       p.URL,
		})
	}

	return pokemonList, nil
}

func GetPokemonDetails(pokedexID int) (models.PokemonDetails, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(pokemonDetailsEndpoint, pokedexID), nil)
	if err != nil {
		return models.PokemonDetails{}, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return models.PokemonDetails{}, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.PokemonDetails{}, fmt.Errorf("error: unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.PokemonDetails{}, fmt.Errorf("error reading response body: %w", err)
	}

	var pokemonDetails models.PokemonDetails
	if err := json.Unmarshal(body, &pokemonDetails); err != nil {
		return models.PokemonDetails{}, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return pokemonDetails, nil
}
