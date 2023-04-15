package main

import "fmt"

var results [][]int

func permute(nums []int) [][]int {
	results = nil
	var paths []int
	var excluded = make([]bool, len(nums))
	backtracking(nums, paths, excluded)
	return results
}

func backtracking(nums []int, paths []int, excluded []bool) {
	if len(paths) == len(nums) {
		var tmp = make([]int, len(paths))
		copy(tmp, paths)
		results = append(results, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if excluded[i] {
			continue
		}
		paths = append(paths, nums[i])
		excluded[i] = true
		backtracking(nums, paths, excluded)
		paths = paths[0 : len(paths)-1] //回滚到上一步操作
		excluded[i] = false             //回滚到上一步操作
	}
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permute([]int{0, 1}))
	fmt.Println(permute([]int{1}))
}
