package pokemoninteam

import (
	"context"
	"errors"
	"fmt"

	"github.com/Fairuzzzzz/pokedex-api/internal/models/pokemoninteam"
	"github.com/rs/zerolog/log"
)

func (s *service) AddPokemonToTeam(ctx context.Context, userID uint, request pokemoninteam.PokemonRequestWithName) error {
	pokemon, err := s.pokemonOutbound.SearchPokemon(ctx, request.PokemonName)
	if err != nil {
		log.Error().Err(err).Msg("error searching pokemon from pokeapi")
		return err
	}

	// Get current pokemon count in team
	count, err := s.repository.GetPokemonCount(ctx, request.TeamID)
	if err != nil {
		log.Error().Err(err).Msg("error getting pokemon count in team")
		return err
	}

	// check if team already has 6 pokemon
	if count >= 6 {
		return errors.New("team already has maximum number of pokemon (6)")
	}

	model := pokemoninteam.PokemonNameInTeam{
		TeamID:      request.TeamID,
		PokemonID:   pokemon.ID,
		PokemonName: pokemon.Name,
		CreatedBy:   fmt.Sprintf("%d", userID),
		UpdatedBy:   fmt.Sprintf("%d", userID),
	}

	err = s.repository.Create(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("error create pokemon in team")
		return err
	}
	return nil
}

func (s *service) DeletePokemon(ctx context.Context, request pokemoninteam.PokemonRequestWithID) error {
	err := s.repository.DeletePokemon(ctx, request.TeamID, request.PokemonID)
	if err != nil {
		log.Error().Err(err).Msg("error deleting pokemon from team")
		return err
	}
	return nil
}

func (s *service) ListPokemon(ctx context.Context, teamID uint) ([]pokemoninteam.PokemonTeamResponse, error) {
	pokemons, err := s.repository.List(ctx, teamID)
	if err != nil {
		log.Error().Err(err).Msg("error getting list pokemon from team")
		return nil, err
	}
	return pokemons, nil
}
