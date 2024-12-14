package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	// Check if the deck has the correct number of cards (52 cards)
	if len(deck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(deck))
	}

	// Check if the first card is correct
	expectedFirstCard := "Spades of Ace"
	if deck[0] != expectedFirstCard {
		t.Errorf("Expected first card to be '%v', but got '%v'", expectedFirstCard, deck[0])
	}

	// Check if the last card is correct
	expectedLastCard := "Clubs of King"
	if deck[len(deck)-1] != expectedLastCard {
		t.Errorf("Expected last card to be '%v', but got '%v'", expectedLastCard, deck[len(deck)-1])
	}
}
