package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type called deck deck is equal to an array of strings
type deck []string

// Expect a deck as output
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

// This function can be called to the deck type lide deck.print:
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

	//Return time as a int64 value with that you create a source
	source := rand.NewSource(time.Now().UnixNano())

	//With a source you can create a random number
	//Rand new ewturns an type rand so r can use Intn
	r := rand.New(source)

	for index, _ := range d {

		random_number := r.Intn(len(d) - 1)

		d[index], d[random_number] = d[random_number], d[index]
	}

}
