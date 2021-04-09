package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but got %d", len(d))
	}

	if d[0] != "Spades of Two" {
		t.Errorf("Expected Spades of Two but got %v", d[1])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected Four of Clubs but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()

	deck.saveToFile("_decktesting")
	loadedDeck := readFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards but got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
