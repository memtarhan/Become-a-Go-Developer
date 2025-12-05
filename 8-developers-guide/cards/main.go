package main

import "fmt"

func main() {
	card := "Ace of Spades"
	fmt.Println(card) 

	card2 := newCard()
	fmt.Println(card2)
}

func newCard() string {
	return "Five of Diamonds"
}