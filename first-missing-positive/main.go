package main

import "fmt"

/*
*
给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
*/
func firstMissingPositive(nums []int) int {
	//原地哈希
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			// 满足在指定范围内、并且没有放在正确的位置上，才交换
			// 例如：数值 3 应该放在索引 2 的位置上
			swap(nums[i]-1, i, nums)
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	// 都正确则返回数组长度 + 1
	return n + 1
}

func swap(i, j int, nums []int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
}
