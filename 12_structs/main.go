package main

import (
	"fmt"
	"strconv"
)

// Person Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// Get married (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	}
	p.lastName = spouseLastName

}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}
	person2 := Person{"Bob", "Johnson", "New York", "m", 30}

	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1)

	person1.hasBirthday()
	person1.getMarried("Williams")
	fmt.Println(person1.greet())

	person2.getMarried("Thompson")
	fmt.Println(person2.greet())
}
