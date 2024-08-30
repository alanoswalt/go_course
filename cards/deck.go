package main

import "fmt"

//Create a new type of eck
type deck []string

func newDeck() deck {
	cards2 := deck{}

	cardSuits := []string{"Spades", "Hearts", "Dimonds", "Clubs"}

	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, value := range cardValues {

		for _, suit := range cardSuits {
			cards2 = append(cards2, value+" of "+suit)
		}
	}

	return cards2
}

//
func (d deck) print_cards() {
	for _, card := range d {
		fmt.Println(card)
	}
}

//Two inputs and two outputs
//Both outputs are type deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
