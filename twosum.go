package main

import "fmt"

func twoSum(nums []int, target int) []int {
	// Create map to store the number and its index
	// key: number, value: index
	m := make(map[int]int)
	// Iterate over the array, O(n)
	for i, num := range nums {
		// Check if the target - num is in the map, return the index of the target - num and the current index
		if j, ok := m[target-num]; ok {
			return []int{j, i}
		}
		// If not, add the current number and its index to the map
		m[num] = i
	}
	return nil
}

func main() {
	nums := []int{2, 11, 11, 7}
	target := 9
	fmt.Println(twoSum(nums, target))
}
