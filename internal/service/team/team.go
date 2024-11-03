package team

import (
	"context"
	"errors"
	"fmt"

	"github.com/Fairuzzzzz/pokedex-api/internal/models/team"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
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

	if err != nil && err != gorm.ErrRecordNotFound {
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

// TODO Get Team and Delete Team
