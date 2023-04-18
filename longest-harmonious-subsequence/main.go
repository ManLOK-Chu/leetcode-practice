package main

import "fmt"

//和谐数组是指一个数组里元素的最大值和最小值之间的差别 正好是 1 。
//现在，给你一个整数数组 nums ，请你在所有可能的子序列中找到最长的和谐子序列的长度。
func findLHS(nums []int) int {
	cnt := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		cnt[nums[i]]++
	}
	var max int
	for key, val := range cnt {
		if c := cnt[key+1]; c > 0 && c+val > max {
			max = c + val
		}
	}
	return max
}

func main() {
	fmt.Println(findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}))
}
