package main

func lengthOfLastWord(s string) int {
	count := 0
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		if count != 0 && s[i] == ' ' {
			break
		}
		if s[i] == ' ' {
			continue
		}
		count++
	}
	return count
}
