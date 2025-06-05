package service

import (
	"database/sql"
	"fmt"

	"github.com/Philo-Ultra/poke-list-api/pkg/models"
	"github.com/Philo-Ultra/poke-list-api/pkg/pokeapi"
	_ "github.com/lib/pq"
)

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{db: db}
}

func (s *Service) MyPokemon() ([]models.MyPokemon, error) {
	fmt.Println("HERE")
	query := `
		SELECT id, pokedex_id, nickname, created_at, deleted_at
		FROM my_pokemon;`

	myPokemon := []models.MyPokemon{}
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying database: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var mp models.MyPokemon
		err := rows.Scan(&mp.ID, &mp.PokedexID, &mp.Nickname, &mp.CreatedAt, &mp.DeletedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		myPokemon = append(myPokemon, mp)
	}

  for _, p := range myPokemon {
		fmt.Println("Getting pokemon details for", p.PokedexID)
		pokemonDetails, err := pokeapi.GetPokemonDetails(p.PokedexID)
		if err != nil {
			return nil, fmt.Errorf("error getting pokemon details: %w", err)
		}
		fmt.Println(pokemonDetails)
	}

	return myPokemon, nil
}
