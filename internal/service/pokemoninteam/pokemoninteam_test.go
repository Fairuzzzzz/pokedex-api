package pokemoninteam

import (
	"context"
	"reflect"
	"testing"

	"github.com/Fairuzzzzz/pokedex-api/internal/models/pokemoninteam"
	"github.com/Fairuzzzzz/pokedex-api/internal/repository/poke"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_service_AddPokemonToTeam(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockPokemonRepo := NewMockrepository(ctrlMock)

	mockPokemonOutbound := NewMockpokemonOutbound(ctrlMock)

	type args struct {
		userID  uint
		request pokemoninteam.PokemonRequestWithName
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				userID: 1,
				request: pokemoninteam.PokemonRequestWithName{
					TeamID:      1,
					PokemonName: "bulbasaur",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockPokemonOutbound.EXPECT().SearchPokemon(gomock.Any(), args.request.PokemonName).Return(&poke.Pokemon{
					ID:   1,
					Name: "bulbasaur",
				}, nil)
				mockPokemonRepo.EXPECT().GetPokemonCount(gomock.Any(), args.request.TeamID).Return(int64(2), nil)
				mockPokemonRepo.EXPECT().Create(gomock.Any(), pokemoninteam.PokemonNameInTeam{
					TeamID:      args.request.TeamID,
					PokemonID:   1,
					PokemonName: "bulbasaur",
					CreatedBy:   "1",
					UpdatedBy:   "1",
				}).Return(nil)
			},
		},
		{
			name: "failed: search pokemon",
			args: args{
				userID: 1,
				request: pokemoninteam.PokemonRequestWithName{
					TeamID:      1,
					PokemonName: "bulbasaur",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockPokemonOutbound.EXPECT().SearchPokemon(gomock.Any(), args.request.PokemonName).Return(nil, assert.AnError)
			},
		},
		{
			name: "failed: when team have 6 pokemon",
			args: args{
				userID: 1,
				request: pokemoninteam.PokemonRequestWithName{
					TeamID:      1,
					PokemonName: "bulbasaur",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockPokemonOutbound.EXPECT().SearchPokemon(gomock.Any(), args.request.PokemonName).Return(&poke.Pokemon{
					ID:   1,
					Name: "bulbasaur",
				}, nil)
				mockPokemonRepo.EXPECT().GetPokemonCount(gomock.Any(), args.request.TeamID).Return(int64(7), assert.AnError)
			},
		},
		{
			name: "failed: when create pokemon in team",
			args: args{
				userID: 1,
				request: pokemoninteam.PokemonRequestWithName{
					TeamID:      1,
					PokemonName: "bulbasaur",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockPokemonOutbound.EXPECT().SearchPokemon(gomock.Any(), args.request.PokemonName).Return(&poke.Pokemon{
					ID:   1,
					Name: "bulbasaur",
				}, nil)
				mockPokemonRepo.EXPECT().GetPokemonCount(gomock.Any(), args.request.TeamID).Return(int64(2), nil)
				mockPokemonRepo.EXPECT().Create(gomock.Any(), pokemoninteam.PokemonNameInTeam{
					TeamID:      args.request.TeamID,
					PokemonID:   1,
					PokemonName: "bulbasaur",
					CreatedBy:   "1",
					UpdatedBy:   "1",
				}).Return(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				repository:      mockPokemonRepo,
				pokemonOutbound: mockPokemonOutbound,
			}
			if err := s.AddPokemonToTeam(context.Background(), tt.args.userID, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("service.AddPokemonToTeam() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_DeletePokemon(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockPokemonRepo := NewMockrepository(ctrlMock)

	mockPokemonOutbound := NewMockpokemonOutbound(ctrlMock)

	type args struct {
		request pokemoninteam.PokemonRequestWithID
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				request: pokemoninteam.PokemonRequestWithID{
					TeamID:    1,
					PokemonID: 1,
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockPokemonRepo.EXPECT().DeletePokemon(gomock.Any(), args.request.TeamID, args.request.PokemonID).Return(nil)
			},
		},
		{
			name: "failed",
			args: args{
				request: pokemoninteam.PokemonRequestWithID{
					TeamID:    1,
					PokemonID: 1,
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockPokemonRepo.EXPECT().DeletePokemon(gomock.Any(), args.request.TeamID, args.request.PokemonID).Return(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				repository:      mockPokemonRepo,
				pokemonOutbound: mockPokemonOutbound,
			}
			if err := s.DeletePokemon(context.Background(), tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("service.DeletePokemon() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_ListPokemon(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockPokemonRepo := NewMockrepository(ctrlMock)

	mockPokemonOutbound := NewMockpokemonOutbound(ctrlMock)

	type args struct {
		teamID uint
	}
	tests := []struct {
		name    string
		args    args
		want    []pokemoninteam.PokemonTeamResponse
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				teamID: 1,
			},
			want: []pokemoninteam.PokemonTeamResponse{
				{
					TeamID:      1,
					PokemonID:   1,
					PokemonName: "bulbasaur",
				},
				{
					TeamID:      1,
					PokemonID:   132,
					PokemonName: "ditto",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockPokemonRepo.EXPECT().List(gomock.Any(), args.teamID).Return([]pokemoninteam.PokemonTeamResponse{
					{
						TeamID:      1,
						PokemonID:   1,
						PokemonName: "bulbasaur",
					},
					{
						TeamID:      1,
						PokemonID:   132,
						PokemonName: "ditto",
					},
				}, nil)
			},
		},
		{
			name: "failed",
			args: args{
				teamID: 1,
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mockPokemonRepo.EXPECT().List(gomock.Any(), args.teamID).Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				repository:      mockPokemonRepo,
				pokemonOutbound: mockPokemonOutbound,
			}
			got, err := s.ListPokemon(context.Background(), tt.args.teamID)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.ListPokemon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.ListPokemon() = %v, want %v", got, tt.want)
			}
		})
	}
}
