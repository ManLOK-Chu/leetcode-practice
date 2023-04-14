package main

import (
	"fmt"
)

//时间复杂度O(n)
//空间复杂度O(n)
func majorityElement(nums []int) int {
	var m = make(map[int]int)
	n := len(nums)
	for i := 0; i < n; i++ {
		m[nums[i]]++
		if m[nums[i]] > n/2 {
			return nums[i]
		}
	}
	return 0
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
