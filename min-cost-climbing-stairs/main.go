package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	var length = len(cost)
	var dp = make([]int, length+1)
	dp[0] = 0 // 默认第一步都是不花费体力的
	dp[1] = 0
	for i := 2; i <= length; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[length]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}
