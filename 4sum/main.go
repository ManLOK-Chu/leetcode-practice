package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	var res = make([][]int, 0)
	var length = len(nums)
	sort.Ints(nums)
	for i := 0; i < length; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < length; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left := j + 1
			right := length - 1
			for left < right {
				if left > j+1 && nums[left] == nums[left-1] {
					left++
					continue
				}
				if right < length-1 && nums[right] == nums[right+1] {
					right--
					continue
				}
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					right--
				} else if sum > target {
					right--
				} else if sum < target {
					left++
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
}
