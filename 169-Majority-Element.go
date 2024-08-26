package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	// Variables for the candidate and its count.
	var candidate int
	var count int

	// Traverse the array to determine the majority element.
	for _, num := range nums {
		if count == 0 {
			// If the count is 0, set the current element as the candidate.
			candidate = num
		}
		// Increment or decrement the count based on whether the current element matches the candidate.
		if num == candidate {
			count++
		} else {
			count--
		}
	}

	// The candidate now holds the majority element.
	return candidate
}

func main() {
	// Test cases
	fmt.Println(majorityElement([]int{3, 2, 3}))             // Output: 3
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2})) // Output: 2
}
