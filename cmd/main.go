package main

import (
	"log"
	"net/http"

	"github.com/Fairuzzzzz/pokedex-api/internal/configs"
	membershipsHandler "github.com/Fairuzzzzz/pokedex-api/internal/handler/memberships"
	pokesHandler "github.com/Fairuzzzzz/pokedex-api/internal/handler/poke"
	"github.com/Fairuzzzzz/pokedex-api/internal/models/memberships"
	"github.com/Fairuzzzzz/pokedex-api/internal/models/team"
	membershipsRepo "github.com/Fairuzzzzz/pokedex-api/internal/repository/memberships"
	pokesOutbound "github.com/Fairuzzzzz/pokedex-api/internal/repository/poke"
	membershipsSvc "github.com/Fairuzzzzz/pokedex-api/internal/service/memberships"
	pokesSvc "github.com/Fairuzzzzz/pokedex-api/internal/service/poke"
	"github.com/Fairuzzzzz/pokedex-api/pkg/httpclient"
	"github.com/Fairuzzzzz/pokedex-api/pkg/internalsql"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs/"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("Gagal initialisasi config", err)
	}

	cfg = configs.Get()
	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to database, err: %+v\n", err)
	}

	db.AutoMigrate(&memberships.User{})
	db.AutoMigrate(&team.PokeTeam{})

	httpClient := httpclient.NewClient(&http.Client{})

	membershipRepo := membershipsRepo.NewRepository(db)

	pokeOutbound := pokesOutbound.NewPokeOutbound(cfg, httpClient)

	membershipSvc := membershipsSvc.NewService(cfg, membershipRepo)

	pokeSvc := pokesSvc.NewOutbound(pokeOutbound)

	membershipHandler := membershipsHandler.NewHandler(r, membershipSvc)
	membershipHandler.RegisterRoute()

	pokeHandler := pokesHandler.NewHandler(r, pokeSvc)
	pokeHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
