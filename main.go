package main

func main() {
	cards := newDeck()

	cards = cards.shuffle()
	cards.print()

}
