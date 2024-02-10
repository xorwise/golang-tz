package usecase

import (
	"context"
	"time"

	"github.com/xorwise/golang-tz/domain"
	"github.com/xorwise/golang-tz/internal/utils"
	"github.com/xorwise/golang-tz/internal/uuid"
)

type refreshUsecase struct {
	repo    domain.UserRepository
	timeout time.Duration
}

func NewRefreshUsecase(repo domain.UserRepository, timeout time.Duration) domain.RefreshUsecase {
	return &refreshUsecase{
		repo:    repo,
		timeout: timeout,
	}
}

func (u *refreshUsecase) Update(ctx context.Context, user *domain.User) error {
	return u.repo.Update(ctx, user)
}

func (u *refreshUsecase) GetUserByRefresh(ctx context.Context, refreshToken string) (domain.User, error) {
	return u.repo.GetByRefresh(ctx, refreshToken)
}

func (u *refreshUsecase) CreateAccessToken(secret string, id uuid.UUID, expiry int) (string, error) {
	return utils.CreateAccessToken(secret, id, expiry)
}

func (u *refreshUsecase) CreateRefreshToken(secret string, expiry int) (string, error) {
	return utils.CreateRefreshToken(secret, expiry)
}

func (u *refreshUsecase) HashRefreshToken(refreshToken string) string {
	return utils.Hash(refreshToken)
}

func (u *refreshUsecase) CheckRefreshToken(secret string, refreshToken string) error {
	return utils.CheckToken(secret, refreshToken)
}
