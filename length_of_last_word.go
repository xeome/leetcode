package main

func lengthOfLastWord(s string) int {
	var count int
	i := len(s) - 1
	for i >= 0 && s[i] == ' ' {
		i--
	}
	for i >= 0 && s[i] != ' ' {
		count++
		i--
	}
	return count
}
