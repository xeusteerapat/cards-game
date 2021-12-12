package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type of 'deck' which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Any variables of type 'deck' now gets access to the 'print' method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

/* NOTE:
slice[indexIncluding : upToNotIncluding]
slice[:2] give me everything from 0 -> 1
slice[3:] give me everything from 3 -> end
*/

// Return multiple values with all types of deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Type conversion []byte(someString) string -> []byte
// deck -> []string -> string -> []byte
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
