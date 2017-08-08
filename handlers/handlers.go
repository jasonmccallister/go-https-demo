package handlers

import (
	"net/http"

	"github.com/themccallister/go-https-demo/response"
)

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	response.JSON(w, "message loaded ok", http.StatusOK)
}
