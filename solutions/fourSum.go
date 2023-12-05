package main

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	ans := [][]int{}
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l := j + 1
			r := n - 1
			for l < r {
				temp := nums[i] + nums[j] + nums[l] + nums[r]
				if temp == target {
					ans = append(ans, []int{nums[i], nums[j], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				} else if temp < target {
					l++
				} else {
					r--
				}
			}
		}
	}
	return ans
}
