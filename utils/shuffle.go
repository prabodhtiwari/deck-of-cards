package utils

import (
	"math/rand"
	"time"

	"github.com/deck-of-cards/constants"
)

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
