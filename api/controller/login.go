package controller

import (
	"encoding/json"
	"net/http"

	"github.com/xorwise/golang-tz/bootstrap"
	"github.com/xorwise/golang-tz/domain"
	"github.com/xorwise/golang-tz/internal/uuid"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Env          *bootstrap.Env
}

func (lc *LoginController) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := uuid.FromString(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	accessToken, err := lc.LoginUsecase.CreateAccessToken(lc.Env.AccessTokenSecret, id, lc.Env.AccessTokenExpiry)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiry)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	refreshHash := lc.LoginUsecase.HashRefreshToken(refreshToken)
	user := domain.User{
		ID:           id,
		RefreshToken: refreshHash,
	}
	err = lc.LoginUsecase.CreateOrUpdate(r.Context(), &user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	loginReponse := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(loginReponse)
}
