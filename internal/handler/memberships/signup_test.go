package memberships

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Fairuzzzzz/pokedex-api/internal/models/memberships"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandler_SignUp(t *testing.T) {
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
	}{
		{
			name: "success",
			mockFn: func() {
				mockSvc.EXPECT().SignUp(memberships.SignUpRequest{
					Email:    "test@gmail.com",
					Username: "testusername",
					Password: "password",
				}).Return(nil)
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "failed",
			mockFn: func() {
				mockSvc.EXPECT().SignUp(memberships.SignUpRequest{
					Email:    "test@gmail.com",
					Username: "testusername",
					Password: "password",
				}).Return(errors.New("email or username exists"))
			},
			expectedCode: http.StatusBadRequest,
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

			endpoint := `/memberships/sign-up`
			model := memberships.SignUpRequest{
				Email:    "test@gmail.com",
				Username: "testusername",
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
		})
	}
}
