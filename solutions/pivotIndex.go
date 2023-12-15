package main

// 17 + 11 = 28 - 6 = 22
func pivotIndex(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	left := 0
	for i, v := range nums {
		//fmt.Println(left, v, sum-left-v)
		if left == sum-left-v {
			return i
		}
		left += v
	}
	return -1
}
