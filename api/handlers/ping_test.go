package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	req, _ := http.NewRequest("GET", "/ping", nil)

	handler := http.HandlerFunc(Ping)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	checkResponseCode(t, http.StatusOK, response.Code)
}
