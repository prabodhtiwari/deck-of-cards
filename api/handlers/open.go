package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/deck-of-cards/utils"
)

type OpenDeckResponse struct {
	DeckID    string        `json:"deck_id"`
	Shuffled  bool          `json:"shuffled"`
	Remaining int           `json:"remaining"`
	Cards     []*utils.Card `json:"cards"`
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
