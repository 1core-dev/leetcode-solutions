package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// Assume first slice element is longest prefix
	prefix := strs[0]
	// Loop through rest of slice elements starting from second
	for _, str := range strs[1:] {
		for i := 0; i < len(prefix) && i < len(str); i++ {
			// When str and prefix chars mismatch at index i, truncate prefix value to last match
			// break current loop and pick next str value for comparison
			if str[i] != prefix[i] {
				prefix = prefix[:i]
				break
			}
		}
		// When prefix len == 0 - no common prefix exist in slice
		if len(prefix) == 0 {
			return ""
		}
		// When len(prefix) > len(current str) -> update prefix to current str
		if len(prefix) > len(str) {
			prefix = str
		}
	}

	return prefix
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	longestCommonPrefix(strs)
}
