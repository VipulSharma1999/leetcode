package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	prefix := make([]int, n)
	suffix := make([]int, n)

	prefix[0] = 1
	suffix[n-1] = 1

	fmt.Println(len(prefix))
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}
	for i := n - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}
	for i := 0; i < n; i++ {
		nums[i] = prefix[i] * suffix[i]
	}
	return nums
}
