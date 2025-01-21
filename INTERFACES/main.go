package main

import "fmt"

//Remember
//Every value in a go program has a typE
//Every function has to specify the type of its arguments so the arguments that I recive in a function have a type, or the ones to be outputed have a type

//Every function we ever write has to be rewritten if there to have the correct type?

//Incorrect way

type englishBot struct {
}

type spanoshBot struct {
}

//Interface

// This interface is telling:
// If you are a type in this code, with a function called getgreeting and you return a string then you are also a type bot
type bot interface {
	getgreeting() string
	//More complex interface
	//You can have multiple types or argumenrts and returns
	// getgreeting(string, int) (string, error)

	//Also posible to define
	//getBotVersion
}

//You can have same name function if the type linked is different
func (eb englishBot) getgreeting() string {
	return "Hello"

}

func (sb spanoshBot) getgreeting() string {
	return "Hola"

}

//Correct way using a function,
//The type change, so all bots can use same function
func printGreeing(b bot) {
	fmt.Println(b.getgreeting())
}

//Incorrect way, two functions that do the same but expect different types
//func printGreeing(eb englishBot) {
//	fmt.Println(eb.getgreeting())
//}

//func printGreeing(sb spanoshBot) {
//	fmt.Println(sb.getgreeting())
//}

func main() {

	eb := englishBot{}
	sb := spanoshBot{}

	printGreeing(sb)
	printGreeing(eb)

}
