package memberships

import (
	"fmt"
	"testing"

	"github.com/Fairuzzzzz/pokedex-api/internal/configs"
	"github.com/Fairuzzzzz/pokedex-api/internal/models/memberships"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func Test_service_Login(t *testing.T) {
	ctrlMock := gomock.NewController(t)

	defer ctrlMock.Finish()

	mockRepo := NewMockrepository(ctrlMock)

	type args struct {
		request memberships.LoginRequest
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				request: memberships.LoginRequest{
					Email:    "test@gmail.com",
					Password: "password",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, "", uint(0)).Return(&memberships.User{
					Model: gorm.Model{
						ID: 1,
					},
					Email:    "test@gmail.com",
					Password: "$2a$10$ceTSPChiQ7GfGGIGDkpiA.DZsvivBT5KA498uuSsU1UTL8XxU7mnm",
					Username: "fairuz",
				}, nil)
			},
		},
		{
			name: "failed",
			args: args{
				request: memberships.LoginRequest{
					Email:    "test@gmail.com",
					Password: "password",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, "", uint(0)).Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		tt.mockFn(tt.args)
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				cfg: &configs.Config{
					Service: configs.Service{
						SecretJWT: "abc",
					},
				},
				repository: mockRepo,
			}
			got, err := s.Login(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				fmt.Printf("test case : %+v", tt.name)
				assert.NotEmpty(t, got)
			} else {
				assert.Empty(t, got)
			}
		})
	}
}
