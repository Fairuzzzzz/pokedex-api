package team

import (
	"context"

	"github.com/Fairuzzzzz/pokedex-api/internal/middleware"
	"github.com/Fairuzzzzz/pokedex-api/internal/models/team"
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=handler.go -destination=handler_mock_test.go -package=team
type service interface {
	CreateTeam(ctx context.Context, userID uint, request team.PokeTeamNameRequest) error
	ListTeam(ctx context.Context, userID uint) ([]team.PokeTeam, error)
	GetTeam(ctx context.Context, request team.PokeTeamRequestByID) (*team.PokeTeam, error)
	DeleteTeam(ctx context.Context, request team.PokeTeamRequestByID) error
}

type Handler struct {
	*gin.Engine
	service service
}

func NewHandler(api *gin.Engine, service service) *Handler {
	return &Handler{
		api,
		service,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("/team")
	route.Use(middleware.AuthMiddleware())
	route.POST("/create-team", h.CreateTeam)
	route.GET("/list-team", h.ListTeam)
	route.GET("/get-team", h.GetTeam)
	route.POST("/delete-team", h.DeleteTeam)
}
