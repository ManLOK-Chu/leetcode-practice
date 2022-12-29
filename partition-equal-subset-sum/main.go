package main

import "fmt"

func canPartition(nums []int) bool {
	var length = len(nums)
	if length == 1 {
		return false
	}
	var sum int
	for _, v := range nums {
		sum += v
	}
	if sum%2 == 1 {
		return false
	}
	var target = sum / 2
	var dp = make([]int, target+1)
	for i := 0; i < length; i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	if dp[target] == target {
		return true
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
	fmt.Println(canPartition([]int{1, 2, 3, 5}))
}
