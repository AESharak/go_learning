package main

func main() {

	cards := deck{"Ace of Diamongs", newCard()}

	cards.print()


	cards = append(cards, "Ace of Spades")


	cards.print()
}


func newCard() string {
	return "Five of Diamonds"
}
