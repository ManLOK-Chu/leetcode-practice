package main

import "fmt"

func numTrees(n int) int {
	if n < 3 {
		return n
	}
	var dp = make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(numTrees(3))
	fmt.Println(numTrees(1))
	fmt.Println(numTrees(5))
}
