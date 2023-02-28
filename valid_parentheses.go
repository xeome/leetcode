package main

func isValid(s string) bool {
	// Stack for storing the left parentheses
	var stack []rune
	// Iterate through the string, O(n)
	for _, r := range s {
		// If the current character is a left parentheses, push it to the stack
		if r == '(' || r == '[' || r == '{' {
			stack = append(stack, r)
		} else {
			// If stack is empty, return false
			if len(stack) == 0 {
				return false
			}
			// If the current character is a right parentheses and does not match the top of the stack (appropriate left parentheses), return false
			top := stack[len(stack)-1]
			if (r == ')' && top != '(') || (r == ']' && top != '[') || (r == '}' && top != '{') {
				return false
			}
			// Pop the top of the stack, as the current character matches the top of the stack
			stack = stack[:len(stack)-1]
		}
	}
	// If the stack is empty, return true, otherwise return false, as there are still unmatched left parentheses meaning the string is invalid
	return len(stack) == 0
}
