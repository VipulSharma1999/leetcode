package main

import (
	"bytes"
)

func mergeAlternately(word1 string, word2 string) string {
	var b bytes.Buffer
	n := len(word1)
	m := len(word2)
	i := 0
	j := 0
	flag := true
	for i < n && j < m {
		if flag {
			b.WriteByte(word1[i])
			i++
		} else {
			b.WriteByte(word2[j])
			j++
		}
		flag = !flag
	}
	for i < n {
		b.WriteByte(word1[i])
		i++
	}
	for j < m {
		b.WriteByte(word2[j])
		j++
	}
	return b.String()
}
