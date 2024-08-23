package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	// Start from the end of the array
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			// If the current digit is less than 9, just increment it
			digits[i]++
			return digits
		}
		// Set the current digit to 0 and carry over 1
		digits[i] = 0
	}

	// If we exited the loop, it means all digits were 9 and have been set to 0
	// We need to insert 1 at the beginning of the array
	result := make([]int, n+1)
	result[0] = 1
	return result
}

func main() {
	// Example 1
	digits1 := []int{1, 2, 3}
	fmt.Printf("Output: %v\n", plusOne(digits1)) // Output: [1, 2, 4]

	// Example 2
	digits2 := []int{4, 3, 2, 1}
	fmt.Printf("Output: %v\n", plusOne(digits2)) // Output: [4, 3, 2, 2]

	// Example 3
	digits3 := []int{9}
	fmt.Printf("Output: %v\n", plusOne(digits3)) // Output: [1, 0]
}
