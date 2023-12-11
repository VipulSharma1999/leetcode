package main

import (
	"strings"
)

func reverseWords(s string) string {
	t := strings.Fields(s)
	start := 0
	n := len(t) - 1
	for start <= n {
		t[start], t[n] = t[n], t[start]
		start++
		n--
	}

	return strings.Join(t, " ")
}
