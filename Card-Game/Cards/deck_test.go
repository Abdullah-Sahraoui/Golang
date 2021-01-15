package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// To ensure length of deck is 16
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	// To see if first and last cards are ace of spades and four of clubs respectively
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	// To ensure saveToFile() and newDeckFromFile() work properly
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
