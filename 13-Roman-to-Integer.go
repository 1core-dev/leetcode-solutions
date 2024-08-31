package main

import (
	"fmt"
)

// Function to convert Roman numeral to integer
func romanToInt(s string) int {
	// Map of Roman numerals to their integer values
	romanToValue := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	// Initialize the integer result
	result := 0

	// Loop through the string
	length := len(s)
	for i := 0; i < length; i++ {
		// Get the value of the current numeral
		currentValue := romanToValue[s[i]]

		// Check if there is a next numeral and get its value
		if i < length-1 {
			nextValue := romanToValue[s[i+1]]
			// If the current value is less than the next value, subtract it
			if currentValue < nextValue {
				result -= currentValue
			} else {
				result += currentValue
			}
		} else {
			// Last numeral, just add its value
			result += currentValue
		}
	}

	return result
}

func main() {
	// Test cases
	testCases := []string{"III", "LVIII", "MCMXCIV"}

	for _, testCase := range testCases {
		fmt.Printf("Roman numeral %s is %d\n", testCase, romanToInt(testCase))
	}
}
