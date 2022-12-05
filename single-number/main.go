package main

import "fmt"

//var _exist struct{}
//
//func singleNumber(nums []int) int {
//	var length = len(nums)
//  if length == 1 {
//     return nums[0]
//  }
//	var m = make(map[int]struct{}, length/2+1)
//	var key int
//	for i := 0; i < length; i++ {
//		key = nums[i]
//		if _, ok := m[key]; !ok { //不存在
//			m[key] = _exist
//		} else { //已存在，删去
//			delete(m, key)
//		}
//	}
//	for k := range m {
//		return k
//	}
//	return -1
//}

func singleNumber(nums []int) int {
	ans := 0
	for _, val := range nums {
		ans = ans ^ val
	}
	return ans
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println(singleNumber([]int{1}))
}
