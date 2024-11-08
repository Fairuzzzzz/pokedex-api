package pokemoninteam

import (
	"context"

	"github.com/Fairuzzzzz/pokedex-api/internal/middleware"
	"github.com/Fairuzzzzz/pokedex-api/internal/models/pokemoninteam"
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=handler.go -destination=handler_mock_test.go -package=pokemoninteam
type service interface {
	AddPokemonToTeam(ctx context.Context, userID uint, request pokemoninteam.PokemonRequestWithName) error
	DeletePokemon(ctx context.Context, request pokemoninteam.PokemonRequestWithID) error
	ListPokemon(ctx context.Context, teamID uint) ([]pokemoninteam.PokemonTeamResponse, error)
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
	route := h.Group("/pokemon-team")
	route.Use(middleware.AuthMiddleware())
	route.GET("/list-pokemon", h.ListPokemon)
	route.POST("/add", h.AddPokemon)
	route.POST("/remove", h.DeletePokemon)
}
