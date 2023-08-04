package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func createV1Router(apiCfg *apiConfig) http.Handler {
	v1Router := chi.NewRouter()

	//Test endpoint
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)

	//User endpoints
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.handlerGetUser)

	return v1Router
}
