package main

func main() {

	cards := newDeck()

	cards.print()

	hand, remainingDeck := deal(cards, 4)

	println("-----------------")
	hand.print()
	println("-----------------")

	remainingDeck.print()
}
