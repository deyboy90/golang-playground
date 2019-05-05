package main

func main() {
	cards := newDeck()
	deal, _ := cards.deal(5)
	deal.print()

	dealShuffle := deal.shuffle()
	dealShuffle.print()

	// deal.saveToFile("my_deal")
	// dealFromFile := newDeckFromFile("my_deal")
	// dealFromFile.print()
	// cards.print()
}
