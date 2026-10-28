package sharedtypes

// Hello world: proves the Go → TS type generation (the `generate` task writes its TS
// twin into generated/) and the Zod schema checked against it (src/hello-world.ts).
// Delete it, with them, when the real work starts.

// HelloWorld is what every private API's /hello-world/shared-types returns.
type HelloWorld struct {
	Message string `json:"message"`
}
