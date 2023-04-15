package main

import "fmt"

//时间复杂度O（log n)
//空间复杂度O（1）
func singleNonDuplicate(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] == nums[mid^1] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return nums[low]
}

//func singleNonDuplicate(nums []int) int {
//	var n = len(nums)
//	for i := 0; i < n-1; i += 2 {
//		if nums[i] == nums[i+1] {
//			continue
//		}
//		return nums[i]
//	}
//	return nums[n-1]
//}

func main() {
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
	fmt.Println(singleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11}))
	fmt.Println(singleNonDuplicate([]int{1, 3, 3, 7, 7, 11, 11}))
	fmt.Println(singleNonDuplicate([]int{3, 3, 7, 7, 11, 11, 12}))
}
