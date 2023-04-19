package main

import "fmt"

//给你四个整数数组 nums1、nums2、nums3 和 nums4 ，数组长度都是 n ，请你计算有多少个元组 (i, j, k, l) 能满足：
//0 <= i, j, k, l < n
//nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	var result int
	count := make(map[int]int)
	for _, a := range nums1 {
		for _, b := range nums2 {
			count[a+b]++
		}
	}
	for _, a := range nums3 {
		for _, b := range nums4 {
			result += count[-a-b]
		}
	}
	return result
}

func main() {
	fmt.Println(fourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}))
}
