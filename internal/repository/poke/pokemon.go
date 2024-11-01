package poke

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Pokemon struct {
	ID             int                `json:"id"`
	Name           string             `json:"name"`
	BaseExperience int                `json:"base_experience"`
	Height         int                `json:"height"`
	IsDefault      int                `json:"is_default"`
	Order          int                `json:"order"`
	Weight         int                `json:"weight"`
	Abilities      []PokemonAbilities `json:"abilities"`
	Moves          []PokemonMoves     `json:"moves"`
	Species        PokemonSpecies     `json:"species"`
	Stats          []PokemonStats     `json:"stats"`
	Types          []PokemonType      `json:"types"`
}

type PokemonAbilities struct {
	IsHidden bool `json:"is_hidden"`
	Slot     int  `json:"slot"`
	Ability  struct {
		Name string `json"name"`
	} `json:"ability"`
}

type PokemonForms struct {
	Name     string `json:"name"`
	FormName string `json:"form_name"`
}

type PokemonMoves struct {
	Move MoveDetails `json:"move"`
}

type MoveDetails struct {
	Name string `json:"name"`
}

type PokemonSpecies struct {
	Name          string            `json:"name"`
	GenderRate    int               `json:"gender_rate"`
	CaptureRate   int               `json:"capture_rate"`
	BaseHappiness int               `json:"base_happiness"`
	IsBaby        bool              `json:"is_baby"`
	IsLegendary   bool              `json:"is_legendary"`
	IsMythical    bool              `json:"is_mythical"`
	Color         PokemonColor      `json:"color"`
	Habitat       PokemonHabitat    `json:"habitat"`
	Generation    PokemonGeneration `json:"generation"`
}

type PokemonColor struct {
	Name string `json:"name"`
}

type PokemonHabitat struct {
	Name string `json:"name"`
}

type PokemonGeneration struct {
	Name string `json:"name"`
}

type PokemonStats struct {
	Stat     Stat `json:"stat"`
	Effort   int  `json:"effort"`
	BaseStat int  `json:"base_stat"`
}

type Stat struct {
	Name string `json:"name"`
}

type PokemonType struct {
	Slot int  `json:"slot"`
	Type Type `json:"type"`
}

type Type struct {
	Name string `json:"name"`
}

func (o *outbound) SearchPokemon(ctx context.Context, name string) (*Pokemon, error) {
	urlPath := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v", name)
	req, err := http.NewRequest(http.MethodGet, urlPath, nil)
	if err != nil {
		log.Error().Err(err).Msg("error create search request for pokeapi")
		return nil, err
	}

	resp, err := o.client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("error execute search request for pokeapi")
		return nil, err
	}
	defer resp.Body.Close()

	// Unmarshal response
	var response Pokemon
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Error().Err(err).Msg("error unmarshal response from pokeapi")
		return nil, err
	}

	return &response, nil
}
