package main

import (
	"fmt"
	"strconv"
)

func decodeString(s string) string {
	stack := []rune{}
	for _, char := range s {
		if char != ']' {
			stack = append(stack, char)
		} else {
			substr := ""
			for stack[len(stack)-1] != '[' {
				fmt.Printf("stack len-1 element=%v\n", string(stack[len(stack)-1]))
				substr = string(stack[len(stack)-1]) + substr
				stack = stack[:len(stack)-1]
			}
			//pop
			stack = stack[:len(stack)-1]

			numStr := ""
			for len(stack) > 0 && stack[len(stack)-1] >= '0' && stack[len(stack)-1] <= '9' {
				fmt.Printf("stack=%v\n", string(stack))
				numStr = string(stack[len(stack)-1]) + numStr
				stack = stack[:len(stack)-1] // pop
			}

			num, _ := strconv.Atoi(numStr)
			for num > 0 {
				stack = append(stack, []rune(substr)...)
				num--
			}
		}
	}

	return string(stack)
}
