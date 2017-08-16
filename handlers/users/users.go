package users

import (
	"encoding/json"
	"net/http"

	"github.com/themccallister/go-https-demo/response"
	"github.com/themccallister/go-https-demo/services"
)

// Index shows all of the users
func Index(w http.ResponseWriter, req *http.Request) {
	u := services.All()
	userJSON, err := json.Marshal(&u)
	if err != nil {
		response.JSON(w, []byte("malformed request"), http.StatusBadRequest)
		return
	}
	response.JSON(w, userJSON, http.StatusOK)
}
