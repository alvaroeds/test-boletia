package http

import (
	"github.com/alvaroeds/test-boletia/internal/db/postgres"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"net/http"

	currencyHandler "github.com/alvaroeds/test-boletia/pkg/currency"
)

// routes function sets routes handlers.
func routes(dbClient *postgres.Client) http.Handler {
	r := chi.NewRouter()

	co := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			"GET",
		},
		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"Cache-Control",
		},
		AllowCredentials: true,
	})

	r.Use(co.Handler)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	ch := currencyHandler.NewHandler(dbClient.DB)
	r.Route("/currency", func(r chi.Router) {
		r.Get("/", ch.GetCurrencyHandler)
		r.Get("/{currency}", ch.GetCurrencyHandler)
	})

	return r
}
