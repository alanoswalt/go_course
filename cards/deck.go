package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Create a new type of eck
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

func (d deck) print_cards() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// Two inputs and two outputs
// Both outputs are type deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) deckToString() string {

	//La funcion join de la libreria string toma un array de strings y los junta con el separado de coma
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.deckToString()), 066)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)

	if err != nil {
		// Option 1 - log error and call new deck
		// Option 2 - Log error and quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	//Bs is a []byte, chance to string with "string(bs)"
	//Separate that long string to a []string with split
	s := strings.Split(string(bs), ",")

	return deck(s)

}

func (d deck) shuffledeck() {
	for index, _ := range d {

		random_number := rand.Intn(len(d) - 1)

		d[index], d[random_number] = d[random_number], d[index]
	}

}
