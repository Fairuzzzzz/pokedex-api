package poke

import (
	"context"

	"github.com/Fairuzzzzz/pokedex-api/internal/middleware"
	"github.com/Fairuzzzzz/pokedex-api/internal/models/poke"
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=handler.go -destination=handler_mock_test.go -package=poke
type service interface {
	PokemonSearch(ctx context.Context, name string) (*poke.SearchPokemonResponse, error)
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
	route := h.Group("/pokemon")
	route.Use(middleware.AuthMiddleware())
	route.GET("/:name", h.Search)
}
