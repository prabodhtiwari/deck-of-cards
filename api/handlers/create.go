package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/deck-of-cards/constants"
	"github.com/deck-of-cards/utils"
	guuid "github.com/google/uuid"
)

type CreateDeckResponse struct {
	DeckID    string        `json:"deck_id"`
	Shuffled  bool          `json:"shuffled"`
	Remaining int           `json:"remaining"`
	Cards     []*utils.Card `json:"-"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	shuffle, _ := strconv.ParseBool(r.URL.Query().Get("shuffle"))

	var wantedCards []string
	if r.URL.Query().Get("cards") != "" {
		wantedCards = strings.Split(r.URL.Query().Get("cards"), ",")
	}

	valid := utils.CheckCardValidity(wantedCards)
	if !valid {
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

	if shuffle {
		utils.Shuffle(cards)
	}

	response := CreateDeckResponse{DeckID: deckID, Shuffled: shuffle, Remaining: len(cards), Cards: utils.GetDisplayableCards(cards)}
	Decks[deckID] = response
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
