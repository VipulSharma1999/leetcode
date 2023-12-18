package main

import (
	"sort"
)

func closeStrings(word1 string, word2 string) bool {
	n := len(word1)
	m := len(word2)
	if n != m {
		return false
	}
	f1 := make([]int, 26)
	f2 := make([]int, 26)

	for i := 0; i < n; i++ {
		f1[word1[i]-'a']++
	}

	for i := 0; i < m; i++ {
		f2[word2[i]-'a']++
	}

	for i := 0; i < 26; i++ {
		if (f1[i] == 0 && f2[i] > 0) ||
			(f1[i] > 0 && f2[i] == 0) {
			return false
		}
	}

	sort.Ints(f1)
	sort.Ints(f2)

	for i := 0; i < 26; i++ {
		if f1[i] != f2[i] {
			return false
		}
	}

	return true
}
