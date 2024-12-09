package main

import (
	"fmt"
)

// lets create a new type of map with keys and values of type string
type string_map map[string]string

func main() {

	//Map declaaration, first element inside brackets are the keys, the next are the values
	//In this example we are delcareing a map with keys and values as strings
	//REMEMEBER ADD A COMA EVEN ON THE LAST VALUE LIKE STRCTS
	//colors := map[string]string{
	//	"red":  "ff0000",
	//	"blue": "00000",
	//}

	// Way to create Empty maps, other way to declare it
	//This declares a map
	//var colors map[string]string
	//this declares an initialize a map
	colors2 := make(map[string]string)

	//How to add new values to a map
	colors2["Red"] = "FFFFF"
	colors2["blue"] = "FFFFF1"
	colors2["green"] = "00000"

	//How to remove a key and value froma dictionarie
	//delete(colors2, "green")

	fmt.Println(colors2["Red"])
	fmt.Println(colors2)

	iterate_map(colors2)

}

func iterate_map(c string_map) {

	for key, value := range c {

		fmt.Println(key)
		fmt.Println(value)

	}

}
