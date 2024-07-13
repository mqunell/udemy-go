package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected first card to be Ace of Clubs, but got %v", d[0])
	}

	if d[15] != "Four of Spades" {
		t.Errorf("Expected last card to be Four of Spades, but got %v", d[15])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	testFileName := "_decktesting.txt"
	os.Remove(testFileName)

	d := newDeck()
	d.saveToFile(testFileName)
	defer os.Remove(testFileName)

	d2 := newDeckFromFile(testFileName)
	if len(d2) != 16 {
		t.Errorf("Expected deck from file to have length of 16, but got %v", len(d2))
	}
}
