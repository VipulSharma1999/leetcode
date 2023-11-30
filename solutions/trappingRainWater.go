package main

func trap(height []int) int {
	n := len(height)
	left := 0
	right := n - 1
	ans := 0
	leftbar := height[0]
	rightbar := height[n-1]
	for left <= right {
		if leftbar < rightbar {
			if height[left] > leftbar {
				leftbar = height[left]
			} else {
				ans += leftbar - height[left]
				left++
			}
		} else {
			if height[right] > rightbar {
				rightbar = height[right]
			} else {
				ans += rightbar - height[right]
				right--
			}
		}
	}

	return ans
}
