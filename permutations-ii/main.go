package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	var results [][]int
	var paths []int
	sort.Ints(nums) //排序
	var visited = make([]bool, len(nums))
	var backtracking func(int)
	backtracking = func(n int) {
		if n == len(nums) {
			var tmp = append([]int{}, paths...)
			results = append(results, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if visited[i] {
				continue
			}
			if i > 0 && !visited[i-1] && nums[i-1] == nums[i] { //核心
				continue
			}
			paths = append(paths, nums[i])
			visited[i] = true
			backtracking(n + 1)
			paths = paths[:len(paths)-1]
			visited[i] = false
		}
	}
	backtracking(0)
	return results
}

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
	fmt.Println(permuteUnique([]int{1, 2, 3}))

}
