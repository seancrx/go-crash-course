package main

import "fmt"

// functionName(parameterName parameterType) returnType
func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Sean"))
	fmt.Println(getSum(3, 4))
}
