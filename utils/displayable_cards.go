package utils

import (
	"github.com/deck-of-cards/constants"
)

type Card struct {
	Code  string `json:"code"`
	Suit  string `json:"suit"`
	Value string `json:"value"`
}

func GetDisplayableCards(cards []string) []*Card {
	var displayableCards []*Card
	for _, v := range cards {
		displayableCard := new(Card)
		displayableCard.Code = v
		displayableCard.Suit = constants.SUITS[v[1:]]
		displayableCard.Value = getValue(v[:1])
		displayableCards = append(displayableCards, displayableCard)
	}
	return displayableCards
}
