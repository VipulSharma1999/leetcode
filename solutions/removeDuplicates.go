package main

func removeDuplicates(nums []int) int {
	res := 1
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != prev {
			nums[res] = prev
			res++
		}
		prev = nums[i]
	}
	return res
}
