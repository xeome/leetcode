package main

func removeElement(nums []int, val int) int {
	last := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[last] = nums[i]
			last++
		}
	}
	return last
}
