package main

// 121
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	res := 0
	temp := x
	for x > 0 {
		rem := x % 10
		res = (res * 10) + rem
		x /= 10
	}

	return res == temp
}
