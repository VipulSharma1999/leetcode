package main

func reverseVowels(s string) string {
	n := len(s)
	i := 0
	j := n - 1
	runes := make([]rune, len(s))
	for i, r := range s {
		runes[i] = r
	}
	for i <= j {
		if !isVowel(runes[i]) {
			i++
			continue
		} else if !isVowel(runes[j]) {
			j--
			continue
		}
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}

	return string(runes)
}

func isVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U'
}
