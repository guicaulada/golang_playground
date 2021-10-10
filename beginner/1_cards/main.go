package main

func main() {
	// var card string = "Ace of Spades"
	// same thing as above but with type inference
	// card := "Ace of Spades"
	// card := newCard()

	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()
	// fmt.Println(cards.toString())

	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
