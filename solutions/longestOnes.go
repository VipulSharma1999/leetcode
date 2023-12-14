package main

func longestOnes(nums []int, k int) int {
	n := len(nums)
	count := 0
	l, r := 0, 0
	res := 0
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			count++
		}
		for count > k {
			if nums[l] == 0 {
				count--
			}
			l++
		}
		r++
		if res < r-l {
			res = (r - l)
		}
	}

	return res
}
