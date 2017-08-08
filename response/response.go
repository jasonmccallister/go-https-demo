package response

import (
	"encoding/json"
	"net/http"
)

// JSON returns a JSON response with a message and status. This is used for successful responses to standardize the response output.
func JSON(w http.ResponseWriter, msg string, s int) {
	res := []string{msg}
	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(s)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
