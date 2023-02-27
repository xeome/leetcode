package main

func isPalindrome(x int) bool {
	// If the number is negative, return false
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	// Create a variable to store the reverse of the number
	reverse := 0
	// Create a variable to store the original number
	original := x
	// Iterate over the number, O(n)
	for x != 0 {
		// Update the reverse
		reverse = reverse*10 + x%10
		// Update the number
		x /= 10
	}
	// Return true if the reverse is equal to the original number
	return reverse == original
}
