package main

func lengthOfLongestSubstring(s string) int {
	// Create a map to store the character and its index
	// key: character, value: index
	m := make(map[byte]int)
	// Create a variable to store the max length
	max := 0
	// Create a variable to store the current length
	cur := 0
	// Create a variable to store the current start index
	start := 0
	// Iterate over the string, O(n)
	for i := 0; i < len(s); i++ {
		// If the character is in the map and the index is greater than the start index, update the start index
		if index, ok := m[s[i]]; ok && index >= start {
			start = index + 1
		}
		// Update the current length
		cur = i - start + 1
		// Update the max length
		if cur > max {
			max = cur
		}
		// Add the character and its index to the map
		m[s[i]] = i
	}
	return max
}
