package main

func generate(numRows int) [][]int {
	res := make([][]int, 0)
	for i := 1; i <= numRows; i++ {
		temp := generateRows(i)
		res = append(res, temp)
	}
	return res
}

func generateRows(row int) []int {
	ans := make([]int, 0)
	temp := 1
	ans = append(ans, 1)
	for i := 1; i < row; i++ {
		temp = temp * (row - i)
		temp = temp / i
		ans = append(ans, temp)
	}

	return ans
}
