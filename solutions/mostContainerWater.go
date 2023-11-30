package main

func maxArea(height []int) int {
	start := 0
	end := len(height) - 1
	maxarea := 0
	for start < end {
		temp := minn(height[start], height[end])
		area := temp * (end - start)
		maxarea = maxx(maxarea, area)

		if minn(height[start], height[end]) == height[start] {
			start++
		} else {
			end--
		}

	}
	return maxarea
}

func minn(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxx(a, b int) int {
	if a > b {
		return a
	}
	return b
}
