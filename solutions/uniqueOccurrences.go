package main

func uniqueOccurrences(arr []int) bool {
	n := len(arr)
	m := make(map[int]int, n)
	for _, val := range arr {
		m[val]++
	}
	revfreq := make(map[int]int)
	for val, i := range m {

		revfreq[i] = val
	}

	return len(m) == len(revfreq)
}
