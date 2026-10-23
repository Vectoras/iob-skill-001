// Package server builds the API's HTTP handler: middleware, routes and handlers.
package server

import (
	"crypto/subtle"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/Vectoras/iob-skill-001/apps/api/internal/config"
)

// New returns the API's handler. chi's middleware comes first: chi panics on
// middleware added after a route, and huma's adapter registers routes as it is created.
func New(cfg config.Config) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	// Answered here, before the secret check: the health check needs no secret.
	router.Use(middleware.Heartbeat("/healthz"))
	router.Use(requireSecret(cfg.APISecret))

	humaConfig := huma.DefaultConfig("api", "0.0.1")
	// No `$schema` link added to every response body.
	humaConfig.CreateHooks = nil
	api := humachi.New(router, humaConfig)

	// One register<Feature>(api) call per feature, each in its own file.
	registerHelloWorld(api)

	return router
}

// requireSecret refuses every request that doesn't carry the shared secret in its own
// header. Only the frontend's server knows it, so the API answers no one else.
// Authorization stays free for the logged-in user's credentials.
func requireSecret(secret string) func(http.Handler) http.Handler {
	want := []byte(secret)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			got := []byte(r.Header.Get("X-Api-Secret"))
			if subtle.ConstantTimeCompare(got, want) != 1 {
				http.Error(w, "missing or wrong X-Api-Secret", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
