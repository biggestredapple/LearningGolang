package main

func main() {

	var card = getNewCard()
	card.print()

	cards := getNewDeck()
	cards = append(cards, "Six of Spades")
	cards.print("Current Deck")

	hand, remainingCards := deal(cards, 5)
	hand.print("Hand Deck")
	remainingCards.print("Remaining Deck")
}
