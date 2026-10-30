package server

// Hello world: proves the Go tests run, against the real handler, and that the health
// check answers. Delete it when the real work starts, keeping what other tests use.

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Vectoras/hello-020/apps/orders/internal/config"
)

// testSecret is the shared secret the tests' handler expects.
const testSecret = "test-secret"

// newTestHandler returns the service's handler, configured for tests.
func newTestHandler() http.Handler {
	return New(config.Config{APISecret: testSecret})
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

func TestSecretRequired(t *testing.T) {
	handler := newTestHandler()
	if rec := get(handler, "/", ""); rec.Code != http.StatusUnauthorized {
		t.Errorf("without the secret: got %d, want 401", rec.Code)
	}
	if rec := get(handler, "/", "wrong"); rec.Code != http.StatusUnauthorized {
		t.Errorf("with a wrong secret: got %d, want 401", rec.Code)
	}
	// Past the secret check, an unknown path is a 404.
	if rec := get(handler, "/", testSecret); rec.Code != http.StatusNotFound {
		t.Errorf("with the secret: got %d, want 404", rec.Code)
	}
}
