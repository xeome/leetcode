package main

func singleNumber(nums []int) int {
	var result int
	for _, num := range nums {
		result ^= num // XOR, because a ^ a = 0, 0 ^ a = a
	}
	return result
}
