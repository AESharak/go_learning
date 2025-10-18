package main

import "fmt"

// Create a new type of deck, which is a slice of strings
// This custom type "extends" or "borrows" all the behavior of a slice of string
// If you're coming from object-oriented programming, you can think of this as:
// "deck type extends slice of string" or "deck is a subclass of slice of string"
// However, Go doesn't use terms like "extends" or "subclass" - this is just for understanding
type deck []string

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
