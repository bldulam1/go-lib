package main

import (
	"fmt"
	"strconv"
)

// Person struct
type Person struct {
	firstName string
	lastName  string
	email     string
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + strconv.Itoa(14)
}

func main() {
	person := Person{"Brendon", "Dulam", "brendon.dulam@veoneer.com"}
	fmt.Println(person.greet())
}
