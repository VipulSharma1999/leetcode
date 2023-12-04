package main

import (
	"sync"
)

func longestValidParentheses(s string) int {
	n := len(s)
	if n == 0 || n == 1 {
		return 0
	}

	inStack := 0
	ans := 0

	st := NewStack()
	for i := 1; i <= n; i++ {
		if s[i-1] == '(' {
			st.Push(i)
		} else {
			if st.Empty() {
				inStack = i
				continue
			} else {
				st.Pop()
			}
			if st.Empty() {
				ans = max(ans, int(i-inStack))
			} else {
				ans = max(ans, int(i-st.Top()))
			}
		}
	}

	return ans
}

type Stack struct {
	m    sync.Mutex
	item []int
}

func NewStack() *Stack {
	return &Stack{
		item: nil,
	}
}

func (st *Stack) Empty() bool {
	st.m.Lock()
	defer st.m.Unlock()

	return len(st.item) == 0
}

func (st *Stack) Push(n int) {
	st.m.Lock()
	defer st.m.Unlock()

	st.item = append(st.item, n)

}

func (st *Stack) Pop() int {
	st.m.Lock()
	defer st.m.Unlock()
	if len(st.item) == 0 {
		return 0
	}

	lastItem := st.item[len(st.item)-1]
	st.item = st.item[:len(st.item)-1]
	return lastItem
}

func (st *Stack) Top() int {
	st.m.Lock()
	defer st.m.Unlock()

	if len(st.item) == 0 {
		return 0
	}
	return st.item[len(st.item)-1]
}
