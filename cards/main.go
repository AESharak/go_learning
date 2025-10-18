package main

import "fmt"

func main() {
	// Call the newCard function and assign its return value to card
	// Go infers that card is of type string because newCard() returns a string
	card := newCard()

	// Print the card value to the console
	fmt.Println(card)
}

// newCard function returns a string value
// The "string" after the parentheses tells Go this function returns a string type
func newCard() string {
	// Return a string value - this must match the declared return type
	return "Five of Diamonds"
}