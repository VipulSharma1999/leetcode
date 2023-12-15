package main

func largestAltitude(gain []int) int {
	n := len(gain)
	res := make([]int, n+1)
	res[0] = 0
	i := 1
	for _, v := range gain {
		res[i] += v + res[i-1]
		i++
	}
	ans := 0
	for i := 0; i < len(res); i++ {
		if res[i] > 0 && res[i] > ans {
			ans = res[i]
		}
	}

	return ans
}
