// Package server builds the service's HTTP handler: middleware, routes and handlers.
package server

import (
	"crypto/subtle"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/Vectoras/hello-015/apps/accounts/internal/config"
)

// New returns the service's handler. chi's middleware comes first: chi panics on
// middleware added after a route, and huma's adapter registers routes as it is created.
func New(cfg config.Config) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	// Answered here, before the routes and any secret check: the health check needs nothing.
	router.Use(middleware.Heartbeat("/healthz"))
	router.Use(requireSecret(cfg.APISecret))

	register(humachi.New(router, humaConfig()))
	return router
}

// OpenAPI returns the API's OpenAPI description, built from the same routes as New but
// with no middleware and no server (cmd/openapi writes it to files). Handlers are not
// called: a feature that needs a dependency to register its routes gets a zero value here.
func OpenAPI() *huma.OpenAPI {
	api := humachi.New(chi.NewRouter(), humaConfig())
	register(api)
	return api.OpenAPI()
}

func humaConfig() huma.Config {
	humaConfig := huma.DefaultConfig("accounts", "0.0.1")
	// No `$schema` link added to every response body.
	humaConfig.CreateHooks = nil
	// The description says every operation needs the secret header, so API clients
	// (Postman) ask for it once instead of every request failing with 401.
	humaConfig.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"apiSecret": {Type: "apiKey", In: "header", Name: "X-Api-Secret"},
	}
	humaConfig.Security = []map[string][]string{{"apiSecret": {}}}
	return humaConfig
}

// register adds every route: one register<Feature>(api) call per feature, each in its
// own file.
func register(api huma.API) {
	registerHelloWorldGoTypes(api)
	registerHelloWorldSharedTypes(api)
}

// requireSecret refuses every request that doesn't carry the shared secret in its own
// header. Only the frontend apps' servers know it, so the API answers no one else.
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
