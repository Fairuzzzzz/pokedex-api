package pokemoninteam

import (
	"context"

	"github.com/Fairuzzzzz/pokedex-api/internal/models/pokemoninteam"
)

func (r *repository) Create(ctx context.Context, model pokemoninteam.PokemonNameInTeam) error {
	return r.db.Create(&model).Error
}

func (r *repository) DeletePokemon(ctx context.Context, teamID uint, pokemonID int) error {
	res := r.db.WithContext(ctx).Where("team_id = ? AND pokemon_id = ?", teamID, pokemonID).Delete(&pokemoninteam.PokemonNameInTeam{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *repository) List(ctx context.Context, teamID uint) ([]pokemoninteam.PokemonTeamResponse, error) {
	var pokeTeam []pokemoninteam.PokemonTeamResponse
	res := r.db.WithContext(ctx).Select("team_id, pokemon_id, pokemon_name").Where("team_id = ?", teamID).Find(&pokeTeam)
	if res.Error != nil {
		return nil, res.Error
	}
	return pokeTeam, nil
}

func (r *repository) GetPokemonCount(ctx context.Context, teamID uint) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&pokemoninteam.PokemonNameInTeam{}).Where("team_id = ?", teamID).Count(&count).Error
	return count, err
}
