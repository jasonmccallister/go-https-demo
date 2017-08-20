package response_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/themccallister/go-https-demo/response"
)

func TestJSON(t *testing.T) {
	rec := httptest.NewRecorder()

	response.JSON(rec, []byte("hello"), http.StatusAccepted)

	if rec.Code != http.StatusAccepted {
		t.Fatalf("expected the response code to be StatusAccepted, got %v instead", rec.Code)
	}
	if rec.Code == http.StatusBadGateway {
		t.Fatalf("exptected StatusAccepted, got %v instead", rec.Code)
	}
}
