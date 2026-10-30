package server

// Hello world: proves this app imports go-types, the types shared between Go apps.
// Delete it, with its test, when the real work starts.

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	gotypes "github.com/Vectoras/hello-018/packages/go-types"
)

type helloWorldGoTypesOutput struct {
	Body gotypes.HelloWorldApp
}

func registerHelloWorldGoTypes(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-hello-world-go-types",
		Method:      http.MethodGet,
		Path:        "/hello-world/go-types",
		Summary:     "Hello world (go-types)",
	}, func(_ context.Context, _ *struct{}) (*helloWorldGoTypesOutput, error) {
		return &helloWorldGoTypesOutput{Body: gotypes.HelloWorldApp{App: "catalog"}}, nil
	})
}
