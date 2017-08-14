package response

import (
	"net/http"
)

// JSON returns a JSON response with a message and status. This is used for successful responses to standardize the response output.
func JSON(w http.ResponseWriter, msg string, status int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(msg))
}
