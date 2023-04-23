package main

import "fmt"

//给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
//解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
//迭代
func subsets(nums []int) [][]int {
	var n = len(nums)
	var results = make([][]int, 0, 1<<n) //2^n个
	for i := 0; i < 1<<n; i++ {
		var arr = make([]int, 0, i)
		for j := 0; j < n; j++ {
			if i>>j&1 > 0 {
				arr = append(arr, nums[j])
			}
		}
		results = append(results, arr)
	}
	return results
}

//递归
//func subsets(nums []int) [][]int {
//	var n = len(nums)
//	var results = make([][]int, 0, 1<<n) //2^n个
//	var paths []int
//	var backtracking func(int)
//	backtracking = func(start int) {
//		if start == n {
//			var tmp = make([]int, len(paths))
//			copy(tmp, paths)
//			results = append(results, tmp)
//			return
//		}
//		paths = append(paths, nums[start])
//		backtracking(start + 1)
//		paths = paths[:len(paths)-1]
//		backtracking(start + 1)
//	}
//	backtracking(0)
//	return results
//}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets([]int{0}))
}
