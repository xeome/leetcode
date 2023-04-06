package valid

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if !isAlphaNumeric(s[i]) {
			i++
			continue
		}

		if !isAlphaNumeric(s[j]) {
			j--
			continue
		}

		if toLower(s[i]) != toLower(s[j]) {
			return false
		}

		i++
		j--
	}
	return true
}

func isAlphaNumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}
