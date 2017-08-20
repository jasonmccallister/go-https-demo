package main

import (
	"fmt"
	"log"
	"net/http"

	// bugsnag "github.com/bugsnag/bugsnag-go"
	homehandler "github.com/themccallister/go-https-demo/handlers"
	usershandler "github.com/themccallister/go-https-demo/handlers/users"
	"github.com/themccallister/go-https-demo/helpers/env"
	"github.com/themccallister/go-https-demo/middleware"
)

func routes(r *http.ServeMux) {

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

	// setup the routing
	r.HandleFunc("/", homehandler.Index)
	r.HandleFunc("/v2/users", middleware.RequestLogging(usershandler.Index))

	if err := http.ListenAndServeTLS(":"+port, certFile, keyFile, r); err != nil {
		log.Fatal(err)
	}
}
