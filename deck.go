package main

import "fmt"

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
