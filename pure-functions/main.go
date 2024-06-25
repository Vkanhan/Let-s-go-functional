package main

import (
	"fmt"
)

// Pure function: does not modify any external state, and given the same inputs, it always produces the same outputs.
func add(a, b int) int {
	return a + b
}

func main() {
	result := add(2, 4)
	fmt.Println(result)
}
