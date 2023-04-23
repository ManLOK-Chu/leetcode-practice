package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	var cnt, max = 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			cnt++
		} else {
			cnt = 0
			continue
		}
		if cnt > max {
			max = cnt
		}
	}
	return max
}

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
	fmt.Println(findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1}))
}
