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

	router.Post("/version", Version)                      // nolint: typecheck
	router.Post("/v1/add/provider", AddProvider)          // nolint: typecheck
	router.Post("/v1/update/provider", UpdateProvider)    // nolint: typecheck
	router.Post("/v1/get/all/providers", GetAllProviders) // nolint: typecheck
	router.Post("/v1/delete/provider", DeleteProvider)    // nolint: typecheck
}
