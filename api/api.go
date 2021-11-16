// Package Login Door service API
//
// This service is to implement api about login.
//
//      Host: localhost
//      Version: 0.0.1
//
// swagger:meta

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
	router.Post("/v1/login", Login)                       // nolint: typecheck
	router.Post("/v1/get/user/login", GetUserLogin)       // nolint: typecheck
	router.Post("/v1/logout", Logout)                     // nolint: typecheck
	router.Post("/v1/refresh/session", RefreshSession)    // nolint: typecheck
}
