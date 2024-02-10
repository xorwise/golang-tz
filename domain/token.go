package domain

import (
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/xorwise/golang-tz/internal/uuid"
)

type JwtCustomClaims struct {
	ID uuid.UUID `json:"id"`
	jwt.RegisteredClaims
}
