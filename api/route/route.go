package route

import (
	"net/http"
	"time"

	"github.com/xorwise/golang-tz/bootstrap"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, mux *http.ServeMux) {
	NewLoginRouter(env, timeout, db, mux)
	NewRefreshRouter(env, timeout, db, mux)
}
