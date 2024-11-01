package poke

import (
	"context"

	"github.com/Fairuzzzzz/pokedex-api/internal/models/poke"
	"github.com/gin-gonic/gin"
)

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
