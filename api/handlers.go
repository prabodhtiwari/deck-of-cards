package api

import(
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	guuid "github.com/google/uuid"
	"github.com/deck-of-cards/constants"
	"github.com/deck-of-cards/utils"
)

type CreateDeckResponse struct {
	DeckID    string        `json:"deck_id"`
	Shuffled  bool          `json:"shuffled"`
	Remaining int           `json:"remaining"`
	Cards     []*utils.Card `json:"-"`
}

type OpenDeckResponse struct {
	DeckID    string        `json:"deck_id"`
	Shuffled  bool          `json:"shuffled"`
	Remaining int           `json:"remaining"`
	Cards     []*utils.Card `json:"cards"`
}

var Decks = map[string]CreateDeckResponse{}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong"))
}

func Create(w http.ResponseWriter, r *http.Request) {
	shuffle, _ := strconv.ParseBool(r.URL.Query().Get("shuffle"))

	var wantedCards []string
	if r.URL.Query().Get("cards") != "" {
		wantedCards = strings.Split(r.URL.Query().Get("cards"), ",")
	}

	cardCheck := utils.CheckCardValidity(wantedCards)
	if cardCheck != true {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	wantedCardsLen := len(wantedCards)

	deckID := guuid.New().String()
	var cards []string

	if wantedCardsLen > 0 {
		cards = make([]string, wantedCardsLen)
		copy(cards, wantedCards)
	} else {
		cards = make([]string, 52)
		copy(cards, constants.CARDS)
	}

	if shuffle == true {
		utils.Shuffle(cards)
	}

	response := CreateDeckResponse{DeckID: deckID, Shuffled: shuffle, Remaining: len(cards), Cards: utils.GetDisplayableCards(cards)}
	Decks[deckID] = response
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func Open(w http.ResponseWriter, r *http.Request) {
	deckID := r.URL.Query().Get("deck_id")

	deck := Decks[deckID]

	if deck.DeckID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := OpenDeckResponse{DeckID: deck.DeckID, Shuffled: deck.Shuffled, Remaining: deck.Remaining, Cards: deck.Cards}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}