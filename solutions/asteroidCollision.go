package main

func asteroidCollision(asteroids []int) []int {
	res := []int{}
	var stack []int
	i := 0
	for i < len(asteroids) {
		if asteroids[i] > 0 {
			stack = append(stack, asteroids[i])
			i++
		} else {
			if len(stack) == 0 {
				res = append(res, asteroids[i])
				i++
			} else {
				value := 0 - asteroids[i]
				if value > stack[len(stack)-1] {
					stack = stack[:len(stack)-1]
				} else if value < stack[len(stack)-1] {
					i++
				} else {
					stack = stack[:len(stack)-1]
					i++
				}
			}
		}
	}

	res = append(res, stack...)

	return res
}
