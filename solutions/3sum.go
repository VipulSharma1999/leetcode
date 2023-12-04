package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	var ans [][]int
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := n - 1
		for l < r {
			temp := nums[i] + nums[l] + nums[r]
			if temp == 0 {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if temp < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return ans
}
