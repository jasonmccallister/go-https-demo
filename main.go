package main

import (
	"log"
	"net/http"

	"github.com/themccallister/go-https-demo/helpers"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", handlers.HomeHandler)
	port := helpers.EnvOr("APP_PORT", "4430")
	if err := http.ListenAndServeTLS(":"+port, "certs/server.crt", "certs/server.key", r); err != nil {
		log.Fatal(err)
	}
}
