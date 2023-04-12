package main

import "fmt"

func maxProfit(prices []int) int {
	var n = len(prices)
	if n <= 1 {
		return 0
	}
	var dp = make([]int, n)
	dp[0] = 0
}

func main() {
	fmt.Println(maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
}
