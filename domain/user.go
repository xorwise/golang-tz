package domain

import (
	"context"

	"github.com/xorwise/golang-tz/internal/uuid"
)

type Collection string

const (
	UserCollection Collection = "users"
)

type User struct {
	ID           uuid.UUID `bson:"_id"`
	RefreshToken string    `bson:"refresh_token"`
}

type UserRepository interface {
	Update(c context.Context, user *User) error
	GetByRefresh(c context.Context, refreshToken string) (User, error)
	InsertOrUpdate(c context.Context, user *User) error
}
