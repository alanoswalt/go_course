package main

import "fmt"

func main() {
	//fmt.Println("Hello World")

	//Deck is a new type created in deck.go
	//deck == []string
	//cards := newDeck()
	//cards2 := []string{newCard(), "3 OF hearts"}

	//cards.print_cards()

	//hand, remainingDeck := deal(cards, 5)
	//hand.print_cards()
	//remainingDeck.print_cards()

	cards := newDeck()

	fmt.Println(cards.deckToString())

	cards.saveToFile("mycards")

	cards.shuffledeck()

	cards.print_cards()
	//cards2 := newDeckFromFile("mycards")

	//cards2.print_cards()

}
