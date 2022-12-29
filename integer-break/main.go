package main

import "fmt"

func integerBreak(n int) int {
	var dp = make([]int, n+1)
	dp[1] = 0
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max((i-j)*j, dp[i-j]*j))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(integerBreak(2))
	fmt.Println(integerBreak(10))
}
