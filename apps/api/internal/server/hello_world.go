package server

// Hello world: proves the API's wiring (chi, huma, the shared secret) and the shared
// Go types. Delete it, with its test, when the real work starts.

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"github.com/Vectoras/iob-skill-001/packages/types"
)

type helloWorldOutput struct {
	Body types.HelloWorld
}

func registerHelloWorld(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-hello-world",
		Method:      http.MethodGet,
		Path:        "/hello-world",
		Summary:     "Hello world",
	}, func(_ context.Context, _ *struct{}) (*helloWorldOutput, error) {
		return &helloWorldOutput{Body: types.HelloWorld{Message: "Hello world!"}}, nil
	})
}
