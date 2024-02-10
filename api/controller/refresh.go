package controller

import (
	"encoding/json"
	"net/http"

	"github.com/xorwise/golang-tz/bootstrap"
	"github.com/xorwise/golang-tz/domain"
	"github.com/xorwise/golang-tz/internal/utils"
)

type RefreshController struct {
	RefreshUsecase domain.RefreshUsecase
	Env            *bootstrap.Env
}

func (rc *RefreshController) Refresh(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var request domain.RefreshRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	refreshHash := rc.RefreshUsecase.HashRefreshToken(request.RefreshToken)
	user, err := rc.RefreshUsecase.GetUserByRefresh(r.Context(), refreshHash)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("invalid refresh token")
		return
	}

	err = rc.RefreshUsecase.CheckRefreshToken(rc.Env.RefreshTokenSecret, request.RefreshToken)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("invalid refresh token")
		return
	}

	accessToken, err := rc.RefreshUsecase.CreateAccessToken(rc.Env.AccessTokenSecret, user.ID, rc.Env.AccessTokenExpiry)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	refreshToken, err := rc.RefreshUsecase.CreateRefreshToken(rc.Env.RefreshTokenSecret, rc.Env.RefreshTokenExpiry)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	refreshHash = utils.Hash(refreshToken)
	user.RefreshToken = refreshHash
	err = rc.RefreshUsecase.Update(r.Context(), &user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	response := domain.RefreshResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
