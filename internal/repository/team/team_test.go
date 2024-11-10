package team

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Fairuzzzzz/pokedex-api/internal/models/team"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_repository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	assert.NoError(t, err)

	now := time.Now()
	type args struct {
		model team.PokeTeam
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
				model: team.PokeTeam{
					Model: gorm.Model{
						CreatedAt: now,
						UpdatedAt: now,
					},
					UserID:    1,
					TeamName:  "team1",
					CreatedBy: "1",
					UpdatedBy: "1",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(`INSERT INTO "poke_teams" (.+) VALUES (.+)`).WithArgs(
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					args.model.UserID,
					args.model.TeamName,
					args.model.CreatedBy,
					args.model.UpdatedBy,
				).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(uint(1)))
				mock.ExpectCommit()
			},
		},
		{
			name: "failed",
			args: args{
				model: team.PokeTeam{
					Model: gorm.Model{
						CreatedAt: now,
						UpdatedAt: now,
					},
					UserID:    1,
					TeamName:  "team1",
					CreatedBy: "1",
					UpdatedBy: "1",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(`INSERT INTO "poke_teams" (.+) VALUES (.+)`).WithArgs(
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					args.model.UserID,
					args.model.TeamName,
					args.model.CreatedBy,
					args.model.UpdatedBy,
				).WillReturnError(assert.AnError)
				mock.ExpectRollback()
			},
		},
	}
	for _, tt := range tests {
		tt.mockFn(tt.args)
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: gormDB,
			}
			if err := r.Create(context.Background(), tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("repository.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func Test_repository_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	assert.NoError(t, err)

	now := time.Now()

	type args struct {
		model team.PokeTeam
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
				model: team.PokeTeam{
					Model: gorm.Model{
						ID:        123,
						CreatedAt: now,
						UpdatedAt: now,
					},
					UserID:    1,
					TeamName:  "team2",
					CreatedBy: "1",
					UpdatedBy: "1",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectExec(`UPDATE "poke_teams" SET (.+) WHERE (.+)`).WithArgs(
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					args.model.UserID,
					args.model.TeamName,
					args.model.CreatedBy,
					args.model.UpdatedBy,
					args.model.ID,
				).WillReturnResult(sqlmock.NewResult(123, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "failed",
			args: args{
				model: team.PokeTeam{
					Model: gorm.Model{
						ID:        123,
						CreatedAt: now,
						UpdatedAt: now,
					},
					UserID:    1,
					TeamName:  "team2",
					CreatedBy: "1",
					UpdatedBy: "1",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectExec(`UPDATE "poke_teams" SET (.+) WHERE (.+)`).WithArgs(
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					args.model.UserID,
					args.model.TeamName,
					args.model.CreatedBy,
					args.model.UpdatedBy,
					args.model.ID,
				).WillReturnError(assert.AnError)
				mock.ExpectRollback()
			},
		},
	}
	for _, tt := range tests {
		tt.mockFn(tt.args)
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: gormDB,
			}
			if err := r.Update(context.Background(), tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("repository.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func Test_repository_Get(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	assert.NoError(t, err)

	now := time.Now()
	type args struct {
		userID uint
		ID     uint
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
				userID: 1,
				ID:     1,
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
				mock.ExpectQuery(`SELECT \* FROM "poke_teams" .+`).WithArgs(args.userID, args.ID, 1).WillReturnRows(
					sqlmock.NewRows([]string{
						"id",
						"created_at",
						"updated_at",
						"user_id",
						"team_name",
						"created_by",
						"updated_by",
					}).AddRow(1, now, now, 1, "team1", "test@gmail.com", "test@gmail.com"),
				)
			},
		},
		{
			name: "failed",
			args: args{
				userID: 1,
				ID:     1,
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectQuery(`SELECT \* FROM "poke_teams" .+`).WithArgs(args.userID, args.ID, 1).WillReturnError(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		tt.mockFn(tt.args)
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: gormDB,
			}
			got, err := r.Get(context.Background(), tt.args.userID, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.Get() = %v, want %v", got, tt.want)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func Test_repository_List(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	assert.NoError(t, err)

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
				mock.ExpectQuery(`SELECT \* FROM "poke_teams" WHERE user_id = \$1 AND "poke_teams"."deleted_at" IS NULL ORDER BY id ASC`).
					WillReturnRows(sqlmock.NewRows([]string{
						"id",
						"created_at",
						"updated_at",
						"user_id",
						"team_name",
						"created_by",
						"updated_by",
					}).AddRow(1, now, now, 1, "team1", "test@gmail.com", "test@gmail.com").AddRow(2, now, now, 1, "team2", "test@gmail.com", "test@gmail.com"))
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
				mock.ExpectQuery(`SELECT \* FROM "poke_teams" WHERE user_id = \$1 AND "poke_teams"."deleted_at" IS NULL ORDER BY id ASC`).
					WillReturnError(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		tt.mockFn(tt.args)
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: gormDB,
			}
			got, err := r.List(context.Background(), tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.List() = %v, want %v", got, tt.want)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func Test_repository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	assert.NoError(t, err)

	type args struct {
		userID uint
		ID     uint
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
				ID:     123,
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectBegin()
				mock.ExpectExec(`UPDATE "poke_teams" SET "deleted_at"=\$1 WHERE \(user_id = \$2 AND id = \$3\) AND "poke_teams"."deleted_at" IS NULL`).
					WithArgs(sqlmock.AnyArg(), args.userID, args.ID).WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "failed",
			args: args{
				userID: 1,
				ID:     123,
			},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectBegin()
				mock.ExpectExec(`UPDATE "poke_teams" SET "deleted_at"=\$1 WHERE \(user_id = \$2 AND id = \$3\) AND "poke_teams"."deleted_at" IS NULL`).
					WithArgs(sqlmock.AnyArg(), args.userID, args.ID).WillReturnError(assert.AnError)
				mock.ExpectRollback()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &repository{
				db: gormDB,
			}
			if err := r.Delete(context.Background(), tt.args.userID, tt.args.ID); (err != nil) != tt.wantErr {
				t.Errorf("repository.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
