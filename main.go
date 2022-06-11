package main

import "fmt"

// Defining struct type for person
type person struct {
	firstName string
	lastName string
}

func main() {
	// Third way to define struct
	var person1 person

	person1.firstName = "Alex"
	person1.lastName = "Anderson"

	fmt.Println(person1)
	fmt.Printf("%+v", person1) // Using this feature of print we can print the struct in the form of object
}