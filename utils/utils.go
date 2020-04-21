package utils

import (
	"math/rand"
	"time"

	"github.com/deck-of-cards/constants"

)

type Card struct {
	Code  string `json:"code"`
	Suit  string `json:"suit"`
	Value string `json:"value"`
}

func CheckCardValidity(cards []string) bool {
	var cardCodeMap = make(map[string]bool)
	for _, card := range constants.CARDS {
		cardCodeMap[card] = true
	}

	for _, card := range cards {
		if cardCodeMap[card] == false {
			return false
		} else {
			cardCodeMap[card] = false
		}
	}
	return true
}

func Shuffle(cards []string) {
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(cards), func(i int, j int) { cards[i], cards[j] = cards[j], cards[i] })
}


func getValue(valueCode string) string {
	if constants.VALUES[valueCode] != "" {
		return constants.VALUES[valueCode]
	}
	return valueCode
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
