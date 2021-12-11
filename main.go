package main

import "fmt"

func main() {
	cards := newDeck()

	cards.print()
	hand, remainingDeck := deal(cards, 5)

	fmt.Println("======================")

	hand.print()
	fmt.Println("======================")
	remainingDeck.print()
}

func newCard() string {
	return "Five of Diamonds"
}
