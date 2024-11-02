package poke

import (
	"context"

	"github.com/Fairuzzzzz/pokedex-api/internal/repository/poke"
)

//go:generate mockgen -source=service.go -destination=service_mock_test.go -package=poke
type pokemonOutbound interface {
	SearchPokemon(ctx context.Context, name string) (*poke.Pokemon, error)
}

type service struct {
	pokemonOutbound pokemonOutbound
}

func NewOutbound(pokemonOutbound pokemonOutbound) *service {
	return &service{
		pokemonOutbound: pokemonOutbound,
	}
}
