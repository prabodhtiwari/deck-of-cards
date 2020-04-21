package integrationtest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/deck-of-cards/api"
)


func TestCreateDeckWithoutCards(t *testing.T) {
	res := new(api.CreateDeckResponse)
	req, _ := http.NewRequest("GET", "/deck/create", nil)

	handler := http.HandlerFunc(api.Create)
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
	if res.Shuffled != false {
		t.Errorf("response of shuffled is ture expected false")
	}
}

func TestCreateDeckWithoutCardsAndWithShuffle(t *testing.T) {

	res := new(api.CreateDeckResponse)
	req, _ := http.NewRequest("GET", "/deck/create", nil)
	q := req.URL.Query()
	q.Add("shuffle", "true")
	req.URL.RawQuery = q.Encode()

	handler := http.HandlerFunc(api.Create)
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
	if res.Shuffled != true {
		t.Errorf("response of shuffled is false expected true")
	}
}

func TestCreateDeckWithWrongCards(t *testing.T) {

	cards := "AS,KD,AC,2C,KH,LL"
	req, _ := http.NewRequest("GET", "/deck/create", nil)
	q := req.URL.Query()
	q.Add("cards", cards)
	req.URL.RawQuery = q.Encode()

	handler := http.HandlerFunc(api.Create)
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

	handler := http.HandlerFunc(api.Create)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)

}

func TestCreateDeckWithCards(t *testing.T) {

	cards := "AS,KD,AC,2C,KH"
	cardsSlice := strings.Split(cards, ",")

	res := new(api.CreateDeckResponse)
	req, _ := http.NewRequest("GET", "/deck/create", nil)
	q := req.URL.Query()
	q.Add("cards", cards)
	req.URL.RawQuery = q.Encode()

	handler := http.HandlerFunc(api.Create)
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
	if res.Shuffled != false {
		t.Errorf("response of shuffled is ture expected false")
	}
}

func TestCreateDeckWithCardsAndWithShuffle(t *testing.T) {

	cards := "AS,KD,AC,2C,KH"
	cardsSlice := strings.Split(cards, ",")

	res := new(api.CreateDeckResponse)
	req, _ := http.NewRequest("GET", "/deck/create", nil)
	q := req.URL.Query()
	q.Add("cards", cards)
	q.Add("shuffle", "true")
	req.URL.RawQuery = q.Encode()

	handler := http.HandlerFunc(api.Create)
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
	if res.Shuffled != true {
		t.Errorf("response of shuffled is false expected true")
	}
}

