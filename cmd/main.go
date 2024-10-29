package main

import (
	"log"

	"github.com/Fairuzzzzz/pokedex-api/internal/configs"
	membershipsHandler "github.com/Fairuzzzzz/pokedex-api/internal/handler/memberships"
	"github.com/Fairuzzzzz/pokedex-api/internal/models/memberships"
	membershipsRepo "github.com/Fairuzzzzz/pokedex-api/internal/repository/memberships"
	membershipsSvc "github.com/Fairuzzzzz/pokedex-api/internal/service/memberships"
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
		log.Fatalf("Gagal initialisasi config", err)
	}

	cfg = configs.Get()
	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to database, err: %+v\n", err)
	}

	db.AutoMigrate(&memberships.User{})

	membershipRepo := membershipsRepo.NewRepository(db)

	membershipSvc := membershipsSvc.NewService(cfg, membershipRepo)

	membershipHandler := membershipsHandler.NewHandler(r, membershipSvc)
	membershipHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
