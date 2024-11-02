package poke

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Fairuzzzzz/pokedex-api/internal/models/poke"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandler_Search(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSvc := NewMockservice(ctrlMock)

	tests := []struct {
		name               string
		expectedStatusCode int
		expectedBody       poke.SearchPokemonResponse
		wantErr            bool
		mockFn             func()
	}{
		{
			name:               "success",
			expectedStatusCode: http.StatusOK,
			expectedBody: poke.SearchPokemonResponse{
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
			mockFn: func() {
				mockSvc.EXPECT().PokemonSearch(gomock.Any(), "bulbasaur").Return(&poke.SearchPokemonResponse{
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
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			api := gin.New()
			h := &Handler{
				Engine:  api,
				service: mockSvc,
			}
			h.RegisterRoute()
			w := httptest.NewRecorder()

			endpoint := `/pokemon/bulbasaur`
			req, err := http.NewRequest(http.MethodGet, endpoint, nil)
			assert.NoError(t, err)

			h.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)

			if !tt.wantErr {
				res := w.Result()
				defer res.Body.Close()

				response := poke.SearchPokemonResponse{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedBody, response)
			}
		})
	}
}
