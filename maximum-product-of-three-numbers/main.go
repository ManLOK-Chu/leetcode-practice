package main

import (
	"fmt"
	"sort"
)

//首先将数组排序。
//如果数组中全是非负数，则排序后最大的三个数相乘即为最大乘积；如果全是非正数，则最大的三个数相乘同样也为最大乘积。
//如果数组中有正数有负数，则最大乘积既可能是三个最大正数的乘积，也可能是两个最小负数（即绝对值最大）与最大正数的乘积。
func maximumProduct(nums []int) int {
	var n = len(nums)
	sort.Ints(nums)
	result := nums[n-1] * nums[n-2] * nums[n-3]
	if nums[0] > 0 { //说明数组不存在负数
		return result
	}
	return max(nums[0]*nums[1]*nums[n-1], result)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maximumProduct([]int{-100, -98, -1, 2, 3, 4}))
	fmt.Println(maximumProduct([]int{1, 2, 3, 4, 5, 6}))
}
