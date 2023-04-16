package main

import "fmt"

//时间复杂度O(n) 空间复杂度O(n)
// 成环情况1：考虑不包含首尾元素
// 成环情况2：考虑包含首元素，不包含尾元素
// 成环情况3: 考虑包含尾元素，不包含首元素
// 情况2 和 情况3 都包含了情况一了，所以只考虑情况二和情况三就可以了。
func rob(nums []int) int {
	var n = len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	var robRange func(int, int) int
	robRange = func(start, end int) int {
		if end == start {
			return nums[start]
		}
		var dp = make([]int, n)
		dp[start] = nums[start]
		dp[start+1] = max(nums[start], nums[start+1])
		//从前到后遍历
		for i := start + 2; i <= end; i++ {
			dp[i] = max(dp[i-2]+nums[i], dp[i-1])
		}
		return dp[end]
	}
	return max(
		robRange(0, n-2), //情况2：考虑包含首元素，不包含尾元素
		robRange(1, n-1), //情况3：考虑包含尾元素，不包含首元素
	)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(rob([]int{2, 3, 2}))
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{1, 2, 1}))
}
