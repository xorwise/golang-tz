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

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, mux *http.ServeMux) {
	repo := repository.NewUserRepository(db)
	sc := controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(repo, timeout),
		Env:          env,
	}

	mux.HandleFunc("GET /api/login/{id}", sc.Login)
}
