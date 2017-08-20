package handlers

import (
	"errors"
	"net/http"
)

// HELPERS

// This contains all of the helper methods, variables and constants
// used throughout the package. Before attempting to create a new
// function to share logic in this package, you should check for
// pre-existing code here, if it does not exist you can add it here.

func requirePOST(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return nil
	}
	return errors.New("text")
}
