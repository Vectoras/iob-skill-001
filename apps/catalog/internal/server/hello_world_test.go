package server

// Hello world: proves the Go tests run, against the real handler, and that the health
// check answers. Delete it when the real work starts, keeping what other tests use.

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Vectoras/hello-013/apps/catalog/internal/config"
)

// testSecret is empty: this service requires no secret.
const testSecret = ""

// newTestHandler returns the service's handler, configured for tests.
func newTestHandler() http.Handler {
	return New(config.Config{})
}

// get sends a GET request to the handler, with the secret when one is given.
func get(handler http.Handler, path, secret string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, path, nil)
	if secret != "" {
		req.Header.Set("X-Api-Secret", secret)
	}
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	return rec
}

func TestHealthz(t *testing.T) {
	// The secret, when there is one, changes nothing: the health check is answered first.
	if rec := get(newTestHandler(), "/healthz", testSecret); rec.Code != http.StatusOK {
		t.Errorf("/healthz: got %d, want 200", rec.Code)
	}
}
