package utils

import (
	"coffe-life/config"
	"coffe-life/internal/entity"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwtToken(user entity.User, tokenCfg config.JwtToken) (string, error) {
	tokenTimeLimit := tokenCfg.GetTokenTimeLimit()

	hmacSampleSecret := tokenCfg.GetJwtTokenSecret()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"exp":    time.Now().Add(tokenTimeLimit).Unix(),
	})

	tokenString, err := token.SignedString([]byte(hmacSampleSecret))
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}
