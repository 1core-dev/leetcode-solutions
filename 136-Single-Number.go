package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	// Initialize the result variable to 0.
	// This will store the cumulative XOR of all elements.
	result := 0

	for _, num := range nums {
		// XOR the current number with the result.
		// Due to the properties of XOR, pairs will cancel out, leaving only the number that appears once.
		result ^= num
	}

	// After all XOR operations, 'result' will hold the unique element.
	return result
}

func main() {
	// Test cases
	fmt.Println(singleNumber([]int{2, 2, 1}))       // Output: 1
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2})) // Output: 4
	fmt.Println(singleNumber([]int{1}))             // Output: 1
}
