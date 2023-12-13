package main

import (
	"sort"
)

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	l, r := 0, n-1
	count := 0
	for l < r {
		if nums[l]+nums[r] > k {
			r--
		} else if nums[l]+nums[r] < k {
			l++
		} else {
			count++
			l++
			r--
		}
	}
	return count
}
