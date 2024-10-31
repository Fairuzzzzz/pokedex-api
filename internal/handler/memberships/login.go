package memberships

import (
	"net/http"

	"github.com/Fairuzzzzz/pokedex-api/internal/models/memberships"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	var request memberships.LoginRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	accessToken, err := h.service.Login(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, memberships.LoginResponse{
		AccessToken: accessToken,
	})
}
