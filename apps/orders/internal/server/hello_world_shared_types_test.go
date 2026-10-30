package server

// Hello world: delete it with hello_world_shared_types.go when the real work starts.

import (
	"encoding/json"
	"net/http"
	"testing"

	sharedtypes "github.com/Vectoras/hello-020/packages/shared-types"
)

func TestHelloWorldSharedTypes(t *testing.T) {
	handler := newTestHandler()
	if rec := get(handler, "/hello-world/shared-types", ""); rec.Code != http.StatusUnauthorized {
		t.Errorf("without the secret: got %d, want 401", rec.Code)
	}

	rec := get(handler, "/hello-world/shared-types", testSecret)
	if rec.Code != http.StatusOK {
		t.Fatalf("with the secret: got %d, want 200", rec.Code)
	}
	var body sharedtypes.HelloWorld
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("decoding the body: %v", err)
	}
	if body.Message != "Hello world!" {
		t.Errorf("message: got %q, want %q", body.Message, "Hello world!")
	}
}
