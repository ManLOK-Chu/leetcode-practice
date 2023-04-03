package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var res = make([][]int, 0)
	sort.Ints(nums) // O(nlogn)
	var length = len(nums)
	// O(n^2)
	for i := 0; i < length; i++ {
		if i > 0 && nums[i] == nums[i-1] { //排序后跳过重复值
			continue
		}
		target := -nums[i] // 在 i + 1 ... nums.length - 1 中查找相加等于 -nums[i] 的两个数
		left := i + 1
		right := length - 1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				//去重
				for left < right {
					left++
					if nums[left-1] != nums[left] {
						break
					}
				}
				for left < right {
					right--
					if nums[right] != nums[right+1] {
						break
					}
				}
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return res
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 1, 1}))
	fmt.Println(threeSum([]int{0, 0, 0}))
}
