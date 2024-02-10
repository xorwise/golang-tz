package domain

import (
	"context"

	"github.com/xorwise/golang-tz/internal/uuid"
)

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginUsecase interface {
	CreateAccessToken(secret string, id uuid.UUID, expiry int) (string, error)
	CreateRefreshToken(secret string, expiry int) (string, error)
	CreateOrUpdate(ctx context.Context, user *User) error
	HashRefreshToken(refreshToken string) string
}
