// The service's entry point. It only wires things together: config, handler, server.
package main

import (
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/Vectoras/hello-020/apps/accounts/internal/config"
	"github.com/Vectoras/hello-020/apps/accounts/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	srv := &http.Server{
		// An empty host listens on every interface, as a server in a container must.
		Addr:              net.JoinHostPort("", strconv.Itoa(cfg.Port)),
		Handler:           server.New(cfg),
		ReadHeaderTimeout: 10 * time.Second,
	}
	log.Printf("listening on %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
