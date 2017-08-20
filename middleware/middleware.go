package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
	"net/http"
)

// HELPERS

// checkForRequestID checks if the ResponseWriter has a header, otherwise
// it will set a new random ID for the request and assign to "Request-ID" on
// the ResponseWriter header.
func checkForRequestID(w http.ResponseWriter, r *http.Request) {
	if reqID := w.Header().Get("Request-ID"); reqID == "" {
		w.Header().Add("Request-ID", requestID())
	}
}

// logRequest takes a ResponseWriter and Request and generates the new log
// for the request.
func logRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("[METHOD] %s | [URI] %s | [REQUEST ID] %s", r.Method, r.URL.String(), w.Header().Get("Request-ID"))
}

// requestID generates a random string for a unique request ID. This is used
// for assigning a unique string to a ResponseWriter header.
func requestID() string {
	b := make([]byte, 6)
	io.ReadFull(rand.Reader, b)
	return base64.StdEncoding.EncodeToString(b)
}
