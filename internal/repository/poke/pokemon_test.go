package poke

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/Fairuzzzzz/pokedex-api/internal/configs"
	"github.com/Fairuzzzzz/pokedex-api/pkg/httpclient"
	"go.uber.org/mock/gomock"
)

func Test_outbound_SearchPokemon(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockHTTPClient := httpclient.NewMockHTTPClient(ctrlMock)

	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    *Pokemon
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				name: "bulbasaur",
			},
			want: &Pokemon{
				ID:             1,
				Name:           "bulbasaur",
				BaseExperience: 64,
				Height:         7,
				Weight:         69,
				IsDefault:      true,
				Order:          1,
				Abilities: []PokemonAbilities{
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
				Moves: []PokemonMoves{
					{
						Move: MoveDetails{
							Name: "razor-wind",
						},
					},
					{
						Move: MoveDetails{
							Name: "swords-dance",
						},
					},
				},
				Species: PokemonSpecies{
					Name: "bulbasaur",
				},
				Stats: []PokemonStats{
					{
						BaseStat: 45,
						Effort:   0,
						Stat: Stat{
							Name: "hp",
						},
					},
					{
						BaseStat: 49,
						Effort:   0,
						Stat: Stat{
							Name: "attack",
						},
					},
					{
						BaseStat: 49,
						Effort:   0,
						Stat: Stat{
							Name: "defense",
						},
					},
					{
						BaseStat: 65,
						Effort:   1,
						Stat: Stat{
							Name: "special-attack",
						},
					},
					{
						BaseStat: 65,
						Effort:   0,
						Stat: Stat{
							Name: "special-defense",
						},
					},
					{
						BaseStat: 45,
						Effort:   0,
						Stat: Stat{
							Name: "speed",
						},
					},
				},
				Types: []PokemonType{
					{
						Slot: 1,
						Type: Type{
							Name: "grass",
						},
					},
					{
						Slot: 2,
						Type: Type{
							Name: "poison",
						},
					},
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				urlPath := "https://pokeapi.co/api/v2/pokemon/bulbasaur"
				req, _ := http.NewRequest(http.MethodGet, urlPath, nil)
				mockHTTPClient.EXPECT().Do(req).Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewReader([]byte(strings.TrimSpace(response)))),
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
				urlPath := "https://pokeapi.co/api/v2/pokemon/bulbasaur"
				req, _ := http.NewRequest(http.MethodGet, urlPath, nil)
				mockHTTPClient.EXPECT().Do(req).Return(&http.Response{
					StatusCode: http.StatusInternalServerError,
					Body:       io.NopCloser(bytes.NewBufferString("Internal Server Error")),
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			o := &outbound{
				cfg:    &configs.Config{},
				client: mockHTTPClient,
			}
			got, err := o.SearchPokemon(context.Background(), tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("outbound.SearchPokemon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("outbound.SearchPokemon() = %v, want %v", got, tt.want)
			}
		})
	}
}
