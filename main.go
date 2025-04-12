package main

import "fmt"

func main() {
	cards := newDeck()

	cards.saveToFile("my_cards1")

	cards2 := newDeckFromFile("my_cards1")
	cards2.print()

	fmt.Println("====================")
	cards.shuffle()
	cards.print()

	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// fmt.Println("====================")
	// remainingDeck.print()
}
