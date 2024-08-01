package main

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := []int{nums[0]}
	for i := 1; i < n; i++ {
		index := binarySearch(dp, nums[i])
		if index == len(dp) {
			dp = append(dp, nums[i])
		} else {
			dp[index] = nums[i]
		}
	}
	return len(dp)
}

// to fetch the lower bound
func binarySearch(nums []int, i int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == i {
			return mid
		} else if nums[mid] < i {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return low
}
