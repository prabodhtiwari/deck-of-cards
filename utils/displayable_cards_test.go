package utils

import (
	"testing"
)

func TestGetDisplayableCards(t *testing.T) {

	cards := []string{"AS", "KH"}
	res := GetDisplayableCards(cards)

	if res[0].Code != "AS" || res[0].Suit != "SPADES" || res[0].Value != "ACE" {
		t.Errorf("displayable card response is not matching expected response")
	}

	if res[1].Code != "KH" || res[1].Suit != "HEARTS" || res[1].Value != "KING" {
		t.Errorf("displayable card response is not matching expected response")
	}

}
