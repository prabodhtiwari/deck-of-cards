package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/deck-of-cards/utils"
)

type DrawDeckResponse struct {
	Cards []*utils.Card `json:"cards"`
}

func Draw(w http.ResponseWriter, r *http.Request) {
	deckID := r.URL.Query().Get("deck_id")
	count, err := strconv.Atoi(r.URL.Query().Get("count"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	deck := Decks[deckID]

	if deck.DeckID == "" || count > len(deck.Cards) || count < 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := DrawDeckResponse{Cards: deck.Cards[:count]}

	deck.Cards = deck.Cards[count:]
	deck.Remaining = len(deck.Cards)

	Decks[deckID] = deck

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
