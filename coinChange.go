package main

import "math"

func coinChange(amount int, coins []int) int {
	n := len(coins)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
	}

	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = 0
		}
	}

	for i := 1; i < amount+1; i++ {
		dp[0][i] = math.MaxInt - 1
	}

	for i := 0; i < n+1; i++ {
		dp[i][0] = 0
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			if coins[i-1] <= j {
				dp[i][j] = min(dp[i-1][j], 1+dp[i][j-coins[i-1]])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	if dp[n][amount] == math.MaxInt-1 {
		return -1
	}

	return dp[n][amount]
}
