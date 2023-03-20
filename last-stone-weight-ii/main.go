package main

import "fmt"

func lastStoneWeightII(stones []int) int {
	var length = len(stones)
	if length == 1 {
		return stones[0]
	}
	var sum int
	for _, v := range stones {
		sum += v
	}
	var target = sum / 2
	var dp = make([]int, target+1)
	for i := 0; i < length; i++ { // 遍历物品
		for j := target; j >= stones[i]; j-- { // 遍历背包
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}
	return sum - dp[target] - dp[target]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(lastStoneWeightII([]int{2, 7, 4, 1, 8, 1}))
	fmt.Println(lastStoneWeightII([]int{31, 26, 33, 21, 40}))
}
