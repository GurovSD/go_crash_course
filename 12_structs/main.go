package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init pesron using struct
	person1 := Person{
		firstName: "Samanta",
		lastName:  "Smith",
		city:      "Boston",
		gender:    "f",
		age:       25}

	person2 := Person{"Bob", "Johnson", "New York", "m", 30}

	// fmt.Println(person1)
	// person1.age++
	// fmt.Println(person1)

	// fmt.Println(person1.greet())
	// person1.hasBirthday()
	// fmt.Println(person1.greet())

	person1.getMarried("Williams")
	person2.getMarried("Smith")

	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
}
