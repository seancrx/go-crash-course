package main

import "fmt"

func main() {
	// Using var
	// var name string = "Sean"
	var age int = 25
	const isCool = true

	// Shorthand
	name, size := "Sean", 1.3
	// size := 1.3

	fmt.Println(name, age, isCool)
	fmt.Printf("%T\n", size)
}
