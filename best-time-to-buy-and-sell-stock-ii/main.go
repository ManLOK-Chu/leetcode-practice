package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	var n = len(prices)
	if n == 0 {
		return 0
	}
	//每天交易结束后只可能存在手里有一支股票或者没有股票
	//dp[i][0]表示没有持有股票，dp[i][1]表示持有股票
	var dp = make([][2]int, n)
	dp[0] = [2]int{0, -prices[0]}
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
