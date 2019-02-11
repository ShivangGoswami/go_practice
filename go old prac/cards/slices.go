package main

func main() {
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()
	// cards.print()
	//fmt.Println(cards)
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))
	//cards := newDeck()
	// fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")
	// cards := newDeckFromFile("my_cards") //"my"-->Error: open my: The system cannot find the file specified. exit status 1
	// cards.print()
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

// func newCard() string {
// 	return "Five Of Diamonds"
// }
