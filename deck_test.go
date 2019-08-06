package main

import (
	"os"
	"testing"
)

// newDeck function should create a deck length of 48 items, first card should be equal to Ace of Spades, and last card should be equal to King of Clubs

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 48 {
		t.Errorf("Expected deck length of 48 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected first card of King of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 48 {
		t.Errorf("Expected 48 cards, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
