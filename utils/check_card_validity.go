package utils

import "github.com/deck-of-cards/constants"

func CheckCardValidity(cards []string) bool {
	var cardCodeMap = make(map[string]bool)
	for _, card := range constants.CARDS {
		cardCodeMap[card] = true
	}

	for _, card := range cards {
		if !cardCodeMap[card] {
			return false
		}
		cardCodeMap[card] = false
	}
	return true
}
