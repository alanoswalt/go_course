package main

import "fmt"

//Costume type
type contact_info struct {
	//Here are not separated by comas
	email   string
	zipcode string
}

//Costume type
type person struct {
	//Here are not separated by comas
	firstName string
	lastName  string
	contact   contact_info //Possible way to embed a struct inside another
}

func main() {
	//How to define an username with a conctat into
	user_name := person{
		firstName: "Jimmy",
		lastName:  "Newtron",
		contact: contact_info{
			email:   "helo@helo.com",
			zipcode: "12345", //Even if is the last element all elements need
		},
	}

	//Aqui estoy dando la direccion
	//user_name_pointer := &user_name

	//Remove pointer variable but still works
	//This works because Go can recieve a type person or a type *person
	user_name.update_name("Mario")
	user_name.print()
}

//Structs can also be Recivers for functions
func (p person) print() {
	fmt.Printf("%+v", p)
}

//The struct is not getting updated here because the struct is just only a copy inside the function is not the actual function
func (pointerToUser *person) update_name(newFirstName string) {
	(*pointerToUser).firstName = newFirstName
}
