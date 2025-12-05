package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Size of Spades")

	for index, card := range cards {
		fmt.Println(index, card)
	}

	// fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
