package main

import (
	"os"
	"reflect"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	expected := 52
	if len(d) != expected {
		t.Errorf("Expected: %v Actual: %v", expected, len(d))
	}
}

func TestNewDeckWithStartAndEndSlice(t *testing.T) {
	d := newDeck()

	expected := "Ace of Spades"
	if d[0] != expected {
		t.Errorf("Expected: %v Actual: %v", expected, d[0])
	}

	expected = "King of Clubs"
	if d[len(d)-1] != expected {
		t.Errorf("Expected: %v Actual: %v", expected, d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	testFileName := "_deckTest"
	os.Remove(testFileName)

	deck := newDeck()
	deck.saveToFile(testFileName)
	loadedDeck := newDeckFromFile(testFileName)

	if reflect.DeepEqual(deck, loadedDeck) == false {
		t.Errorf("Loaded deck does not match created deck")
	}
	os.Remove(testFileName)
}
