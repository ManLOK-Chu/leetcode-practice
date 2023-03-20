package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	var result = math.MinInt32
	var count int
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if count > result { // 取区间累计的最大值（相当于不断确定最大子序终止位置）
			result = count
		}
		if count <= 0 {
			count = 0 // 相当于重置最大子序起始位置，因为遇到负数一定是拉低总和
		}
	}
	return result
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{1}))
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
}
