package main

func main() {
	cards := newDeck()
	cards.saveToFile("myCard")
}

func newCard() string {
	return "Five of Diamonds"
}
