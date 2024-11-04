package team

import (
	"gorm.io/gorm"
)

type (
	PokeTeam struct {
		gorm.Model
		UserID    uint   `gorm:"not null"`
		TeamName  string `gorm:"not null"`
		CreatedBy string `gorm:"not null"`
		UpdatedBy string `gorm:"not null"`
	}
)

type (
	PokeTeamNameRequest struct {
		TeamName string `json:"teamName"`
	}

	PokeTeamRequestByID struct {
		ID     uint `json:"id"`
		UserID uint `json:"userID"`
	}
)
