package main

import "fmt"

func main() {
	cards := newDeckFromFile("my_cards")
	cards.print()

	fmt.Println("==================")
	cards.shuffle()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
