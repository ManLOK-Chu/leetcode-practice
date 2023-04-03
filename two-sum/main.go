package main

import "fmt"

//func twoSum(nums []int, target int) (result []int) {
//	var m = make(map[int]int, len(nums))
//	result = make([]int, 2)
//	for i, v := range nums {
//		m[v] = i
//	}
//	for i := range nums {
//		result[0] = i
//		val := target - nums[i]
//		if j, ok := m[val]; ok && j != i {
//			result[1] = j
//			return
//		}
//		//for j := range nums {
//		//	if j == i || val != nums[j] {
//		//		continue
//		//	}
//		//	result[1] = j
//		//	return
//		//}
//	}
//	return
//}
func twoSum(nums []int, target int) []int {
	var m = make(map[int]int, len(nums))
	for i := range nums {
		m[nums[i]] = i
	}
	for i := range nums {
		if val, ok := m[target-nums[i]]; ok && val != i {
			return []int{i, val}
		}
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}
