package pokemoninteam

import (
	"context"

	"github.com/Fairuzzzzz/pokedex-api/internal/models/pokemoninteam"
	"github.com/Fairuzzzzz/pokedex-api/internal/repository/poke"
)

//go:generate mockgen -source=service.go -destination=service_mock_test.go -package=pokemoninteam
type pokemonOutbound interface {
	SearchPokemon(ctx context.Context, name string) (*poke.Pokemon, error)
}

type repository interface {
	Create(ctx context.Context, model pokemoninteam.PokemonNameInTeam) error
	DeletePokemon(ctx context.Context, teamID uint, pokemonID int) error
	List(ctx context.Context, teamID uint) ([]pokemoninteam.PokemonTeamResponse, error)
	GetPokemonCount(ctx context.Context, teamID uint) (int64, error)
}

type service struct {
	repository      repository
	pokemonOutbound pokemonOutbound
}

func NewService(repository repository, pokemonOutbound pokemonOutbound) *service {
	return &service{
		repository:      repository,
		pokemonOutbound: pokemonOutbound,
	}
}
