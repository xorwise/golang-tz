package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"

	"github.com/xorwise/golang-tz/domain"
	"github.com/xorwise/golang-tz/internal/uuid"
)

func CreateAccessToken(secret string, id uuid.UUID, expiry int) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry))
	claims := &domain.JwtCustomClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func CreateRefreshToken(secret string, expiry int) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(exp),
	})
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func Hash(s string) string {
	hasher := sha256.New()
	hasher.Write([]byte(s))
	return hex.EncodeToString(hasher.Sum(nil))
}

func CheckToken(secret string, token string) error {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	})

	if err != nil || !t.Valid {
		return err
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok {
		expTime := time.Unix(int64(claims["exp"].(float64)), 0)
		curTime := time.Now()
		if expTime.Before(curTime) {
			return jwt.ErrTokenExpired
		}
	}
	return nil
}
