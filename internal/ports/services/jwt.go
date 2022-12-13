package services

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtService interface {
	GetCookieDuration() time.Duration
	ValidateToken(encodedToken string) (jwt.RegisteredClaims, error)
}
