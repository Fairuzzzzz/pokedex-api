package team

import (
	"context"

	"github.com/Fairuzzzzz/pokedex-api/internal/models/team"
)

func (r *repository) Create(ctx context.Context, model team.PokeTeam) error {
	return r.db.WithContext(ctx).Create(&model).Error
}

func (r *repository) Update(ctx context.Context, model team.PokeTeam) error {
	return r.db.Save(&model).Error
}

func (r *repository) Get(ctx context.Context, userID uint, ID uint) (*team.PokeTeam, error) {
	pokeTeam := team.PokeTeam{}
	res := r.db.Where("user_id = ? AND id = ?", userID, ID).First(&pokeTeam)
	if res.Error != nil {
		return nil, res.Error
	}
	return &pokeTeam, nil
}

func (r *repository) List(ctx context.Context, userID uint) ([]team.PokeTeam, error) {
	var pokeTeams []team.PokeTeam
	res := r.db.WithContext(ctx).Where("user_id = ?", userID).Order("id ASC").Find(&pokeTeams)
	if res.Error != nil {
		return nil, res.Error
	}
	return pokeTeams, nil
}

func (r *repository) Delete(ctx context.Context, userID uint, ID uint) error {
	res := r.db.WithContext(ctx).Where("user_id = ? AND id = ?", userID, ID).Delete(&team.PokeTeam{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}
