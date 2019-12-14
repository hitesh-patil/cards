package main

import "testing"

func TestNewCards(t *testing.T) {
	cards := newCards()

	if len(cards) != 52 {
		t.Errorf("Expected 52 cards, but got %v", len(cards))
	}
}
