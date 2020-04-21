package unittest

import (
	"testing"

	"github.com/deck-of-cards/utils"
)

func TestCheckCardValidityWithCorrectCards(t *testing.T) {

	cards := []string{"AS", "2S", "3S", "4S", "5S", "6S", "7S", "8S", "9S", "0S"}
	res := utils.CheckCardValidity(cards)

	if res != true {
		t.Errorf("correct cards check not passed")
	}

}

func TestCheckCardValidityWithDuplicateCorrectCards(t *testing.T) {

	cards := []string{"AS", "2S", "3S", "4S", "5S", "6S", "7S", "8S", "9S", "0S", "0S"}
	res := utils.CheckCardValidity(cards)

	if res != false {
		t.Errorf("correct cards check not passed")
	}

}

func TestCheckCardValidityWithIncorrectCards(t *testing.T) {

	cards := []string{"AS", "2U", "3S", "4S", "5S", "6S", "7S", "8S", "9S", "0S"}
	res := utils.CheckCardValidity(cards)

	if res != false {
		t.Errorf("Incorrect cards check not passed")
	}

}
