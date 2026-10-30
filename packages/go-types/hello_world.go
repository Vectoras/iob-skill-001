package gotypes

// Hello world: proves a type shared between Go apps (every Go app's
// /hello-world/go-types returns it). Delete it when the real work starts.

// HelloWorldApp names the app that answered. (Not HelloWorld: huma names each
// response schema after its Go type, without the package, and shared-types has one.)
type HelloWorldApp struct {
	App string `json:"app"`
}
