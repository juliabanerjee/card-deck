package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected first card to be Ace of Clubs, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Hearts" {
		t.Errorf("Expected last card to be King of Hearts, but got %v", d[len(d)-1])
	}
}

//test names clearly indicate what function is being tested
func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")

	deck := newDeck()
	deck.saveToFile("_deckTesting")

	loadedDeck := newDeckFromFile(("_deskTesting"))

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(loadedDeck))
	}

	os.Remove("_deckTesting")
}
