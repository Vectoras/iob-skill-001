package server

// Hello world: proves the Go tests run, against the real handler. Delete it with
// hello_world.go when the real work starts.

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Vectoras/iob-skill-001/apps/api/internal/config"
	"github.com/Vectoras/iob-skill-001/packages/types"
)

func TestHelloWorld(t *testing.T) {
	handler := New(config.Config{APISecret: "test-secret"})

	get := func(path, secret string) *httptest.ResponseRecorder {
		req := httptest.NewRequest(http.MethodGet, path, nil)
		if secret != "" {
			req.Header.Set("X-Api-Secret", secret)
		}
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
		return rec
	}

	if rec := get("/healthz", ""); rec.Code != http.StatusOK {
		t.Errorf("/healthz without the secret: got %d, want 200", rec.Code)
	}
	if rec := get("/hello-world", ""); rec.Code != http.StatusUnauthorized {
		t.Errorf("/hello-world without the secret: got %d, want 401", rec.Code)
	}
	if rec := get("/hello-world", "wrong"); rec.Code != http.StatusUnauthorized {
		t.Errorf("/hello-world with a wrong secret: got %d, want 401", rec.Code)
	}

	rec := get("/hello-world", "test-secret")
	if rec.Code != http.StatusOK {
		t.Fatalf("/hello-world with the secret: got %d, want 200", rec.Code)
	}
	var body types.HelloWorld
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("decoding the body: %v", err)
	}
	if body.Message != "Hello world!" {
		t.Errorf("message: got %q, want %q", body.Message, "Hello world!")
	}
}
