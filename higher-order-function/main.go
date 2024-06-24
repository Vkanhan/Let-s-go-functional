package main 

import(
	"fmt"
)

//squareFunction takes function f as an argument and returns another function that squares the input
func squareFunction(f func(int) int) func(int) int {
    return func(x int) int {
        squared := x * x
        return f(squared)
    }
}

//double func doubles the input
func double(n int) int {
	return n * 2
}

func main() {
	//create a new function by applying squareFunction to the double func
	squareDoubled := squareFunction(double)

	inputs := []int{2, 4, 6, 8}

	for _, input := range inputs {
		result := squareDoubled(input)
		fmt.Printf("input: %d, squared and doubled: %d\n", input, result)

	}

}