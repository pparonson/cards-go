package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	// newDeck() should return 52 cards
	if len(d) != 52 {
		t.Errorf("Expected 52, but got: %v", len(d))
	}
	// newDeck() should return Ace of Spades at [0] and King of Clubs at [51]
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, but got: %v", d[0])
	}
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected King of Clubs, but got: %v", d[len(d)-1])
	}
}

// func (d deck) TestPrint() {

// }

// func TestDeal(d deck, handSize int) (deck, deck) {

// }

func TestSaveToFileAndReadFromFile(t *testing.T) {
	// clean up to remove any testDeck from project dir
	os.Remove("_testDeck")

	// create a new deck for tests
	createdDeck := newDeck()

	// save deck to local storage
	createdDeck.saveToFile("_testDeck")

	// read deck from localstorage
	savedDeck := readFromFile("_testDeck")

	// assert newly created deck exists
	if len(savedDeck) != 52 {
		t.Errorf("Expected 52, but got: %v", len(savedDeck))
	}
}
