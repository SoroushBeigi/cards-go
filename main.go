package main

import "fmt"

func main() {
	cards := newDeck()
	deck1, _ := deal(cards, 4)
	fmt.Println(deck1.toString())
}
