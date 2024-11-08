package pokemoninteam

import (
	"net/http"
	"strconv"

	"github.com/Fairuzzzzz/pokedex-api/internal/models/pokemoninteam"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AddPokemon(c *gin.Context) {
	ctx := c.Request.Context()

	var request pokemoninteam.PokemonRequestWithName
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	userID := c.GetUint("userID")
	err := h.service.AddPokemonToTeam(ctx, userID, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) DeletePokemon(c *gin.Context) {
	ctx := c.Request.Context()

	var request pokemoninteam.PokemonRequestWithID
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.service.DeletePokemon(ctx, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) ListPokemon(c *gin.Context) {
	ctx := c.Request.Context()

	teamIDStr := c.Param("teamID")
	teamID, err := strconv.ParseUint(teamIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid team ID",
		})
		return
	}

	pokemons, err := h.service.ListPokemon(ctx, uint(teamID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, pokemons)
}
