package main

func longestSubarray(nums []int) int {
	l, r := 0, 0
	zero := 0
	n := len(nums)
	for r = 0; r < n; r++ {
		if nums[r] == 0 {
			zero++
		}

		if zero > 1 {
			if nums[l] == 0 {
				zero--
			}
			l++
		}
	}

	return r - l - 1
}
