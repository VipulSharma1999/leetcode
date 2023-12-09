package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	n := len(candies)
	maxi := 0
	res := make([]bool, n)
	for _, v := range candies {
		maxi = maxx(maxi, v)
	}
	for i := 0; i < n; i++ {
		if extraCandies+candies[i] >= maxi {
			res[i] = true
		} else {
			res[i] = false
		}
	}
	return res
}

func maxx(a, b int) int {
	if a > b {
		return a
	}
	return b
}
