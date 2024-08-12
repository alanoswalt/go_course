package main

import (
	"fmt"
)

func main() {

	//var: informs go to declare a variable
	//card: name of the variable
	//string: type of variable
	//Se puede usar dentro y fuera de una funcion
	var card string = "1"

	//Can declare and later assing
	var card3 string
	card3 = "2"

	//Using := helps omit the var word and assumes the type of the variable
	//just use to declare new variables
	//This type of delcaration can only be use inside functions
	card2 := "3"

	//Assinge with function
	card4 := newCard()

	fmt.Println(card)
	fmt.Println(card2)
	fmt.Println(card3)
	fmt.Println(card4)

	practiceArray()
}

// Func: inform go this is a function
// newCard: name of function
// inside () the expected input
// string: expected return
func newCard() string {

	return "4"
}

func practiceArray() {

	//Array is fix in size
	//Slice is an array that can grow or shrink
	//Both need a  fixed type

	var cards = [2]string{newCard(), "2 OF hearts"}
	cards2 := []string{newCard(), "3 OF hearts"}

	//Add values
	//This doesn't add a new value to existing arrat
	//It creates a new array and assigens to same value
	cards2 = append(cards2, "Four of spades")

	fmt.Println(cards)
	fmt.Println(cards2)

	// i: current index
	// card: element in the
	//use range to iterare all records inside an array
	for i, card := range cards2 {
		fmt.Println(i, card)
	}

	//To not use i variable
	for _, card := range cards2 {
		fmt.Println(card)
	}

	//this will take only the index
	for i := range cards2 {
		fmt.Println(i)
	}

	//for loop c++ style
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	//for loop c++ style
	var i int
	for i = 0; i < 3; i++ {
		fmt.Println(i)
	}
}
