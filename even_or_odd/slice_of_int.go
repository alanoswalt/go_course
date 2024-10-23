package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Create a new type called deck deck is equal to an array of strings
type slice_of_int []int

func newSlice() slice_of_int {

	new_Slice_of_int := slice_of_int{}

	counter := 0
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for counter < 10 {
		random_number := r.Intn(100)
		new_Slice_of_int = append(new_Slice_of_int, random_number)
		counter++
	}

	return new_Slice_of_int

}

func (s slice_of_int) ever_or_odd() {

	for _, num := range s {
		if num%2 == 0 {
			fmt.Println(num, " is even")
		} else {
			fmt.Println(num, " is odd")
		}
	}

}
