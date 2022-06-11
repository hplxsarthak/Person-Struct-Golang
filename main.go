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
	contactInfo
}

func main() {
	person1 := person{
		firstName: "Jim",
		lastName: "Anderson",
		contactInfo: contactInfo{
			email: "jim.an@gmail.com",
			pinCode: 281001,
		},
	}

	// Dealing with pointers
	// Turn "address" into "value" with "*address"
	// Turn "value" into "address" with "&value"
	person1Pointer := &person1 // Taking the address of the value of the person1 variable using "&"
	person1Pointer.updateName("Jimmy") // Passing the value to be change using pointers
	person1.print()
}

// Updating using pointers
func (pointerToPerson *person) updateName (newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// Receiver function in struct
func (p person) print() {
	fmt.Printf("%+v", p)
}