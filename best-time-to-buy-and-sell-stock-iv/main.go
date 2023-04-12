package main

import "fmt"

func maxProfit(k int, prices []int) int {
	var n = len(prices)
	if n <= 1 || k == 0 {
		return 0
	}
	var local = make([][]int, n)
	var global = make([][]int, n)
	local[0] = make([]int, k+1)
	global[0] = make([]int, k+1)
	for i := 1; i < n; i++ {
		local[i] = make([]int, k+1)
		global[i] = make([]int, k+1)
		curProfit := prices[i] - prices[i-1]
		for j := k; j >= 1; j-- {
			local[i][j] = max(global[i-1][j-1]+curProfit, local[i-1][j]+curProfit)
			global[i][j] = max(local[i][j], global[i-1][j])
		}
	}
	return global[n-1][k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProfit(2, []int{2, 4, 1}))
}
