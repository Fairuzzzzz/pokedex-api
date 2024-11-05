package team

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/Fairuzzzzz/pokedex-api/internal/models/team"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func Test_service_CreateTeam(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockTeamRepo := NewMockrepository(ctrlMock)

	type args struct {
		userID  uint
		request team.PokeTeamNameRequest
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
				request: team.PokeTeamNameRequest{
					TeamName: "team1",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockTeamRepo.EXPECT().Create(gomock.Any(), team.PokeTeam{
					UserID:    args.userID,
					TeamName:  args.request.TeamName,
					CreatedBy: "1",
					UpdatedBy: "1",
				}).Return(nil)
			},
		},
		{
			name: "failed",
			args: args{
				userID: 1,
				request: team.PokeTeamNameRequest{
					TeamName: "team1",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockTeamRepo.EXPECT().Create(gomock.Any(), team.PokeTeam{
					UserID:    args.userID,
					TeamName:  args.request.TeamName,
					CreatedBy: "1",
					UpdatedBy: "1",
				}).Return(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				repository: mockTeamRepo,
			}
			if err := s.CreateTeam(context.Background(), tt.args.userID, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("service.CreateTeam() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_ListTeam(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockTeamRepo := NewMockrepository(ctrlMock)

	now := time.Now()
	type args struct {
		userID uint
	}
	tests := []struct {
		name    string
		args    args
		want    []team.PokeTeam
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				userID: 1,
			},
			want: []team.PokeTeam{
				{
					Model: gorm.Model{
						ID:        1,
						CreatedAt: now,
						UpdatedAt: now,
					},
					UserID:    1,
					TeamName:  "team1",
					CreatedBy: "test@gmail.com",
					UpdatedBy: "test@gmail.com",
				},
				{
					Model: gorm.Model{
						ID:        2,
						CreatedAt: now,
						UpdatedAt: now,
					},
					UserID:    1,
					TeamName:  "team2",
					CreatedBy: "test@gmail.com",
					UpdatedBy: "test@gmail.com",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockTeamRepo.EXPECT().List(gomock.Any(), args.userID).Return([]team.PokeTeam{
					{
						Model: gorm.Model{
							ID:        1,
							CreatedAt: now,
							UpdatedAt: now,
						},
						UserID:    args.userID,
						TeamName:  "team1",
						CreatedBy: "test@gmail.com",
						UpdatedBy: "test@gmail.com",
					},
					{
						Model: gorm.Model{
							ID:        2,
							CreatedAt: now,
							UpdatedAt: now,
						},
						UserID:    args.userID,
						TeamName:  "team2",
						CreatedBy: "test@gmail.com",
						UpdatedBy: "test@gmail.com",
					},
				}, nil)
			},
		},
		{
			name: "failed",
			args: args{
				userID: 1,
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mockTeamRepo.EXPECT().List(gomock.Any(), args.userID).Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				repository: mockTeamRepo,
			}
			got, err := s.ListTeam(context.Background(), tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.ListTeam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.ListTeam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetTeam(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockTeamRepository := NewMockrepository(ctrlMock)

	now := time.Now()
	type args struct {
		request team.PokeTeamRequestByID
	}
	tests := []struct {
		name    string
		args    args
		want    *team.PokeTeam
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				request: team.PokeTeamRequestByID{
					ID:     1,
					UserID: 1,
				},
			},
			want: &team.PokeTeam{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: now,
					UpdatedAt: now,
				},
				UserID:    1,
				TeamName:  "team1",
				CreatedBy: "test@gmail.com",
				UpdatedBy: "test@gmail.com",
			},
			wantErr: false,
			mockFn: func(args args) {
				mockTeamRepository.EXPECT().Get(gomock.Any(), args.request.UserID, args.request.ID).Return(&team.PokeTeam{
					Model: gorm.Model{
						ID:        args.request.ID,
						CreatedAt: now,
						UpdatedAt: now,
					},
					UserID:    args.request.UserID,
					TeamName:  "team1",
					CreatedBy: "test@gmail.com",
					UpdatedBy: "test@gmail.com",
				}, nil)
			},
		},
		{
			name: "failed",
			args: args{
				request: team.PokeTeamRequestByID{
					ID:     1,
					UserID: 1,
				},
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mockTeamRepository.EXPECT().Get(gomock.Any(), args.request.UserID, args.request.ID).Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				repository: mockTeamRepository,
			}
			got, err := s.GetTeam(context.Background(), tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetTeam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetTeam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_DeleteTeam(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockTeamRepository := NewMockrepository(ctrlMock)

	type args struct {
		request team.PokeTeamRequestByID
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
				request: team.PokeTeamRequestByID{
					ID:     1,
					UserID: 1,
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockTeamRepository.EXPECT().Delete(gomock.Any(), args.request.UserID, args.request.ID).Return(nil)
			},
		},
		{
			name: "failed",
			args: args{
				request: team.PokeTeamRequestByID{
					ID:     1,
					UserID: 1,
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockTeamRepository.EXPECT().Delete(gomock.Any(), args.request.UserID, args.request.ID).Return(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				repository: mockTeamRepository,
			}
			if err := s.DeleteTeam(context.Background(), tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("service.DeleteTeam() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
