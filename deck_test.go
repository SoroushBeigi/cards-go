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

func TestDeal(t *testing.T) {
	deck := newDeck()
	handSize := 5
	hand, remainingDeck := deal(deck, handSize)
	if len(hand) != handSize {
		t.Errorf("Expected hand size of %v, but got %v", handSize, len(hand))
	}
	expectedRemainingDeckSize := len(deck) - handSize
	if len(remainingDeck) != expectedRemainingDeckSize {
		t.Errorf("Expected remaining deck size of %v, but got %v", expectedRemainingDeckSize, len(remainingDeck))
	}

	// Ensure that the deck was split correctly (i.e., no card is repeated)
	for i := 0; i < handSize; i++ {
		if hand[i] == remainingDeck[i] {
			t.Errorf("Card '%v' was found in both hand and remaining deck at position %v", hand[i], i)
		}
	}
}
