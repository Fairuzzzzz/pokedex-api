package team

import (
	"github.com/Fairuzzzzz/pokedex-api/internal/models/memberships"
	"gorm.io/gorm"
)

type (
	PokeTeam struct {
		gorm.Model
		UserID    uint             `gorm:"not null;index"`
		User      memberships.User `gorm:"foreigenKey:UserID"`
		TeamName  string           `gorm:"not null"`
		TeamID    uint             `gorm:"not null;index"`
		CreatedBy string           `gorm:"not null"`
		UpdatedBy string           `gorm:"not null"`
	}
)

type (
	PokeTeamNameRequest struct {
		TeamName string `json:"teamName"`
	}

	PokeTeamRequestByID struct {
		UserID uint `json:"userID"`
		TeamID uint `json:"teamID"`
	}
)
