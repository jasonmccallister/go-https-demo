package main

import (
	"log"
	"net/http"

	"github.com/themccallister/go-https-demo/handlers"
	"github.com/themccallister/go-https-demo/helpers"
)

func main() {
	// grab the environment vars required
	port := helpers.EnvStringOr("APP_PORT", "4430")
	certFile := helpers.EnvStringOr("APP_CERT_FILE", "certs/server.crt")
	keyFile := helpers.EnvStringOr("APP_KEY_FILE", "certs/server.key")

	r := http.NewServeMux()

	// setup the routing
	r.HandleFunc("/", handlers.HomeIndexHandler)
	r.HandleFunc("/v2/users", handlers.UsersIndexHandler)

	if err := http.ListenAndServeTLS(":"+port, certFile, keyFile, r); err != nil {
		log.Fatal(err)
	}
}
