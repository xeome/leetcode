package main

func romanToInt(s string) int {
	// Create a map to store the roman character and its value
	// key: character, value: value
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	// Create a variable to store the result
	result := 0
	// Iterate over the string, O(n)
	for i := 0; i < len(s); i++ {
		// If the current character is not the last character
		// If the value of the current character is less than the next character, subtract the value of the current character from the result
		// Reason: Certain letters are used to represent subtractive values. For example, the numeral IX represents the value 9 (10 - 1), and the numeral XL represents the value 40 (50 - 10).
		if i != len(s)-1 && m[s[i]] < m[s[i+1]] {
			result -= m[s[i]]
		} else {
			// Otherwise, add the value of the current character to the result
			result += m[s[i]]
		}
	}
	return result
}
