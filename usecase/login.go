package usecase

import (
	"context"
	"time"

	"github.com/xorwise/golang-tz/domain"
	"github.com/xorwise/golang-tz/internal/utils"
	"github.com/xorwise/golang-tz/internal/uuid"
)

type loginUsecase struct {
	repo    domain.UserRepository
	timeout time.Duration
}

func NewLoginUsecase(repo domain.UserRepository, timeout time.Duration) domain.LoginUsecase {
	return &loginUsecase{
		repo:    repo,
		timeout: timeout,
	}
}

func (l *loginUsecase) CreateAccessToken(secret string, id uuid.UUID, expiry int) (string, error) {
	return utils.CreateAccessToken(secret, id, expiry)
}

func (l *loginUsecase) CreateRefreshToken(secret string, expiry int) (string, error) {
	return utils.CreateRefreshToken(secret, expiry)
}

func (l *loginUsecase) CreateOrUpdate(ctx context.Context, user *domain.User) error {
	return l.repo.InsertOrUpdate(ctx, user)
}

func (l *loginUsecase) HashRefreshToken(refreshToken string) string {
	return utils.Hash(refreshToken)
}
