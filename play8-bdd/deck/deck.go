package deck

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

/*
  golang does not have the idea of objects, type is similar to a construct
  which can have receiver functions associated with it, it helps delegate functionality
*/
type deck []string

// this is a normal function which can be directly called from anywhere inside package main
func NewDeck() deck {
	var cardSuites = []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	var cardNumbers = []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	cardDeck := deck{}
	for _, suite := range cardSuites {
		for _, number := range cardNumbers {
			cardDeck = append(cardDeck, number+" of "+suite)
		}
	}
	return cardDeck
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) shuffle() deck {
	for i := range d {
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(len(d) - 1)
		d[i], d[r] = d[r], d[i]
	}
	return d
}

/*
  Any variable of type deck gets access to call this receiver function print()
  the usual convention is to have single letter name as the variable which gets passed into the receiver
  it's similar to the idea of this/self but go avoids those constructs
*/
func (d deck) print() {
	for i, card := range d {
		println(i, card)
	}
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	content, error := ioutil.ReadFile(filename)
	if error != nil {
		fmt.Println("Failed to load file from disk: ", error)
		os.Exit(1)
	}
	return deck(strings.Split(string(content), ","))
}

func (d deck) toString() string {
	// convert deck -> []string -> string
	return strings.Join([]string(d), ",")
}
