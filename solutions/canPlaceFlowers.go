package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	res := 0
	first1 := -1
	last1 := -1
	l := len(flowerbed)
	for i := 0; i < l; i++ {
		if flowerbed[i] == 0 {
			continue
		} else if first1 == -1 {
			first1 = i
			last1 = i
		} else {
			last1 = i
		}
	}

	// 0 0 0 0 0 0 0 0 0
	if first1 == -1 {
		return n <= (l+1)/2
	}

	res = first1 / 2
	res += ((l - last1 - 1) / 2)

	count := 0
	for i := first1 + 1; i <= last1-1; i++ {
		if flowerbed[i] == 0 {
			count++
		} else {
			res += (count - 1) / 2
			count = 0
		}
	}

	res += (count - 1) / 2
	return n <= res
}
