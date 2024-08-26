package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	// Convert the string to lowercase to handle case insensitivity
	s = strings.ToLower(s)

	// Initialize two indexes: one at the start and one at the end of the string
	leftIndex := 0
	rightIndex := len(s) - 1

	// Loop until the indexes meet in the middle
	for leftIndex < rightIndex {
		// Move the left index forward if the current character is not alphanumeric
		if !isAlphanumeric(s[leftIndex]) {
			leftIndex++
			continue
		}

		// Move the right index backward if the current character is not alphanumeric
		if !isAlphanumeric(s[rightIndex]) {
			rightIndex--
			continue
		}

		// Compare the characters at both indexes
		if s[leftIndex] != s[rightIndex] {
			return false // If they don't match, it's not a palindrome
		}

		// Move both indexes towards the center
		leftIndex++
		rightIndex--
	}

	return true // If all characters match, it's a palindrome
}

// isAlphanumeric checks if a character is a letter or a digit;
// when you index a string like s[i], the result is a byte
func isAlphanumeric(c byte) bool {
	return unicode.IsLetter(rune(c)) || unicode.IsDigit(rune(c))
}

func main() {
	// Test cases
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // Output: true
	fmt.Println(isPalindrome("race a car"))                     // Output: false
	fmt.Println(isPalindrome(" "))                              // Output: true
}
