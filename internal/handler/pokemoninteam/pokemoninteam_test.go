package pokemoninteam

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Fairuzzzzz/pokedex-api/internal/models/pokemoninteam"
	"github.com/Fairuzzzzz/pokedex-api/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandler_AddPokemon(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSvc := NewMockservice(ctrlMock)

	tests := []struct {
		name               string
		expectedStatusCode int
		wantErr            bool
		mockFn             func()
	}{
		{
			name:               "success",
			expectedStatusCode: http.StatusCreated,
			wantErr:            false,
			mockFn: func() {
				mockSvc.EXPECT().AddPokemonToTeam(gomock.Any(), uint(1), pokemoninteam.PokemonRequestWithName{
					TeamID:      1,
					PokemonName: "bulbasaur",
				}).Return(nil)
			},
		},
		{
			name:               "failed",
			expectedStatusCode: http.StatusBadRequest,
			wantErr:            true,
			mockFn: func() {
				mockSvc.EXPECT().AddPokemonToTeam(gomock.Any(), uint(1), pokemoninteam.PokemonRequestWithName{
					TeamID:      1,
					PokemonName: "bulbasaur",
				}).Return(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		api := gin.New()
		tt.mockFn()
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				Engine:  api,
				service: mockSvc,
			}
			h.RegisterRoute()
			w := httptest.NewRecorder()

			endpoint := `/pokemon-team/add`

			model := pokemoninteam.PokemonRequestWithName{
				TeamID:      1,
				PokemonName: "bulbasaur",
			}

			val, err := json.Marshal(model)
			assert.NoError(t, err)

			body := bytes.NewReader(val)
			req, err := http.NewRequest(http.MethodPost, endpoint, body)
			assert.NoError(t, err)

			// Token JWT Authorization
			token, err := jwt.CreateToken(1, "", "")
			assert.NoError(t, err)
			req.Header.Set("Authorization", token)

			h.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestHandler_DeletePokemon(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSvc := NewMockservice(ctrlMock)

	tests := []struct {
		name               string
		expectedStatusCode int
		wantErr            bool
		mockFn             func()
	}{
		{
			name:               "success",
			expectedStatusCode: http.StatusOK,
			wantErr:            false,
			mockFn: func() {
				mockSvc.EXPECT().DeletePokemon(gomock.Any(), pokemoninteam.PokemonRequestWithID{
					TeamID:    1,
					PokemonID: 1,
				}).Return(nil)
			},
		},
		{
			name:               "failed",
			expectedStatusCode: http.StatusBadRequest,
			wantErr:            true,
			mockFn: func() {
				mockSvc.EXPECT().DeletePokemon(gomock.Any(), pokemoninteam.PokemonRequestWithID{
					TeamID:    1,
					PokemonID: 1,
				}).Return(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := gin.New()
			tt.mockFn()
			h := &Handler{
				Engine:  api,
				service: mockSvc,
			}
			h.RegisterRoute()
			w := httptest.NewRecorder()

			endpoint := `/pokemon-team/remove`

			model := pokemoninteam.PokemonRequestWithID{
				TeamID:    1,
				PokemonID: 1,
			}

			val, err := json.Marshal(model)
			assert.NoError(t, err)

			body := bytes.NewReader(val)
			req, err := http.NewRequest(http.MethodPost, endpoint, body)
			assert.NoError(t, err)

			// Token JWT Authorization
			token, err := jwt.CreateToken(1, "", "")
			assert.NoError(t, err)
			req.Header.Set("Authorization", token)

			h.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestHandler_ListPokemon(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSvc := NewMockservice(ctrlMock)

	tests := []struct {
		name               string
		expectedStatusCode int
		expectedBody       []pokemoninteam.PokemonTeamResponse
		wantErr            bool
		mockFn             func()
	}{
		{
			name:               "success",
			expectedStatusCode: http.StatusOK,
			expectedBody: []pokemoninteam.PokemonTeamResponse{
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
			mockFn: func() {
				mockSvc.EXPECT().ListPokemon(gomock.Any(), uint(1)).Return([]pokemoninteam.PokemonTeamResponse{
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
			name:               "failed",
			expectedStatusCode: http.StatusBadRequest,
			expectedBody: []pokemoninteam.PokemonTeamResponse{
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
			wantErr: true,
			mockFn: func() {
				mockSvc.EXPECT().ListPokemon(gomock.Any(), uint(1)).Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := gin.New()
			tt.mockFn()
			h := &Handler{
				Engine:  api,
				service: mockSvc,
			}
			h.RegisterRoute()
			w := httptest.NewRecorder()

			endpoint := `/pokemon-team/list/1`
			req, err := http.NewRequest(http.MethodGet, endpoint, nil)
			assert.NoError(t, err)

			// Token JWT Authorization
			token, err := jwt.CreateToken(1, "", "")
			assert.NoError(t, err)
			req.Header.Set("Authorization", token)

			h.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)

			if !tt.wantErr {
				res := w.Result()
				defer res.Body.Close()

				response := []pokemoninteam.PokemonTeamResponse{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedBody, response)
			}
		})
	}
}
