package memberships

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Fairuzzzzz/pokedex-api/internal/models/memberships"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandler_Login(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSvc := NewMockservice(ctrlMock)

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name         string
		mockFn       func()
		expectedCode int
		expectedBody memberships.LoginResponse
		wantErr      bool
	}{
		{
			name: "success",
			mockFn: func() {
				mockSvc.EXPECT().Login(memberships.LoginRequest{
					Email:    "test@gmail.com",
					Password: "password",
				}).Return("accessToken", nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: memberships.LoginResponse{
				AccessToken: "accessToken",
			},
			wantErr: false,
		},
		{
			name: "failed",
			mockFn: func() {
				mockSvc.EXPECT().Login(memberships.LoginRequest{
					Email:    "test@gmail.com",
					Password: "password",
				}).Return("", assert.AnError)
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: memberships.LoginResponse{},
			wantErr:      true,
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

			endpoint := `/memberships/login`
			model := memberships.LoginRequest{
				Email:    "test@gmail.com",
				Password: "password",
			}

			val, err := json.Marshal(model)
			assert.NoError(t, err)

			// Mengubah val menjadi bytes
			body := bytes.NewReader(val)
			req, err := http.NewRequest(http.MethodPost, endpoint, body)
			assert.NoError(t, err)

			// Membuat ServeHTTP
			h.ServeHTTP(w, req)

			// Pengecekan status code
			assert.Equal(t, tt.expectedCode, w.Code)

			if !tt.wantErr {
				res := w.Result()
				defer res.Body.Close()

				response := memberships.LoginResponse{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedBody, response)
			}
		})
	}
}
