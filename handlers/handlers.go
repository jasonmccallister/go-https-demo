package handlers

import (
	"net/http"

	"github.com/themccallister/go-https-demo/response"
)

func HomeIndexHandler(w http.ResponseWriter, req *http.Request) {
	response.JSON(w, []byte("message loaded ok"), http.StatusOK)
}
