package pokemoninteam

import (
	"github.com/Fairuzzzzz/pokedex-api/internal/models/team"
	"gorm.io/gorm"
)

type (
	PokemonNameInTeam struct {
		gorm.Model
		TeamID      uint          `gorm:"not null;foreignKey:ID;references:team.PokeTeam"`
		Team        team.PokeTeam `gorm:"foreignKey:TeamID"`
		PokemonID   int           `gorm:"not null"`
		PokemonName string        `gorm:"not null"`
		CreatedBy   string        `gorm:"not null"`
		UpdatedBy   string        `gorm:"not null"`
	}
)

type (
	PokemonRequestWithName struct {
		TeamID      uint   `json:"teamID"`
		PokemonName string `json:"pokemonName"`
	}

	PokemonRequestWithID struct {
		TeamID    uint `json:"teamID"`
		PokemonID int  `json:"pokemonID"`
	}
)

type (
	PokemonTeamResponse struct {
		TeamID      uint   `json:"teamID"`
		PokemonID   int    `json:"pokemonID"`
		PokemonName string `json:"pokemonName"`
	}
)
