package main

import (
	"fmt"
)

func titleToNumber(columnTitle string) int {
	columnNumber := 0
	for _, char := range columnTitle {
		// Convert character to integer (A=1, B=2, ..., Z=26)
		value := int(char - 'A' + 1)
		// Accumulate the result considering the base-26 system
		columnNumber = columnNumber*26 + value
	}
	return columnNumber
}

func main() {
	examples := []string{"A", "AB", "ZY", "FXSHRXW"} // Test cases

	for _, example := range examples {
		result := titleToNumber(example)
		fmt.Printf("The column number of '%s' is %d\n", example, result)
	}
}
