package types

// Hello world: proves the Go → TS type generation (the `generate` task writes its TS
// twin into generated/). Delete it when the real work starts.

// HelloWorld is what the API's /hello-world returns.
type HelloWorld struct {
	Message string `json:"message"`
}
