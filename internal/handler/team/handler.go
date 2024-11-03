package team

import (
	"context"

	"github.com/Fairuzzzzz/pokedex-api/internal/models/team"
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=handler.go -destination=handler_mock_test.go -package=team
type service interface {
	CreateTeam(ctx context.Context, userID uint, request team.PokeTeamNameRequest) error
	ListTeam(ctx context.Context, userID uint) ([]team.PokeTeam, error)
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

// TODO: add handler for create team and list team
func (h *Handler) RegisterRoute() {
	route := h.Group("/team")
	route.POST("/create-team")
	route.GET("/list-team")
}
