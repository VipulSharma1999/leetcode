package main

func moveZeroes(nums []int) {
	n := len(nums)
	idx, count := 0, 0
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			count++
		} else {
			nums[idx] = nums[i]
			idx++
		}
	}
	for i := idx; i < n; i++ {
		nums[i] = 0
	}
}
