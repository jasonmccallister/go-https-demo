package main

import (
	"fmt"
	"log"
	"net/http"

	// bugsnag "github.com/bugsnag/bugsnag-go"

	"github.com/themccallister/go-https-demo/handlers/home"
	"github.com/themccallister/go-https-demo/handlers/users"
	"github.com/themccallister/go-https-demo/helpers/env"
	"github.com/themccallister/go-https-demo/middleware"
)

// routes handles the path pattern to a specific HTTP handler
func routes(r *http.ServeMux) {
	r.HandleFunc("/", home.Index)
	r.HandleFunc("/v2/users", middleware.RequestLogging(users.Index))
}

func main() {
	// grab the environment vars required
	port := env.GetOr("APP_PORT", "443")
	certFile := env.GetOr("APP_CERT_FILE", "certs/local.demodomain.com.crt")
	keyFile := env.GetOr("APP_KEY_FILE", "certs/local.demodomain.com.key")
	// bugsnagKey := helpers.EnvStringOr("BUGSNAG_KEY", "")
	// if bugsnagKey != "" {
	// 	bugsnag.Configure(bugsnag.Configuration{
	// 		APIKey: bugsnagKey,
	// 	})
	// }

	fmt.Println("Running on port", port)

	r := http.NewServeMux()
	routes(r)
	if err := http.ListenAndServeTLS(":"+port, certFile, keyFile, r); err != nil {
		log.Fatal(err)
	}
}
