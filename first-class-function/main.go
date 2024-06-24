package main

import (
	"fmt"
)

// A function type declaration
type operation func(int, int) int

// A function that takes another function as an argument
func applyOperation(a, b int, op operation) int {
	return op(a, b)
}

// A function that returns another function
func getAdder() operation {
	return func(a, b int) int {
		return a + b
	}
}

func main() {
	// Assign a function to a variable
	add := func(a, b int) int {
		return a + b
	}

	subtract := func(a, b int) int {
		return a - b
	}

	// Using the function variables
	fmt.Println("Add: ", add(5, 3))         
	fmt.Println("Subtract: ", subtract(5, 3)) 

	// Passing function as an argument
	result := applyOperation(10, 5, add)
	fmt.Println("Apply Add Operation: ", result) 

	result = applyOperation(10, 5, subtract)
	fmt.Println("Apply Subtract Operation: ", result) 

	// Getting a function as a return value
	adder := getAdder()
	fmt.Println("Adder Function: ", adder(7, 8)) 
}