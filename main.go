package main

import "fmt"

// Defining struct which will be embedded in other struct
type contactInfo struct {
	email string
	pinCode int
}

// Defining struct type for person
type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func main() {
	person1 := person{
		firstName: "Jim",
		lastName: "Anderson",
		contact: contactInfo{
			email: "jim.an@gmail.com",
			pinCode: 281001,
		},
	}

	fmt.Printf("%+v", person1)
}