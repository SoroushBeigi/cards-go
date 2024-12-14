package main

func main() {
	cards := newDeck()
	deck1, deck2 := deal(cards, 5)
	deck1.print()
	deck2.print()
}
