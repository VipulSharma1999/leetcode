package main

import (
	"math"
)

func increasingTriplet(nums []int) bool {
	n := len(nums)
	smallest := math.MaxInt64
	secondsmallest := math.MaxInt64
	for i := 0; i < n; i++ {
		if nums[i] <= smallest {
			smallest = nums[i]

		} else if nums[i] <= secondsmallest {
			secondsmallest = nums[i]
		} else {
			return true
		}
	}
	return false
}
