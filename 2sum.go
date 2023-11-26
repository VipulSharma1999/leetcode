package main

import (
	"fmt"
	"sync"
)

func main() {
	target := 6
	a := []int{1, 3, 4, 2}
	x := twoSum(a, target)
	fmt.Println(x)

}

func twoSum(a []int, t int) []int {
	//x := map[int]int{}
	s := newSet()
	// for in, i := range a {
	// 	x[i] = in
	// }
	// fmt.Println(x)
	for in, i := range a {
		if check := t - i; s.has(check) {
			idx := s.get(check)
			return []int{idx, in}
		}
		s.add(in, i)
	}

	return []int{}
}

type set struct {
	item map[int]int
	mu   sync.Mutex
}

func newSet() *set {
	return &set{
		item: make(map[int]int),
	}
}

func (s *set) add(pos, val int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.item[val] = pos
}

func (s *set) has(i int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, x := s.item[i]
	return x
}

func (s *set) get(pos int) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	num, ok := s.item[pos]
	if !ok {
		return -1
	}

	return num
}
