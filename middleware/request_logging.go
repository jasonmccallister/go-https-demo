package middleware

import (
	"net/http"
)

// RequestLogging handles all of the request and sets a new request ID for
// logging purposes.
var RequestLogging = func(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		checkForRequestID(w, r)
		f(w, r)
		logRequest(w, r)
	}
}
