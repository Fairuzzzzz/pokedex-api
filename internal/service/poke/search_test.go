package poke

import (
	"context"
	"reflect"
	"testing"

	"github.com/Fairuzzzzz/pokedex-api/internal/models/poke"
	pokeRepo "github.com/Fairuzzzzz/pokedex-api/internal/repository/poke"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_service_PokemonSearch(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockPokemonOutbound := NewMockpokemonOutbound(ctrlMock)
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    *poke.SearchPokemonResponse
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				name: "bulbasaur",
			},
			want: &poke.SearchPokemonResponse{
				ID:             1,
				Name:           "bulbasaur",
				BaseExperience: 64,
				Height:         7,
				Weight:         69,
				Order:          1,
				Abilities: []poke.PokemonAbilities{
					{
						IsHidden: false,
						Slot:     1,
						Ability: poke.AbilityDetails{
							Name: "overgrow",
						},
					},
					{
						IsHidden: true,
						Slot:     3,
						Ability: poke.AbilityDetails{
							Name: "chlorophyll",
						},
					},
				},
				Moves: []poke.PokemonMoves{
					{
						Move: poke.MoveDetails{
							Name: "razor-wind",
						},
					},
					{
						Move: poke.MoveDetails{
							Name: "swords-dance",
						},
					},
				},
				Species: poke.PokemonSpecies{
					Name: "bulbasaur",
				},
				Stats: []poke.PokemonStats{
					{
						BaseStat: 45,
						Effort:   0,
						Stat: poke.StatDetails{
							Name: "hp",
						},
					},
					{
						BaseStat: 49,
						Effort:   0,
						Stat: poke.StatDetails{
							Name: "attack",
						},
					},
					{
						BaseStat: 49,
						Effort:   0,
						Stat: poke.StatDetails{
							Name: "defense",
						},
					},
					{
						BaseStat: 65,
						Effort:   1,
						Stat: poke.StatDetails{
							Name: "special-attack",
						},
					},
					{
						BaseStat: 65,
						Effort:   0,
						Stat: poke.StatDetails{
							Name: "special-defense",
						},
					},
					{
						BaseStat: 45,
						Effort:   0,
						Stat: poke.StatDetails{
							Name: "speed",
						},
					},
				},
				Types: []poke.PokemonType{
					{
						Slot: 1,
						Type: poke.TypeDetails{
							Name: "grass",
						},
					},
					{
						Slot: 2,
						Type: poke.TypeDetails{
							Name: "poison",
						},
					},
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockPokemonOutbound.EXPECT().SearchPokemon(gomock.Any(), args.name).Return(&pokeRepo.Pokemon{
					ID:             1,
					Name:           "bulbasaur",
					BaseExperience: 64,
					Height:         7,
					Weight:         69,
					IsDefault:      true,
					Order:          1,
					Abilities: []pokeRepo.PokemonAbilities{
						{
							IsHidden: false,
							Slot:     1,
							Ability: struct {
								Name string "json:\"name\""
							}{
								Name: "overgrow",
							},
						},
						{
							IsHidden: true,
							Slot:     3,
							Ability: struct {
								Name string "json:\"name\""
							}{
								Name: "chlorophyll",
							},
						},
					},
					Moves: []pokeRepo.PokemonMoves{
						{
							Move: pokeRepo.MoveDetails{
								Name: "razor-wind",
							},
						},
						{
							Move: pokeRepo.MoveDetails{
								Name: "swords-dance",
							},
						},
					},
					Species: pokeRepo.PokemonSpecies{
						Name: "bulbasaur",
					},
					Stats: []pokeRepo.PokemonStats{
						{
							BaseStat: 45,
							Effort:   0,
							Stat: pokeRepo.Stat{
								Name: "hp",
							},
						},
						{
							BaseStat: 49,
							Effort:   0,
							Stat: pokeRepo.Stat{
								Name: "attack",
							},
						},
						{
							BaseStat: 49,
							Effort:   0,
							Stat: pokeRepo.Stat{
								Name: "defense",
							},
						},
						{
							BaseStat: 65,
							Effort:   1,
							Stat: pokeRepo.Stat{
								Name: "special-attack",
							},
						},
						{
							BaseStat: 65,
							Effort:   0,
							Stat: pokeRepo.Stat{
								Name: "special-defense",
							},
						},
						{
							BaseStat: 45,
							Effort:   0,
							Stat: pokeRepo.Stat{
								Name: "speed",
							},
						},
					},
					Types: []pokeRepo.PokemonType{
						{
							Slot: 1,
							Type: pokeRepo.Type{
								Name: "grass",
							},
						},
						{
							Slot: 2,
							Type: pokeRepo.Type{
								Name: "poison",
							},
						},
					},
				}, nil)
			},
		},
		{
			name: "failed",
			args: args{
				name: "bulbasaur",
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mockPokemonOutbound.EXPECT().SearchPokemon(gomock.Any(), args.name).Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				pokemonOutbound: mockPokemonOutbound,
			}
			got, err := s.PokemonSearch(context.Background(), tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.PokemonSearch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.PokemonSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
