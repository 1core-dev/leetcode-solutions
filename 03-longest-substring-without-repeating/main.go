package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	// Window [left ... right] always contains unique characters
	lastPosition := make(map[rune]int)

	left := 0
	maxLength := 0 // maximum length seen so far

	for right, ch := range s {
		// If duplicate is inside current window, move left boundary
		if prev, exists := lastPosition[ch]; exists && prev >= left {
			left = prev + 1
		}
		lastPosition[ch] = right

		// Measure current valid window
		currentLength := right - left + 1

		// Update the maximum length found so far
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcbdefbac"))
}
