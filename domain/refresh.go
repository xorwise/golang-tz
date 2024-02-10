package domain

import (
	"context"

	"github.com/xorwise/golang-tz/internal/uuid"
)

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshUsecase interface {
	Update(ctx context.Context, user *User) error
	GetUserByRefresh(ctx context.Context, refreshToken string) (User, error)
	CreateAccessToken(secret string, id uuid.UUID, expiry int) (string, error)
	CreateRefreshToken(secret string, expiry int) (string, error)
	HashRefreshToken(refreshToken string) string
	CheckRefreshToken(secret string, refreshToken string) error
}
