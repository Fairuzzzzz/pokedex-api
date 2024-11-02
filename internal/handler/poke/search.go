package poke

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Search(c *gin.Context) {
	ctx := c.Request.Context()

	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "pokemon name is required",
		})
		return
	}

	response, err := h.service.PokemonSearch(ctx, name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
