package server

// Hello world: delete it with hello_world_go_types.go when the real work starts.

import (
	"encoding/json"
	"net/http"
	"testing"

	gotypes "github.com/Vectoras/hello-015/packages/go-types"
)

func TestHelloWorldGoTypes(t *testing.T) {
	rec := get(newTestHandler(), "/hello-world/go-types", testSecret)
	if rec.Code != http.StatusOK {
		t.Fatalf("/hello-world/go-types: got %d, want 200", rec.Code)
	}
	var body gotypes.HelloWorldApp
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("decoding the body: %v", err)
	}
	if body.App != "accounts" {
		t.Errorf("app: got %q, want %q", body.App, "accounts")
	}
}
