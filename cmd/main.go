package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/xorwise/golang-tz/api/route"
	"github.com/xorwise/golang-tz/bootstrap"
)

func main() {
	env := bootstrap.NewEnv()
	db, client := bootstrap.InitializeDatabase(env)
	defer client.Disconnect(context.Background())

	mux := http.NewServeMux()

	route.Setup(env, 10*time.Second, db, mux)

	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
