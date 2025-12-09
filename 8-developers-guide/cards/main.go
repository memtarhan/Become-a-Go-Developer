package main

func main() {
	// cards := newDeck()

	// hand, remaining := deal(cards, 5)

	// hand.print()
	// remaining.print()
	cards := newDeck()
	// fmt.Println(cards.toString())
	cards.saveToFile("cards")

	newCards := newDeckFromFile("cards")
	newCards.print()
}
