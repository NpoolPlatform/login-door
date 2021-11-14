package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// https://github.com/grpc/grpc-go/issues/3794
// require_unimplemented_servers=false
type Server struct {
	*chi.Mux
}

func Register(router *chi.Mux) {
	router.Use(middleware.Logger)
	router.Use(middleware.Heartbeat("/"))

	router.Post("/version", Version) // nolint: typecheck
}
