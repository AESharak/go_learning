package main

func main() {
	// Create a slice of strings called 'cards' with initial values
	// []string declares a slice that can hold multiple string values
	// We initialize it with two elements: a string literal and a function call
	cards := []string{"Ace of Diamongs", newCard()}

	// Iterate over the slice using 'range' keyword
	// 'i' gets the index (0, 1, 2, etc.) and 'card' gets the actual value
	// This is the first iteration, so we'll see the original 2 cards
	for i, card := range cards {
		println("first one", i, card)
	}

	// Use the append function to add a new element to the slice
	// append() takes the existing slice and the new element(s) to add
	// IMPORTANT: append() returns a NEW slice, it doesn't modify the original
	// We must assign the result back to 'cards' to update our slice
	cards = append(cards, "Ace of Spades")

	// Iterate over the slice again after adding the new card
	// Now we should see 3 cards total (original 2 + the new "Ace of Spades")
	for i, card := range cards {
		println("second time", i, card)
	}
}

// newCard function returns a string value
// The "string" after the parentheses tells Go this function returns a string type
func newCard() string {
	// Return a string value - this must match the declared return type
	return "Five of Diamonds"
}
