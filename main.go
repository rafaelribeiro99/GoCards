package main

func main() {
	// cards := newDeck()

	// cards.print()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	cards := newDeck()
	//fmt.Println(cards.toString())
	cards.saveToFile("my_cards")

}
