package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	var n = len(prices)
	if n <= 1 {
		return 0
	}
	var dp = make([]int, n+1)
	dp[0] = 0
	dp[1] = 0
	minPrice := prices[0]
	for i := 2; i <= n; i++ {
		minPrice = min(minPrice, prices[i-1])
		dp[i] = max(dp[i-1], prices[i-1]-minPrice)
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//func maxProfit1(prices []int) int {
//	buy := prices[0]
//	sell := prices[0]
//	max := math.MinInt
//	for _, v := range prices {
//		if buy > v {
//			buy = v
//			sell = buy
//		} else if buy < v {
//			sell = v
//			gap := sell - buy
//			if gap > max {
//				max = gap
//			}
//		}
//	}
//	if max > 0 {
//		return max
//	}
//	return 0
//}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit([]int{0}))
	fmt.Println(maxProfit([]int{3, 2, 6, 5, 0, 3}))
}
