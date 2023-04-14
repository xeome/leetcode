package main

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	// variable to store the start and end index of the longest palindrome
	var start, end int
	for i := 0; i < len(s); i++ {
		// expand around the center for odd palindrome length
		len1 := expandAroundCenter(s, i, i)
		// expand around the center for even palindrome length
		len2 := expandAroundCenter(s, i, i+1)
		// pick the longer length
		length := max(len1, len2)
		// if the length is longer than the current longest, update the start and end index
		if length > end-start {
			start = i - (length-1)/2
			end = i + length/2
		}
	}
	// return the longest palindrome
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	// Expand around the center
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	// When the loop ends, we've gone one character too far
	// Return the length of the palindrome
	return right - left - 1
}
