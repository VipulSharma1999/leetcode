package main

import (
	"strconv"
)

func compress(chars []byte) int {
	count, idx := 1, 0
	for i := 1; i <= len(chars); i++ {
		if i < len(chars) && chars[i] == chars[i-1] {
			count++
		} else {
			chars[idx] = chars[i-1]
			idx++
			if count > 1 {
				for _, ch := range strconv.Itoa(count) {
					chars[idx] = byte(ch)
					idx++
				}
			}
			count = 1
		}
	}
	return idx
}
