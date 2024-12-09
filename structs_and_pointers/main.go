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

	//Order of definition in the struct
	//username := person{"Alex", "Garcia"}

	//Posible to change order when defininf
	//username2 := person{firstName: "Alan", lastName: "Sanchez"}

	//With no initialization go assinge a null or 0 value to the fields
	//var username3 person

	//fmt.Println(username)
	//fmt.Println(username2)

	//with this it prints also the fields and its values
	//fmt.Printf("%+v", username)

	//Update values
	//username.firstName = "Pepe"

	//How to define an username with a conctat into

	user_name := person{
		firstName: "Jimmy",
		lastName:  "Newtron",
		contact: contact_info{
			email:   "helo@helo.com",
			zipcode: "12345", //Even if is the last element all elements need
		},
	}

	//fmt.Println(user_name)
	//fmt.Printf("%+v", user_name)

	//Aqui estoy dando la direccion
	user_name_pointer := &user_name
	user_name_pointer.update_name("Mario")

	//user_name.update_name("Mario")
	user_name.print()
}

//Structs can also be Recivers for functions
func (p person) print() {
	fmt.Printf("%+v", p)
}

//The struct is not getting updated here because the struct is just only a copy inside the function is not the actual function
//func (p person) update_name(newFirstName string) {
//	p.firstName = newFirstName
//}

//The struct is not getting updated here because the struct is just only a copy inside the function is not the actual function
func (pointerToUser *person) update_name(newFirstName string) {
	(*pointerToUser).firstName = newFirstName
}
