// Writes the API's OpenAPI description into generated/ (gitignored), for API clients
// such as Postman: openapi.yaml in OpenAPI 3.1, huma's own, and openapi-3.0.yaml in
// 3.0, for tools that don't read 3.1. Built from the code, with no server started.
// Run by the `openapi` script, from the app's folder.
package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Vectoras/hello-015/apps/accounts/internal/server"
)

func main() {
	spec := server.OpenAPI()

	v31, err := spec.YAML()
	if err != nil {
		log.Fatalf("OpenAPI 3.1: %v", err)
	}
	v30, err := spec.DowngradeYAML()
	if err != nil {
		log.Fatalf("OpenAPI 3.0: %v", err)
	}

	if err := os.MkdirAll("generated", 0o755); err != nil {
		log.Fatal(err)
	}
	for name, content := range map[string][]byte{"openapi.yaml": v31, "openapi-3.0.yaml": v30} {
		path := filepath.Join("generated", name)
		if err := os.WriteFile(path, content, 0o644); err != nil {
			log.Fatal(err)
		}
		log.Printf("wrote %s", path)
	}
}
