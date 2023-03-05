package main

func addBinary(a string, b string) string {
	var result string
	carry := 0
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
		}
		if j >= 0 {
			sum += int(b[j] - '0')
		}
		carry = sum / 2
		result = string(rune(sum%2+'0')) + result
	}
	if carry > 0 {
		result = "1" + result
	}
	return result
}
