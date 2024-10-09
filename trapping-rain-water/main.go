package main

import (
	"fmt"
)

/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
*/
//动态规划
func trap(height []int) int {
	sum := 0
	//max_left [i] 代表第 i 列左边最高的墙的高度，max_right[i] 代表第 i 列右边最高的墙的高度
	//第 i 列左（右）边最高的墙，是不包括自身的
	var maxLeft = make([]int, len(height))
	var maxRight = make([]int, len(height))
	for i := 1; i < len(height)-1; i++ {
		//它前边的墙的左边的最高高度和它前边的墙的高度选一个较大的，就是当前列左边最高的墙了。
		maxLeft[i] = max(maxLeft[i-1], height[i-1])
	}
	for i := len(height) - 2; i >= 0; i-- {
		//它后边的墙的右边的最高高度和它后边的墙的高度选一个较大的，就是当前列右边最高的墙了。
		maxRight[i] = max(maxRight[i+1], height[i+1])
	}
	for i := 1; i < len(height)-1; i++ {
		minNum := min(maxLeft[i], maxRight[i])
		if minNum > height[i] {
			sum += minNum - height[i]
		}
	}
	return sum
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

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
}
