package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/deck-of-cards/constants"
	"github.com/deck-of-cards/utils"
)

func TestOpenDeckWithoutCardsWithWrongDeckId(t *testing.T) {

	createDeckRes := new(CreateDeckResponse)
	createDeckReq, _ := http.NewRequest("GET", "/deck/create", nil)

	handler := http.HandlerFunc(Create)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, createDeckReq)
	checkResponseCode(t, http.StatusOK, response.Code)

	str := response.Body.String()
	err := json.Unmarshal([]byte(str), &createDeckRes)
	if err != nil {
		t.Errorf("Incorrect Response %s\n", str)
	}

	openDeckReq, _ := http.NewRequest("GET", "/deck/open", nil)
	openDeckQuery := openDeckReq.URL.Query()
	openDeckQuery.Add("deck_id", "wrong_value")
	openDeckReq.URL.RawQuery = openDeckQuery.Encode()

	handler = http.HandlerFunc(Open)
	response = httptest.NewRecorder()
	handler.ServeHTTP(response, openDeckReq)
	checkResponseCode(t, http.StatusBadRequest, response.Code)

}

func TestOpenDeckWithoutCardsWithEmptyDeckId(t *testing.T) {

	createDeckRes := new(CreateDeckResponse)
	createDeckReq, _ := http.NewRequest("GET", "/deck/create", nil)

	handler := http.HandlerFunc(Create)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, createDeckReq)
	checkResponseCode(t, http.StatusOK, response.Code)

	str := response.Body.String()
	err := json.Unmarshal([]byte(str), &createDeckRes)
	if err != nil {
		t.Errorf("Incorrect Response %s\n", str)
	}

	openDeckReq, _ := http.NewRequest("GET", "/deck/open", nil)
	openDeckQuery := openDeckReq.URL.Query()
	openDeckQuery.Add("deck_id", "")
	openDeckReq.URL.RawQuery = openDeckQuery.Encode()

	handler = http.HandlerFunc(Open)
	response = httptest.NewRecorder()
	handler.ServeHTTP(response, openDeckReq)
	checkResponseCode(t, http.StatusBadRequest, response.Code)

}

func TestOpenDeckWithoutCards(t *testing.T) {

	createDeckRes := new(CreateDeckResponse)
	createDeckReq, _ := http.NewRequest("GET", "/deck/create", nil)

	handler := http.HandlerFunc(Create)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, createDeckReq)
	checkResponseCode(t, http.StatusOK, response.Code)

	str := response.Body.String()
	err := json.Unmarshal([]byte(str), &createDeckRes)
	if err != nil {
		t.Errorf("Incorrect Response %s\n", str)
	}

	openDeckRes := new(OpenDeckResponse)
	openDeckReq, _ := http.NewRequest("GET", "/deck/open", nil)
	openDeckQuery := openDeckReq.URL.Query()
	openDeckQuery.Add("deck_id", createDeckRes.DeckID)
	openDeckReq.URL.RawQuery = openDeckQuery.Encode()

	handler = http.HandlerFunc(Open)
	response = httptest.NewRecorder()
	handler.ServeHTTP(response, openDeckReq)
	checkResponseCode(t, http.StatusOK, response.Code)

	str = response.Body.String()
	err = json.Unmarshal([]byte(str), &openDeckRes)
	if err != nil {
		t.Errorf("Incorrect Response %s\n", str)
	}

	if openDeckRes.DeckID == "" {
		t.Errorf("Empty deck id expected non empty deck id")
	}
	if openDeckRes.Remaining != len(constants.CARDS) {
		t.Errorf("Remaining cards length not matching with expected lenght 52")
	}
	if openDeckRes.Shuffled {
		t.Errorf("response of shuffled is true expected false")
	}
	if !reflect.DeepEqual(openDeckRes.Cards, utils.GetDisplayableCards(constants.CARDS)) {
		t.Errorf("response shuffled but expcted non shuffled")
	}
}

func TestOpenDeckWithoutCardsAndWithShuffle(t *testing.T) {

	createDeckRes := new(CreateDeckResponse)
	createDeckReq, _ := http.NewRequest("GET", "/deck/create", nil)
	createDeckQuery := createDeckReq.URL.Query()
	createDeckQuery.Add("shuffle", "true")
	createDeckReq.URL.RawQuery = createDeckQuery.Encode()

	handler := http.HandlerFunc(Create)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, createDeckReq)
	checkResponseCode(t, http.StatusOK, response.Code)

	str := response.Body.String()
	err := json.Unmarshal([]byte(str), &createDeckRes)
	if err != nil {
		t.Errorf("Incorrect Response %s\n", str)
	}

	openDeckRes := new(OpenDeckResponse)
	openDeckReq, _ := http.NewRequest("GET", "/deck/open", nil)
	openDeckQuery := openDeckReq.URL.Query()
	openDeckQuery.Add("deck_id", createDeckRes.DeckID)
	openDeckReq.URL.RawQuery = openDeckQuery.Encode()

	handler = http.HandlerFunc(Open)
	response = httptest.NewRecorder()
	handler.ServeHTTP(response, openDeckReq)
	checkResponseCode(t, http.StatusOK, response.Code)

	str = response.Body.String()
	err = json.Unmarshal([]byte(str), &openDeckRes)
	if err != nil {
		t.Errorf("Incorrect Response %s\n", str)
	}

	if openDeckRes.DeckID == "" {
		t.Errorf("Empty deck id expected non empty deck id")
	}
	if openDeckRes.Remaining != len(constants.CARDS) {
		t.Errorf("Remaining cards length not matching with expected lenght 52")
	}
	if !openDeckRes.Shuffled {
		t.Errorf("response of shuffled is true expected false")
	}
	if reflect.DeepEqual(openDeckRes.Cards, utils.GetDisplayableCards(constants.CARDS)) {
		t.Errorf("response not shuffled but expcted shuffled")
	}
}

func TestOpenDeckWithCards(t *testing.T) {

	cards := "AS,KD,AC,2C,KH"
	cardsSlice := strings.Split(cards, ",")

	createDeckRes := new(CreateDeckResponse)
	createDeckReq, _ := http.NewRequest("GET", "/deck/create", nil)
	createDeckQuery := createDeckReq.URL.Query()
	createDeckQuery.Add("cards", cards)
	createDeckReq.URL.RawQuery = createDeckQuery.Encode()

	handler := http.HandlerFunc(Create)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, createDeckReq)
	checkResponseCode(t, http.StatusOK, response.Code)

	str := response.Body.String()
	err := json.Unmarshal([]byte(str), &createDeckRes)
	if err != nil {
		t.Errorf("Incorrect Response %s\n", str)
	}

	openDeckRes := new(OpenDeckResponse)
	openDeckReq, _ := http.NewRequest("GET", "/deck/open", nil)
	openDeckQuery := openDeckReq.URL.Query()
	openDeckQuery.Add("deck_id", createDeckRes.DeckID)
	openDeckReq.URL.RawQuery = openDeckQuery.Encode()

	handler = http.HandlerFunc(Open)
	response = httptest.NewRecorder()
	handler.ServeHTTP(response, openDeckReq)
	checkResponseCode(t, http.StatusOK, response.Code)

	str = response.Body.String()
	err = json.Unmarshal([]byte(str), &openDeckRes)
	if err != nil {
		t.Errorf("Incorrect Response %s\n", str)
	}

	if openDeckRes.DeckID == "" {
		t.Errorf("Empty deck id expected non empty deck id")
	}
	if openDeckRes.Remaining != len(cardsSlice) {
		t.Errorf("Remaining cards length not matching with expected lenght 52")
	}
	if openDeckRes.Shuffled {
		t.Errorf("response of shuffled is true expected false")
	}
	if !reflect.DeepEqual(openDeckRes.Cards, utils.GetDisplayableCards(cardsSlice)) {
		t.Errorf("response shuffled but expcted non shuffled")
	}

}

func TestOpenDeckWithCardsAndShuffle(t *testing.T) {

	cards := "AS,KD,AC,2C,KH"
	cardsSlice := strings.Split(cards, ",")

	createDeckRes := new(CreateDeckResponse)
	createDeckReq, _ := http.NewRequest("GET", "/deck/create", nil)
	createDeckQuery := createDeckReq.URL.Query()
	createDeckQuery.Add("cards", cards)
	createDeckQuery.Add("shuffle", "true")
	createDeckReq.URL.RawQuery = createDeckQuery.Encode()

	handler := http.HandlerFunc(Create)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, createDeckReq)
	checkResponseCode(t, http.StatusOK, response.Code)

	str := response.Body.String()
	err := json.Unmarshal([]byte(str), &createDeckRes)
	if err != nil {
		t.Errorf("Incorrect Response %s\n", str)
	}

	openDeckRes := new(OpenDeckResponse)
	openDeckReq, _ := http.NewRequest("GET", "/deck/open", nil)
	openDeckQuery := openDeckReq.URL.Query()
	openDeckQuery.Add("deck_id", createDeckRes.DeckID)
	openDeckReq.URL.RawQuery = openDeckQuery.Encode()

	handler = http.HandlerFunc(Open)
	response = httptest.NewRecorder()
	handler.ServeHTTP(response, openDeckReq)
	checkResponseCode(t, http.StatusOK, response.Code)

	str = response.Body.String()
	err = json.Unmarshal([]byte(str), &openDeckRes)
	if err != nil {
		t.Errorf("Incorrect Response %s\n", str)
	}

	if openDeckRes.DeckID == "" {
		t.Errorf("Empty deck id expected non empty deck id")
	}
	if openDeckRes.Remaining != len(cardsSlice) {
		t.Errorf("Remaining cards length not matching with expected lenght 52")
	}
	if !openDeckRes.Shuffled {
		t.Errorf("response of shuffled is true expected false")
	}
	if reflect.DeepEqual(openDeckRes.Cards, utils.GetDisplayableCards(cardsSlice)) {
		t.Errorf("response not shuffled but expcted shuffled")
	}

}
