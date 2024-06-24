package main

import (
	"fmt"
	"time"
)

// Memoize takes a function f and returns a memoized version of f.
// It uses a map to store previously computed results.
func Memoize(f func(int) int) func(int) int {
	cache := make(map[int]int)
	return func(n int) int {
		if result, ok := cache[n]; ok {
			return result
		}
		result := f(n)
		cache[n] = result
		return result
	}
}

// Example function to be memoized (a simple factorial function).
func factorial(n int) int {
	time.Sleep(100 * time.Millisecond) //Simulate some computation time
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	// Create a memoized version of the factorial function.
	memoizedFactorial := Memoize(factorial)

	numbers := []int{5, 10, 5, 20, 15, 10, 20}

	for _, n := range numbers {
		start := time.Now()
		result := memoizedFactorial(n)
		elapsed := time.Since(start)

		fmt.Printf("Factorial of %d: %d. Time taken: %.2f seconds\n", n, result, elapsed.Seconds())
	}
}
