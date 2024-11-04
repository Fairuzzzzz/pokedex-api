package team

import (
	"context"
	"errors"
	"fmt"

	"github.com/Fairuzzzzz/pokedex-api/internal/models/team"
	"github.com/rs/zerolog/log"
)

func (s *service) CreateTeam(ctx context.Context, userID uint, request team.PokeTeamNameRequest) error {
	if request.TeamName == "" {
		return errors.New("team name is required")
	}

	model := team.PokeTeam{
		UserID:    userID,
		TeamName:  request.TeamName,
		CreatedBy: fmt.Sprintf("%d", userID),
		UpdatedBy: fmt.Sprintf("%d", userID),
	}

	err := s.repository.Create(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("error create team record to database")
		return err
	}
	return nil
}

func (s *service) ListTeam(ctx context.Context, userID uint) ([]team.PokeTeam, error) {
	listTeam, err := s.repository.List(ctx, userID)
	if err != nil {
		log.Error().Err(err).Msg("error getting list team from database")
		return nil, err
	}

	return listTeam, nil
}

func (s *service) GetTeam(ctx context.Context, request team.PokeTeamRequestByID) (*team.PokeTeam, error) {
	pokeTeam, err := s.repository.Get(ctx, request.UserID, request.ID)
	if err != nil {
		log.Error().Err(err).Msg("error getting team from database")
		return nil, err
	}

	return pokeTeam, nil
}

func (s *service) DeleteTeam(ctx context.Context, request team.PokeTeamRequestByID) error {
	err := s.repository.Delete(ctx, request.UserID, request.ID)
	if err != nil {
		log.Error().Err(err).Msg("error delete team from database")
		return err
	}

	return nil
}
