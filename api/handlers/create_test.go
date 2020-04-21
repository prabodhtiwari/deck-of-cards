package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestCreateDeckWithoutCards(t *testing.T) {
	res := new(CreateDeckResponse)
	req, _ := http.NewRequest("GET", "/deck/create", nil)

	handler := http.HandlerFunc(Create)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	checkResponseCode(t, http.StatusOK, response.Code)

	str := response.Body.String()
	err := json.Unmarshal([]byte(str), &res)
	if err != nil {
		t.Errorf("Incorrect Response %s\n", str)
	}

	if res.DeckID == "" {
		t.Errorf("Empty deck id expected non empty deck id")
	}
	if res.Remaining != 52 {
		t.Errorf("Remaining cards length not matching with expected lenght 52")
	}
	if res.Shuffled {
		t.Errorf("response of shuffled is true expected false")
	}
}

func TestCreateDeckWithoutCardsAndWithShuffle(t *testing.T) {

	res := new(CreateDeckResponse)
	req, _ := http.NewRequest("GET", "/deck/create", nil)
	q := req.URL.Query()
	q.Add("shuffle", "true")
	req.URL.RawQuery = q.Encode()

	handler := http.HandlerFunc(Create)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	checkResponseCode(t, http.StatusOK, response.Code)

	str := response.Body.String()
	err := json.Unmarshal([]byte(str), &res)
	if err != nil {
		t.Errorf("Incorrect Response %s\n", str)
	}

	if res.DeckID == "" {
		t.Errorf("Empty deck id expected non empty deck id")
	}
	if res.Remaining != 52 {
		t.Errorf("Remaining cards length not matching with expected lenght 52")
	}
	if !res.Shuffled {
		t.Errorf("response of shuffled is false expected true")
	}
}

func TestCreateDeckWithWrongCards(t *testing.T) {

	cards := "AS,KD,AC,2C,KH,LL"
	req, _ := http.NewRequest("GET", "/deck/create", nil)
	q := req.URL.Query()
	q.Add("cards", cards)
	req.URL.RawQuery = q.Encode()

	handler := http.HandlerFunc(Create)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)

}

func TestCreateDeckWithDuplicateCards(t *testing.T) {

	cards := "AS,KD,AC,2C,KH,KH"
	req, _ := http.NewRequest("GET", "/deck/create", nil)
	q := req.URL.Query()
	q.Add("cards", cards)
	req.URL.RawQuery = q.Encode()

	handler := http.HandlerFunc(Create)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)

}

func TestCreateDeckWithCards(t *testing.T) {

	cards := "AS,KD,AC,2C,KH"
	cardsSlice := strings.Split(cards, ",")

	res := new(CreateDeckResponse)
	req, _ := http.NewRequest("GET", "/deck/create", nil)
	q := req.URL.Query()
	q.Add("cards", cards)
	req.URL.RawQuery = q.Encode()

	handler := http.HandlerFunc(Create)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	checkResponseCode(t, http.StatusOK, response.Code)

	str := response.Body.String()
	err := json.Unmarshal([]byte(str), &res)
	if err != nil {
		t.Errorf("Incorrect Response %s\n", str)
	}

	if res.DeckID == "" {
		t.Errorf("Empty deck id expected non empty deck id")
	}
	if res.Remaining != len(cardsSlice) {
		t.Errorf("Remaining cards length not matching with expected lenght 52")
	}
	if res.Shuffled {
		t.Errorf("response of shuffled is true expected false")
	}
}

func TestCreateDeckWithCardsAndWithShuffle(t *testing.T) {

	cards := "AS,KD,AC,2C,KH"
	cardsSlice := strings.Split(cards, ",")

	res := new(CreateDeckResponse)
	req, _ := http.NewRequest("GET", "/deck/create", nil)
	q := req.URL.Query()
	q.Add("cards", cards)
	q.Add("shuffle", "true")
	req.URL.RawQuery = q.Encode()

	handler := http.HandlerFunc(Create)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	checkResponseCode(t, http.StatusOK, response.Code)

	str := response.Body.String()
	err := json.Unmarshal([]byte(str), &res)
	if err != nil {
		t.Errorf("Incorrect Response %s\n", str)
	}

	if res.DeckID == "" {
		t.Errorf("Empty deck id expected non empty deck id")
	}
	if res.Remaining != len(cardsSlice) {
		t.Errorf("Remaining cards length not matching with expected lenght 52")
	}
	if !res.Shuffled {
		t.Errorf("response of shuffled is false expected true")
	}
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
