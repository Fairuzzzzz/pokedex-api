package team

import (
	"net/http"

	"github.com/Fairuzzzzz/pokedex-api/internal/models/team"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTeam(c *gin.Context) {
	ctx := c.Request.Context()

	var request team.PokeTeamNameRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	userID := c.GetUint("userID")
	err := h.service.CreateTeam(ctx, userID, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) ListTeam(c *gin.Context) {
	ctx := c.Request.Context()

	userID := c.GetUint("userID")
	teams, err := h.service.ListTeam(ctx, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, teams)
}

func (h *Handler) GetTeam(c *gin.Context) {
	ctx := c.Request.Context()

	var request team.PokeTeamRequestByID
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	request.UserID = c.GetUint("userID")
	team, err := h.service.GetTeam(ctx, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, team)
}

func (h *Handler) DeleteTeam(c *gin.Context) {
	ctx := c.Request.Context()

	var request team.PokeTeamRequestByID
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	request.UserID = c.GetUint("userID")
	err := h.service.DeleteTeam(ctx, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}
