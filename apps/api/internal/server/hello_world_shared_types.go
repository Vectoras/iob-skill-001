package server

// Hello world: proves the private API's wiring (chi, huma, the shared secret) and the
// types shared with the frontend apps. Delete it, with its test, when the real work
// starts.

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	sharedtypes "github.com/Vectoras/hello-019/packages/shared-types"
)

type helloWorldSharedTypesOutput struct {
	Body sharedtypes.HelloWorld
}

func registerHelloWorldSharedTypes(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-hello-world-shared-types",
		Method:      http.MethodGet,
		Path:        "/hello-world/shared-types",
		Summary:     "Hello world (shared-types)",
	}, func(_ context.Context, _ *struct{}) (*helloWorldSharedTypesOutput, error) {
		return &helloWorldSharedTypesOutput{Body: sharedtypes.HelloWorld{Message: "Hello world!"}}, nil
	})
}
