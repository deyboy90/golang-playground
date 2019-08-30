package deck

import (
	. "github.com/onsi/ginkgo"
	"os"
	"reflect"
)

var _ = Describe("Testing with Ginkgo", func() {
	It("new deck", func() {

		d := NewDeck()

		expected := 52
		if len(d) != expected {
			GinkgoT().Errorf("Expected: %v Actual: %v", expected, len(d))
		}
	})
	It("new deck with start and end slice", func() {

		d := NewDeck()

		expected := "Ace of Spades"
		if d[0] != expected {
			GinkgoT().Errorf("Expected: %v Actual: %v", expected, d[0])
		}

		expected = "King of Clubs"
		if d[len(d)-1] != expected {
			GinkgoT().Errorf("Expected: %v Actual: %v", expected, d[len(d)-1])
		}
	})
	It("save to deck and new deck from file", func() {

		testFileName := "_deckTest"
		os.Remove(testFileName)

		deck := NewDeck()
		deck.saveToFile(testFileName)
		loadedDeck := newDeckFromFile(testFileName)

		if reflect.DeepEqual(deck, loadedDeck) == false {
			GinkgoT().Errorf("Loaded deck does not match created deck")
		}
		os.Remove(testFileName)
	})
})
