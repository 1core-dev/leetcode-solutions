package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {
	// Use strings.Index to find the first occurrence of needle in haystack
	return strings.Index(haystack, needle)
}

func main() {
	// Example 1
	haystack1 := "sadbutsad"
	needle1 := "sad"
	fmt.Printf("The index of %q in %q is %d\n", needle1, haystack1, strStr(haystack1, needle1)) // Output: 0

	// Example 2
	haystack2 := "leetcode"
	needle2 := "leeto"
	fmt.Printf("The index of %q in %q is %d\n", needle2, haystack2, strStr(haystack2, needle2)) // Output: -1
}
