package main

import "fmt"



func main() {
	// 	the longer syntax that is used only for documentation
	//  var card string = "Ace of Spades"

	//  the alternative more short and concise and 100% similar to the above code
	//  this valid for the initialization only
	// 	when you reassign it to a new value, you mustn't use := anymore for the same value
	card := "Ace of Spades"

	//  reassinging
	card = "Five of Daimonds"

	// card := "Five of Diamonds" // ERROR, No new varaible declared
	fmt.Println(card)
}
