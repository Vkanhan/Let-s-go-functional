package main

import "fmt"

// Function sequenceGenerator returns a closure that generates a sequence of numbers.
func sequenceGenerator(start, stop int) func() int {
	next := start - stop
	return func() int {
		next += stop
		return next
	}
}

func main() {
	// Creating two instances of sequence generators with different configurations
	generator1 := sequenceGenerator(0, 2)  // Start from 0, increment by 2
	generator2 := sequenceGenerator(10, 3) // Start from 10, increment by 3

	// Generate and print sequences from generator1
	fmt.Println("Generating sequence from generator1:")
	for i := 0; i < 5; i++ {
		fmt.Println(generator1())
	}

	fmt.Println()

	// Generate and print sequences from generator2
	fmt.Println("Generating sequence from generator2:")
	for i := 0; i < 5; i++ {
		fmt.Println(generator2())
	}
}
