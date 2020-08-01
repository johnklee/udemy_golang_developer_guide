package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but got %v", len(d))
	}

	if d[0] != "Spades of Ace" {
		t.Errorf("Unexpected first card = %v", d[0])
	}

	if lastCard, err := d.lastElement(); lastCard != "Clubs of Four" || err != nil {
		t.Errorf("Unexpected last card = %v", lastCard)
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	test_file_name := "_decktesting"
	os.Remove(test_file_name)
	deck := newDeck()
	err := deck.saveToFile(test_file_name)
	if err != nil {
		t.Errorf("Something wrong in saving deck to file system")
	}
	loadedDeck, err := newDeckFromFile(test_file_name)
	if err != nil {
		t.Errorf("Something wrong in loading deck from file system")
	}

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck but got %v", len(loadedDeck))
	}
}
