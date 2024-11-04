package team

import (
	"context"

	"github.com/Fairuzzzzz/pokedex-api/internal/models/team"
)

//go:generate mockgen -source=service.go -destination=service_mock_test.go -package=team
type repository interface {
	Create(ctx context.Context, model team.PokeTeam) error
	Update(ctx context.Context, model team.PokeTeam) error
	Get(ctx context.Context, userID uint, ID uint) (*team.PokeTeam, error)
	List(ctx context.Context, userID uint) ([]team.PokeTeam, error)
	Delete(ctx context.Context, userID uint, ID uint) error
}

type service struct {
	repository repository
}

func NewService(repository repository) *service {
	return &service{
		repository: repository,
	}
}
