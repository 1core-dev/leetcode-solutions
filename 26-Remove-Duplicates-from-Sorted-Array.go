package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Initialize the write index
	writeIndex := 0

	// Iterate through the array starting from the second element
	for readIndex := 1; readIndex < len(nums); readIndex++ {
		if nums[readIndex] != nums[writeIndex] {
			// Move to the next position for unique element
			writeIndex++
			// Update the element at the write index with the new unique element
			nums[writeIndex] = nums[readIndex]
		}
	}

	// Return the count of unique elements
	return writeIndex + 1
}

func main() {
	// Example 1
	nums1 := []int{1, 1, 2}
	k1 := removeDuplicates(nums1)
	fmt.Printf("Output: %d, nums = %v\n", k1, nums1[:k1])

	// Example 2
	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// nums2 := []int{0, 1, 2, 3, 4, 2, 2, 3, 3, 4}
	k2 := removeDuplicates(nums2)
	fmt.Printf("Output: %d, nums = %v\n", k2, nums2[:k2])
}
