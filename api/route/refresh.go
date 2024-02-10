package route

import (
	"net/http"
	"time"

	"github.com/xorwise/golang-tz/api/controller"
	"github.com/xorwise/golang-tz/bootstrap"
	"github.com/xorwise/golang-tz/repository"
	"github.com/xorwise/golang-tz/usecase"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRefreshRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, mux *http.ServeMux) {
	repo := repository.NewUserRepository(db)
	rc := controller.RefreshController{
		RefreshUsecase: usecase.NewRefreshUsecase(repo, timeout),
		Env:            env,
	}
	mux.HandleFunc("POST /api/refresh", rc.Refresh)
}
