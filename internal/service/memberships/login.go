package memberships

import (
	"errors"

	"github.com/Fairuzzzzz/pokedex-api/internal/models/memberships"
	"github.com/Fairuzzzzz/pokedex-api/pkg/jwt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *service) Login(request memberships.LoginRequest) (string, error) {
	userDetails, err := s.repository.GetUser(request.Email, "", 0)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("error get user from database")
		return "", err
	}

	if userDetails == nil {
		errors.New("email not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDetails.Password), []byte(request.Password))
	if err != nil {
		return "", errors.New("email or password not match")
	}

	accessToken, err := jwt.CreateToken(int64(userDetails.ID), userDetails.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		log.Error().Err(err).Msg("error get access token")
		return "", nil
	}

	return accessToken, nil
}
