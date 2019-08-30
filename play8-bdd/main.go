package main

import "github.com/golang-playground/play8-bdd/deck"

func main() {
	cards := deck.NewDeck()
	deal, _ := cards.deal(5)
	deal.print()

	dealShuffle := deal.shuffle()
	dealShuffle.print()

	// deal.saveToFile("my_deal")
	// dealFromFile := newDeckFromFile("my_deal")
	// dealFromFile.print()
	// cards.print()
}
