package main

import "fmt"

// Defining struct type for person
type person struct {
	firstName string
	lastName string
}

func main() {
	person1 := person{"Sarthak" , "Agrawal"} // In this type we cannot change order
	person2 := person{firstName: "Alex", lastName: "Anderson"} // In this we can change order

	fmt.Println(person1)
	fmt.Println(person2.firstName)
}