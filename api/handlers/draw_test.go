package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestDrawDeckWithoutCardsWithWrongDeckIdAndCorrectCount(t *testing.T) {

	count := 2

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

	drawDeckReq, _ := http.NewRequest("GET", "/deck/draw", nil)
	drawDeckQuery := drawDeckReq.URL.Query()
	drawDeckQuery.Add("deck_id", "wrong_deck_id")
	drawDeckQuery.Add("count", strconv.Itoa(count))
	drawDeckReq.URL.RawQuery = drawDeckQuery.Encode()

	handler = http.HandlerFunc(Draw)
	response = httptest.NewRecorder()
	handler.ServeHTTP(response, drawDeckReq)
	checkResponseCode(t, http.StatusBadRequest, response.Code)

}

func TestDrawDeckWithoutCardsWithWrongDeckIdAndWrongCount(t *testing.T) {

	count := 200

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

	drawDeckReq, _ := http.NewRequest("GET", "/deck/draw", nil)
	drawDeckQuery := drawDeckReq.URL.Query()
	drawDeckQuery.Add("deck_id", "wrong_deck_id")
	drawDeckQuery.Add("count", strconv.Itoa(count))
	drawDeckReq.URL.RawQuery = drawDeckQuery.Encode()

	handler = http.HandlerFunc(Draw)
	response = httptest.NewRecorder()
	handler.ServeHTTP(response, drawDeckReq)
	checkResponseCode(t, http.StatusBadRequest, response.Code)

}

func TestDrawDeckWithoutCardsWithWrongDeckIdAndNegativeCount(t *testing.T) {

	count := -2

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

	drawDeckReq, _ := http.NewRequest("GET", "/deck/draw", nil)
	drawDeckQuery := drawDeckReq.URL.Query()
	drawDeckQuery.Add("deck_id", "wrong_deck_id")
	drawDeckQuery.Add("count", strconv.Itoa(count))
	drawDeckReq.URL.RawQuery = drawDeckQuery.Encode()

	handler = http.HandlerFunc(Draw)
	response = httptest.NewRecorder()
	handler.ServeHTTP(response, drawDeckReq)
	checkResponseCode(t, http.StatusBadRequest, response.Code)

}

func TestDrawWithoutCardsWithCorrectDeckIdAndCount(t *testing.T) {

	count := 2

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

	drawDeckRes := new(DrawDeckResponse)
	drawDeckReq, _ := http.NewRequest("GET", "/deck/draw", nil)
	drawDeckQuery := drawDeckReq.URL.Query()
	drawDeckQuery.Add("deck_id", createDeckRes.DeckID)
	drawDeckQuery.Add("count", strconv.Itoa(count))
	drawDeckReq.URL.RawQuery = drawDeckQuery.Encode()

	handler = http.HandlerFunc(Draw)
	response = httptest.NewRecorder()
	handler.ServeHTTP(response, drawDeckReq)
	checkResponseCode(t, http.StatusOK, response.Code)

	str = response.Body.String()
	err = json.Unmarshal([]byte(str), &drawDeckRes)
	if err != nil {
		t.Errorf("Incorrect Response %s\n", str)
	}

	if len(drawDeckRes.Cards) != count {
		t.Errorf("Response cards length not matching with requested lenght")
	}

	if drawDeckRes.Cards[0].Code == "" || drawDeckRes.Cards[0].Suit == "" || drawDeckRes.Cards[0].Value == "" {
		t.Errorf("Card code, suit or value empty, exptected non empty")
	}
}

func TestDrawDeckWithoutCardsWithCorrectDeckIdAndWrongCount(t *testing.T) {

	count := 200

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

	drawDeckReq, _ := http.NewRequest("GET", "/deck/draw", nil)
	drawDeckQuery := drawDeckReq.URL.Query()
	drawDeckQuery.Add("deck_id", createDeckRes.DeckID)
	drawDeckQuery.Add("count", strconv.Itoa(count))
	drawDeckReq.URL.RawQuery = drawDeckQuery.Encode()

	handler = http.HandlerFunc(Draw)
	response = httptest.NewRecorder()
	handler.ServeHTTP(response, drawDeckReq)
	checkResponseCode(t, http.StatusBadRequest, response.Code)

}

func TestDrawDeckWithoutCardsWithCorrectDeckIdAndNegativeCount(t *testing.T) {

	count := -2

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

	drawDeckReq, _ := http.NewRequest("GET", "/deck/draw", nil)
	drawDeckQuery := drawDeckReq.URL.Query()
	drawDeckQuery.Add("deck_id", createDeckRes.DeckID)
	drawDeckQuery.Add("count", strconv.Itoa(count))
	drawDeckReq.URL.RawQuery = drawDeckQuery.Encode()

	handler = http.HandlerFunc(Draw)
	response = httptest.NewRecorder()
	handler.ServeHTTP(response, drawDeckReq)
	checkResponseCode(t, http.StatusBadRequest, response.Code)

}
