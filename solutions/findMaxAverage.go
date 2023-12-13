package main

func findMaxAverage(nums []int, k int) float64 {
	n := len(nums)
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	res := float64(sum)
	l := 0
	r := k
	for r < n {
		sum -= nums[l]
		l++

		sum += nums[r]
		r++

		res = maxx(res, float64(sum))
	}

	return res / float64(k)
}

func maxx(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
