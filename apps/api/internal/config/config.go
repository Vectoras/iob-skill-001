// Package config reads the service's settings from the environment, once, at startup.
package config

import "github.com/caarlos0/env/v11"

// Config lists every variable the service reads. A field without a default is
// required, and `notEmpty` refuses an empty value: either stops the service at
// startup. The struct is passed explicitly to what needs it; nothing else reads the
// environment.
type Config struct {
	Port int `env:"PORT"`
	// The shared secret the frontend's server sends in X-Api-Secret.
	APISecret string `env:"API_SECRET,notEmpty"`
}

// Load fills Config from the environment.
func Load() (Config, error) {
	var cfg Config
	err := env.ParseWithOptions(&cfg, env.Options{RequiredIfNoDef: true})
	return cfg, err
}
