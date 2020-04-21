package integrationtest

import (
	"github.com/deck-of-cards/src/api"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestPing(t *testing.T) {
	req, _ := http.NewRequest("GET", "/ping", nil)

	handler := http.HandlerFunc(api.Ping)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	checkResponseCode(t, http.StatusOK, response.Code)

}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}