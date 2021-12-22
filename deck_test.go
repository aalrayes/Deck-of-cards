package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if d[0] != "Ace of Spades" {
		t.Error("Expected Ace of Spades got, ", d[0])
	}
	if d[len(d)-1] != "Six of Diamonds" {
		t.Error("Expected Six of Diamonds got, ", d[len(d)-1])
	}
	if len(d) != 24 {
		t.Error("Expected Deck length of 24", len(d))
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 24 {
		t.Error("Expected Deck of length 24 got, ", len(loadedDeck))
	}
	os.Remove("_decktesting")

}
