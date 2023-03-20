package main

import (
	"fmt"
	"sort"
)

type IntSlice []int

func (s IntSlice) Len() int {
	return len(s)
}

func (s IntSlice) Less(i, j int) bool {
	if s[i] > s[j] {
		return true
	}
	return false
}

func (s IntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//O(n)
func findKthLargest(nums []int, k int) int {
	var array IntSlice = nums
	sort.Sort(array) //快排O(nlogn)
	return array[k-1]
}

//桶排序
//func findKthLargest(nums []int, k int) int {
//	var res int
//	hash := make([]int, 20001)
//	for i := range nums {
//		hash[nums[i]+10000]++
//	}
//	for index := 20000; index >= 0; index-- {
//		if hash[index] != 0 {
//			k -= hash[index]
//			if k <= 0 {
//				res = index - 10000
//				break
//			}
//		}
//	}
//	return res
//}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
