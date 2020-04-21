package unittest

import (
	"os"
	"reflect"
	"testing"

	"github.com/deck-of-cards/utils"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestShuffle(t *testing.T) {

	cards := []string{"AS", "2S", "3S", "4S", "5S", "6S", "7S", "8S", "9S", "0S"}
	shuffledCards := make([]string, len(cards))
	copy(shuffledCards, cards)
	utils.Shuffle(shuffledCards)

	if len(cards) != len(shuffledCards) {
		t.Errorf("Shuffled cards length not matching")
	}

	if reflect.DeepEqual(cards, shuffledCards) {
		t.Errorf("Cards not shuffled")
	}

}
