package main

func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	if len(candidates) == 0 {
		return ans
	}

	backtracking(&ans, make([]int, 0), candidates, 0, target)

	return ans
}

func backtracking(ans *[][]int, temp, candidates []int, i, target int) {
	if target < 0 || i > len(candidates) {
		return
	}

	if target == 0 {
		ctemp := make([]int, len(temp))
		copy(ctemp, temp)
		*ans = append(*ans, ctemp)
	}

	for j := i; j < len(candidates); j++ {
		temp = append(temp, candidates[j])
		backtracking(ans, temp, candidates, j, target-candidates[j])
		temp = temp[:len(temp)-1]
	}
}
