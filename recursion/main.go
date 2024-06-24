package main

import (
	"fmt"
)

// permute generates all permutations of the input string
func permute(input string) []string {
	var result []string
	permuteHelper([]rune(input), 0, &result)
	return result
}

// permuteHelper is the recursive helper function that generates permutations
func permuteHelper(arr []rune, start int, result *[]string) {
	if start == len(arr)-1 {
		*result = append(*result, string(arr))
		return
	}
	for i := start; i < len(arr); i++ {
		arr[start], arr[i] = arr[i], arr[start] // swap
		permuteHelper(arr, start+1, result)
		arr[start], arr[i] = arr[i], arr[start] // backtrack
	}
}

func main() {
	input := "abc"
	permutations := permute(input)
	for i, perm := range permutations {
		fmt.Printf("Permutation %d: %s\n", i+1, perm)
	}
}
