package main

import (
	"fmt"
	"sync"
)

type Stack struct {
	item []rune
	mu   sync.Mutex
}

func NewStack() *Stack {
	return &Stack{
		item: nil,
	}
}

func main() {
	s := "()[]{}"
	ans := validParenthesis(s)
	fmt.Println(ans)
}

func validParenthesis(s string) bool {
	st := NewStack()
	for _, ch := range s {
		if ch == ']' || ch == '}' || ch == ')' {
			if len(st.item) == 0 {
				return false
			}
			if st.Top() != '[' && ch == ']' {
				return false
			}
			if st.Top() != '{' && ch == '}' {
				return false
			}
			if st.Top() != '(' && ch == ')' {
				return false
			}
			st.Pop()
		} else {
			st.Push(ch)
		}
	}

	return len(st.item) == 0
}

func (s *Stack) Top() rune {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.item) == 0 {
		return 0
	}
	return s.item[len(s.item)-1]
}

func (s *Stack) Push(it rune) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.item = append(s.item, it)
}

func (s *Stack) Pop() rune {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.item) == 0 {
		return 0
	}

	lastItem := s.item[len(s.item)-1]
	s.item = s.item[:len(s.item)-1]
	return lastItem
}
