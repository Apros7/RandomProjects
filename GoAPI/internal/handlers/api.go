package handlers

import (
	"goapi/internal/middleware"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) { // starting with a large letter makes other packages import
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})
}
