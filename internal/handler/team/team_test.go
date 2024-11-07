package team

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Fairuzzzzz/pokedex-api/internal/models/team"
	"github.com/Fairuzzzzz/pokedex-api/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestHandler_CreateTeam(t *testing.T) {
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
				mockSvc.EXPECT().CreateTeam(gomock.Any(), uint(1), team.PokeTeamNameRequest{
					TeamName: "team1",
				}).Return(nil)
			},
		},
		{
			name:               "failed",
			expectedStatusCode: http.StatusBadRequest,
			wantErr:            true,
			mockFn: func() {
				mockSvc.EXPECT().CreateTeam(gomock.Any(), uint(1), team.PokeTeamNameRequest{
					TeamName: "team1",
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

			endpoint := `/team/create-team`

			model := team.PokeTeamNameRequest{
				TeamName: "team1",
			}

			val, err := json.Marshal(model)
			assert.NoError(t, err)

			// Mengubah val menjadi bytes
			body := bytes.NewReader(val)
			req, err := http.NewRequest(http.MethodPost, endpoint, body)
			assert.NoError(t, err)

			// Token JWT Authorization
			token, err := jwt.CreateToken(1, "", "")
			assert.NoError(t, err)
			req.Header.Set("Authorization", token)

			// Membuat ServeHTTP
			h.ServeHTTP(w, req)

			// Pengecekan status code
			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestHandler_ListTeam(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSvc := NewMockservice(ctrlMock)

	fixedTime := time.Date(2024, time.November, 7, 19, 17, 4, 308082349, time.Local)
	tests := []struct {
		name               string
		expectedStatusCode int
		expectedBody       *[]team.PokeTeam
		wantErr            bool
		mockFn             func()
	}{
		{
			name:               "success",
			expectedStatusCode: http.StatusOK,
			expectedBody: &[]team.PokeTeam{
				{
					Model: gorm.Model{
						ID:        1,
						CreatedAt: fixedTime,
						UpdatedAt: fixedTime,
					},
					UserID:    1,
					TeamName:  "team1",
					CreatedBy: "test@gmail.com",
					UpdatedBy: "test@gmail.com",
				},
				{
					Model: gorm.Model{
						ID:        2,
						CreatedAt: fixedTime,
						UpdatedAt: fixedTime,
					},
					UserID:    1,
					TeamName:  "team2",
					CreatedBy: "test@gmail.com",
					UpdatedBy: "test@gmail.com",
				},
			},
			wantErr: false,
			mockFn: func() {
				mockSvc.EXPECT().ListTeam(gomock.Any(), uint(1)).Return([]team.PokeTeam{
					{
						Model: gorm.Model{
							ID:        1,
							CreatedAt: fixedTime,
							UpdatedAt: fixedTime,
						},
						UserID:    1,
						TeamName:  "team1",
						CreatedBy: "test@gmail.com",
						UpdatedBy: "test@gmail.com",
					},
					{
						Model: gorm.Model{
							ID:        2,
							CreatedAt: fixedTime,
							UpdatedAt: fixedTime,
						},
						UserID:    1,
						TeamName:  "team2",
						CreatedBy: "test@gmail.com",
						UpdatedBy: "test@gmail.com",
					},
				}, nil)
			},
		},
		{
			name:               "failed",
			expectedStatusCode: http.StatusBadRequest,
			expectedBody: &[]team.PokeTeam{
				{
					Model: gorm.Model{
						ID:        1,
						CreatedAt: fixedTime,
						UpdatedAt: fixedTime,
					},
					UserID:    1,
					TeamName:  "team1",
					CreatedBy: "test@gmail.com",
					UpdatedBy: "test@gmail.com",
				},
				{
					Model: gorm.Model{
						ID:        2,
						CreatedAt: fixedTime,
						UpdatedAt: fixedTime,
					},
					UserID:    1,
					TeamName:  "team2",
					CreatedBy: "test@gmail.com",
					UpdatedBy: "test@gmail.com",
				},
			},
			wantErr: true,
			mockFn: func() {
				mockSvc.EXPECT().ListTeam(gomock.Any(), uint(1)).Return(nil, assert.AnError)
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

			endpoint := `/team/list-team`

			req, err := http.NewRequest(http.MethodGet, endpoint, nil)
			assert.NoError(t, err)

			// Token JWT Authorization
			token, err := jwt.CreateToken(1, "", "")
			assert.NoError(t, err)
			req.Header.Set("Authorization", token)

			// Membuat ServeHTTP
			h.ServeHTTP(w, req)

			// Pengecekan status code
			assert.Equal(t, tt.expectedStatusCode, w.Code)

			if !tt.wantErr {
				res := w.Result()
				defer res.Body.Close()

				response := []team.PokeTeam{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedBody, &response)
			}
		})
	}
}

func TestHandler_GetTeam(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSvc := NewMockservice(ctrlMock)

	fixedTime := time.Date(2024, time.November, 7, 19, 17, 4, 308082349, time.Local)
	tests := []struct {
		name               string
		expectedStatusCode int
		expectedBody       *team.PokeTeam
		wantErr            bool
		mockFn             func()
	}{
		{
			name:               "success",
			expectedStatusCode: http.StatusOK,
			expectedBody: &team.PokeTeam{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: fixedTime,
					UpdatedAt: fixedTime,
				},
				UserID:    1,
				TeamName:  "team1",
				CreatedBy: "test@gmail.com",
				UpdatedBy: "test@gmail.com",
			},
			wantErr: false,
			mockFn: func() {
				mockSvc.EXPECT().GetTeam(gomock.Any(), team.PokeTeamRequestByID{
					ID:     1,
					UserID: 1,
				}).Return(&team.PokeTeam{
					Model: gorm.Model{
						ID:        1,
						CreatedAt: fixedTime,
						UpdatedAt: fixedTime,
					},
					UserID:    1,
					TeamName:  "team1",
					CreatedBy: "test@gmail.com",
					UpdatedBy: "test@gmail.com",
				}, nil)
			},
		},
		{
			name:               "failed",
			expectedStatusCode: http.StatusBadRequest,
			expectedBody: &team.PokeTeam{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: fixedTime,
					UpdatedAt: fixedTime,
				},
				UserID:    1,
				TeamName:  "team1",
				CreatedBy: "test@gmail.com",
				UpdatedBy: "test@gmail.com",
			},
			wantErr: true,
			mockFn: func() {
				mockSvc.EXPECT().GetTeam(gomock.Any(), team.PokeTeamRequestByID{
					ID:     1,
					UserID: 1,
				}).Return(nil, assert.AnError)
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

			endpoint := `/team/get-team`

			model := team.PokeTeamRequestByID{
				ID:     1,
				UserID: 1,
			}

			val, err := json.Marshal(model)
			assert.NoError(t, err)

			// Mengubah val menjadi bytes
			body := bytes.NewReader(val)
			req, err := http.NewRequest(http.MethodGet, endpoint, body)
			assert.NoError(t, err)

			// Token JWT Authorization
			token, err := jwt.CreateToken(1, "", "")
			assert.NoError(t, err)
			req.Header.Set("Authorization", token)

			// Membuat ServeHTTP
			h.ServeHTTP(w, req)

			// Pengecekan status code
			assert.Equal(t, tt.expectedStatusCode, w.Code)

			if !tt.wantErr {
				res := w.Result()
				defer res.Body.Close()

				response := team.PokeTeam{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedBody, &response)
			}
		})
	}
}

func TestHandler_DeleteTeam(t *testing.T) {
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
				mockSvc.EXPECT().DeleteTeam(gomock.Any(), team.PokeTeamRequestByID{
					ID:     1,
					UserID: 1,
				}).Return(nil)
			},
		},
		{
			name:               "failed",
			expectedStatusCode: http.StatusBadRequest,
			wantErr:            true,
			mockFn: func() {
				mockSvc.EXPECT().DeleteTeam(gomock.Any(), team.PokeTeamRequestByID{
					ID:     1,
					UserID: 1,
				}).Return(assert.AnError)
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

			endpoint := `/team/delete-team`

			model := team.PokeTeamRequestByID{
				ID:     1,
				UserID: 1,
			}

			val, err := json.Marshal(model)
			assert.NoError(t, err)

			// Mengubah val menjadi bytes
			body := bytes.NewReader(val)
			req, err := http.NewRequest(http.MethodPost, endpoint, body)
			assert.NoError(t, err)

			// Token JWT Authorization
			token, err := jwt.CreateToken(1, "", "")
			assert.NoError(t, err)
			req.Header.Set("Authorization", token)

			// Membuat ServeHTTP
			h.ServeHTTP(w, req)

			// Pengecekan status code
			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}
