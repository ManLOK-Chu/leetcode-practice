package main

import "fmt"

func search(nums []int, target int) int {
	var length = len(nums)
	var left, right = 0, length - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			// 左半边有序
			if nums[0] <= target && target <= nums[mid] {
				//target的值在左半边
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			//右半边有序
			if nums[mid] <= target && target <= nums[length-1] {
				//target的值在右半边
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	fmt.Println(search([]int{1}, 0))
	fmt.Println(search([]int{3, 1}, 1))
}
