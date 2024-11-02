package poke

import (
	"context"

	"github.com/Fairuzzzzz/pokedex-api/internal/models/poke"
	"github.com/rs/zerolog/log"
)

func (s *service) PokemonSearch(ctx context.Context, name string) (*poke.SearchPokemonResponse, error) {
	pokeDetails, err := s.pokemonOutbound.SearchPokemon(ctx, name)
	if err != nil {
		log.Error().Err(err).Msg("error search pokemon to pokeapi")
		return nil, err
	}

	abilities := make([]poke.PokemonAbilities, len(pokeDetails.Abilities))
	for i, ability := range pokeDetails.Abilities {
		abilities[i] = poke.PokemonAbilities{
			IsHidden: ability.IsHidden,
			Slot:     ability.Slot,
			Ability: poke.AbilityDetails{
				Name: ability.Ability.Name,
			},
		}
	}

	moves := make([]poke.PokemonMoves, len(pokeDetails.Moves))
	for i, move := range pokeDetails.Moves {
		moves[i] = poke.PokemonMoves{
			Move: poke.MoveDetails{
				Name: move.Move.Name,
			},
		}
	}

	stats := make([]poke.PokemonStats, len(pokeDetails.Stats))
	for i, stat := range pokeDetails.Stats {
		stats[i] = poke.PokemonStats{
			BaseStat: stat.BaseStat,
			Effort:   stat.Effort,
			Stat: poke.StatDetails{
				Name: stat.Stat.Name,
			},
		}
	}

	species := poke.PokemonSpecies{
		Name: pokeDetails.Species.Name,
	}

	types := make([]poke.PokemonType, len(pokeDetails.Types))
	for i, t := range pokeDetails.Types {
		types[i] = poke.PokemonType{
			Slot: t.Slot,
			Type: poke.TypeDetails{
				Name: t.Type.Name,
			},
		}
	}

	response := &poke.SearchPokemonResponse{
		ID:             pokeDetails.ID,
		Name:           pokeDetails.Name,
		BaseExperience: pokeDetails.BaseExperience,
		Height:         pokeDetails.Height,
		Weight:         pokeDetails.Weight,
		Order:          pokeDetails.Order,
		Abilities:      abilities,
		Moves:          moves,
		Species:        species,
		Stats:          stats,
		Types:          types,
	}

	return response, nil
}
