package main

import (
	"fmt"
	"log"
	"net/http"

	homehandler "github.com/themccallister/go-https-demo/handlers"
	usershandler "github.com/themccallister/go-https-demo/handlers/users"
	"github.com/themccallister/go-https-demo/helpers"
)

func main() {
	// grab the environment vars required
	port := helpers.EnvStringOr("APP_PORT", "443")
	certFile := helpers.EnvStringOr("APP_CERT_FILE", "certs/local.demodomain.com.crt")
	keyFile := helpers.EnvStringOr("APP_KEY_FILE", "certs/local.demodomain.com.key")

	fmt.Println("Running on port", port)

	r := http.NewServeMux()

	// setup the routing
	r.HandleFunc("/", homehandler.Index)
	r.HandleFunc("/v2/users", usershandler.Index)

	if err := http.ListenAndServeTLS(":"+port, certFile, keyFile, r); err != nil {
		log.Fatal(err)
	}
}
