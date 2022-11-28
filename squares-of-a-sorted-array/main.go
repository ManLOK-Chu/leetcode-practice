package main

import "fmt"

func sortedSquares(nums []int) []int {
	var length = len(nums)
	var result = make([]int, length)
	var index, left, right = length - 1, 0, length - 1
	for left <= right {
		r := nums[right] * nums[right]
		l := nums[left] * nums[left]
		if r > l {
			result[index] = r
			index--
			right--
		} else {
			result[index] = l
			index--
			left++
		}
	}
	return result
}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11}))
}
