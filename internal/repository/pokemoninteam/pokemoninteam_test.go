package pokemoninteam

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Fairuzzzzz/pokedex-api/internal/models/pokemoninteam"
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
		model pokemoninteam.PokemonNameInTeam
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
				model: pokemoninteam.PokemonNameInTeam{
					Model: gorm.Model{
						CreatedAt: now,
						UpdatedAt: now,
					},
					TeamID:      1,
					PokemonID:   1,
					PokemonName: "bulbasaur",
					CreatedBy:   "1",
					UpdatedBy:   "1",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(`INSERT INTO "pokemon_name_in_teams" (.+) VALUES (.+)`).WithArgs(
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					args.model.TeamID,
					args.model.PokemonID,
					args.model.PokemonName,
					args.model.CreatedBy,
					args.model.UpdatedBy,
				).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(uint(1)))
				mock.ExpectCommit()
			},
		},
		{
			name: "failed",
			args: args{
				model: pokemoninteam.PokemonNameInTeam{
					Model: gorm.Model{
						CreatedAt: now,
						UpdatedAt: now,
					},
					TeamID:      1,
					PokemonID:   1,
					PokemonName: "bulbasaur",
					CreatedBy:   "1",
					UpdatedBy:   "1",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(`INSERT INTO "pokemon_name_in_teams" (.+) VALUES (.+)`).WithArgs(
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					args.model.TeamID,
					args.model.PokemonID,
					args.model.PokemonName,
					args.model.CreatedBy,
					args.model.UpdatedBy,
				).WillReturnError(assert.AnError)
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
			if err := r.Create(context.Background(), tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("repository.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func Test_repository_DeletePokemon(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	assert.NoError(t, err)

	type args struct {
		teamID    uint
		pokemonID int
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
				teamID:    1,
				pokemonID: 1,
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectExec(`UPDATE "pokemon_name_in_teams" SET "deleted_at"=\$1 WHERE \(team_id = \$2 AND pokemon_id = \$3\) AND "pokemon_name_in_teams"."deleted_at" IS NULL`).WithArgs(
					sqlmock.AnyArg(),
					args.teamID,
					args.pokemonID,
				).WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "failed",
			args: args{
				teamID:    1,
				pokemonID: 1,
			},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectExec(`UPDATE "pokemon_name_in_teams" SET "deleted_at"=\$1 WHERE \(team_id = \$2 AND pokemon_id = \$3\) AND "pokemon_name_in_teams"."deleted_at" IS NULL`).WithArgs(
					sqlmock.AnyArg(),
					args.teamID,
					args.pokemonID,
				).WillReturnError(assert.AnError)
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
			if err := r.DeletePokemon(context.Background(), tt.args.teamID, tt.args.pokemonID); (err != nil) != tt.wantErr {
				t.Errorf("repository.DeletePokemon() error = %v, wantErr %v", err, tt.wantErr)
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
				mock.ExpectQuery(`SELECT team_id, pokemon_id, pokemon_name FROM "pokemon_name_in_teams" WHERE team_id = \$1`).WithArgs(args.teamID).
					WillReturnRows(sqlmock.NewRows([]string{
						"team_id",
						"pokemon_id",
						"pokemon_name",
					}).AddRow(1, 1, "bulbasaur").AddRow(1, 132, "ditto"))
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
				mock.ExpectQuery(`SELECT team_id, pokemon_id, pokemon_name FROM "pokemon_name_in_teams" WHERE team_id = \$1`).WithArgs(args.teamID).
					WillReturnError(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &repository{
				db: gormDB,
			}
			got, err := r.List(context.Background(), tt.args.teamID)
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

func Test_repository_GetPokemonCount(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	assert.NoError(t, err)

	type args struct {
		teamID uint
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				teamID: 1,
			},
			want:    6,
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectQuery(`SELECT count\(\*\) FROM "pokemon_name_in_teams" WHERE team_id = \$1`).WithArgs(args.teamID).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(6))
			},
		},
		{
			name: "failed",
			args: args{
				teamID: 1,
			},
			want:    0,
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectQuery(`SELECT count\(\*\) FROM "pokemon_name_in_teams" WHERE team_id = \$1`).WithArgs(args.teamID).
					WillReturnError(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &repository{
				db: gormDB,
			}
			got, err := r.GetPokemonCount(context.Background(), tt.args.teamID)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetPokemonCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("repository.GetPokemonCount() = %v, want %v", got, tt.want)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
