package main

import "fmt"

// Create a new type of deck, which is a slice of strings
// This custom type "extends" or "borrows" all the behavior of a slice of string
// If you're coming from object-oriented programming, you can think of this as:
// "deck type extends slice of string" or "deck is a subclass of slice of string"
// However, Go doesn't use terms like "extends" or "subclass" - this is just for understanding
type deck []string

// newDeck creates and returns a complete deck of playing cards
// This function does NOT have a receiver because it creates a new deck from scratch
// You would call it like: cards := newDeck()
// It generates all possible combinations of suits and values
func newDeck() deck {
	// Create an empty deck to start with
	// deck{} creates a new slice of strings with zero length
	cards := deck{}

	// Define all possible card suits
	// We use []string (not deck) because these are just suit names, not complete cards
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	// Define all possible card values
	// Starting with Ace, Two, Three, Four (can be extended to include more values)
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// Nested loops to create every combination of suit and value
	// First loop: iterate through each suit (Spades, Diamonds, Hearts, Clubs)
	for _, suit := range cardSuits {
		// Second loop: iterate through each value for the current suit
		// This creates 4 values × 4 suits = 16 total cards
		for _, value := range cardValues {
			// Create a card name by combining value and suit with " of " in between
			// Example: "Ace" + " of " + "Spades" = "Ace of Spades"
			// We use value + " of " + suit to get natural card naming
			cards = append(cards, value+" of "+suit)
		}
	}

	// Return the complete deck with all 16 card combinations
	return cards
}

// This is a special function called a "receiver function" or "method"
// The syntax (d deck) before the function name is called a "receiver"
// This means the print() function belongs to the deck type
// You can think of it as: "any variable of type deck now has a print() method"
// The 'd' is the receiver variable - it represents the instance of deck that calls this method
// When you call cards.print(), 'd' will be equal to 'cards'
func (d deck) print() {
	// Loop through all cards in the deck and print each one
	// 'i' gets the index (0, 1, 2, etc.) and 'card' gets the actual value
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// deal function splits a deck into two separate decks: a hand and remaining cards
// This function demonstrates multiple return values and slice range syntax in Go
// Parameters:
//   - d deck: the original deck of cards to split
//   - handSize int: number of cards to deal out into the hand
// Returns:
//   - deck: the hand containing the specified number of cards
//   - deck: the remaining cards left in the deck
// 
// Example usage: hand, remaining := deal(cards, 5)
// This would create a hand of 5 cards and return the rest as remaining deck
func deal(d deck, handSize int) (deck, deck) {
	// Use slice range syntax to split the deck into two parts:
	// d[:handSize] - everything from start (index 0) up to handSize (not including handSize)
	// d[handSize:] - everything from handSize index to the end of the slice
	// 
	// Example: if handSize = 4 and deck has 16 cards:
	// d[:4] returns cards at indices 0, 1, 2, 3 (first 4 cards)
	// d[4:] returns cards at indices 4, 5, 6, ..., 15 (remaining 12 cards)
	return d[:handSize], d[handSize:]
}
