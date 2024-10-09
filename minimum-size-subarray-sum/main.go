package main

import "fmt"

/*
给定一个含有 n 个正整数的数组和一个正整数 target 。
找出该数组中满足其总和大于等于 target 的长度最小的子数组[numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。
如果不存在符合条件的子数组，返回 0 。
*/
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	ans, sum, left := n+1, 0, 0
	for right, x := range nums { // 枚举子数组右端点
		sum += x
		for sum-nums[left] >= target { // 尽量缩小子数组长度
			sum -= nums[left]
			left++ // 左端点右移
		}
		if sum >= target {
			ans = min(ans, right-left+1)
		}
	}
	if ans <= n {
		return ans
	}
	return 0
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}
