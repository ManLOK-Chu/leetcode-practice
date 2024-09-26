package main

import (
	"fmt"
)

/*
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组是数组中的一个连续部分。
*/
func maxSubArray(nums []int) int {
	//动态规划
	length := len(nums)
	var result int
	var dp = make([]int, length)
	dp[0] = nums[0]
	result = dp[0]
	for i := 1; i < length; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		result = max(dp[i], result)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//贪心算法
//func maxSubArray(nums []int) int {
//	var result = math.MinInt32
//	var count int
//	for i := 0; i < len(nums); i++ {
//		count += nums[i]
//		if count > result { // 取区间累计的最大值（相当于不断确定最大子序终止位置）
//			result = count
//		}
//		if count <= 0 {
//			count = 0 // 相当于重置最大子序起始位置，因为遇到负数一定是拉低总和
//		}
//	}
//	return result
//}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{1}))
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
	fmt.Println(maxSubArray([]int{-1, -2}))
}
