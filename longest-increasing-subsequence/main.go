package main

import "fmt"

//时间复杂度O(n^2)
//空间复杂度O(n)
func lengthOfLIS(nums []int) int {
	var n = len(nums)
	if n == 0 {
		return 0
	}
	var dp = make([]int, n)
	dp[0] = 1
	var maxAns = 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxAns = max(maxAns, dp[i])
	}
	return maxAns
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
