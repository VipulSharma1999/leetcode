package main

func maxVowels(s string, k int) int {
	n := len(s)
	count := 0
	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			count++
		}
	}
	res := count
	for i := k; i < n; i++ {
		if isVowel(s[i-k]) {
			count--

		}
		if isVowel(s[i]) {
			count++
		}

		if res < count {
			res = count
		}
	}
	return res
}

func isVowel(ch byte) bool {
	if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
		return true
	}
	return false
}
