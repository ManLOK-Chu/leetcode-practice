package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	var short, long = nums1, nums2
	var lengthShort, lengthLong = len(nums1), len(nums2)
	if lengthShort > lengthLong {
		short, long = nums2, nums1
		lengthShort, lengthLong = lengthLong, lengthShort
	}
	var result = make([]int, 0, lengthShort)
	var table = [1001]int{}
	for i := 0; i < lengthLong; i++ {
		table[long[i]] = 1
	}
	var key int
	for i := 0; i < lengthShort; i++ {
		key = short[i]
		if table[key] == 1 {
			result = append(result, key)
			table[key] = 0
		}
	}
	return result
}

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
