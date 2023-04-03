package main

import "fmt"

//func merge(nums1 []int, m int, nums2 []int, n int) {
//	var indexM, indexN int
//	var result = make([]int, 0, m+n)
//	for indexM < m && indexN < n {
//		if nums1[indexM] < nums2[indexN] {
//			result = append(result, nums1[indexM])
//			indexM++
//		} else {
//			result = append(result, nums2[indexN])
//			indexN++
//		}
//	}
//	if indexM < m {
//		result = append(result, nums1[indexM:m]...)
//	}
//	if indexN < n {
//		result = append(result, nums2[indexN:n]...)
//	}
//	for i := range result {
//		nums1[i] = result[i]
//	}
//}
func merge(nums1 []int, m int, nums2 []int, n int) {
	result := make([]int, 0, m+n)
	var p1, p2 int
	for {
		if p1 == m {
			result = append(result, nums2[p2:]...)
			break
		}
		if p2 == n {
			result = append(result, nums1[p1:]...)
			break
		}
		if nums1[p1] < nums2[p2] {
			result = append(result, nums1[p1])
			p1++
		} else {
			result = append(result, nums2[p2])
			p2++
		}
	}
	copy(nums1, result)
}

func main() {
	var nums1 = []int{1, 2, 3, 0, 0, 0}
	var nums2 = []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)

	nums1 = []int{1}
	nums2 = []int{}
	merge(nums1, 1, nums2, 0)
	fmt.Println(nums1)

	nums1 = []int{0}
	nums2 = []int{1}
	merge(nums1, 0, nums2, 1)
	fmt.Println(nums1)

	nums1 = []int{2, 0}
	nums2 = []int{1}
	merge(nums1, 1, nums2, 1)
	fmt.Println(nums1)
}
