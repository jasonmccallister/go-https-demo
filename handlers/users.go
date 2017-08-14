package handlers

import (
	"net/http"

	"github.com/themccallister/go-https-demo/response"
)

// UsersHandler shows all of the users
func UsersHandler(w http.ResponseWriter, req *http.Request) {
	response.JSON(w, "hello from users", http.StatusOK)
}
