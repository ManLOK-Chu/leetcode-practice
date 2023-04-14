package main

import (
	"fmt"
	"sort"
)

var results [][]int

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates) //先按从小到大排个序
	var paths []int
	results = nil
	backtracking(paths, candidates, target, 0)
	return results
}

func backtracking(paths []int, candidates []int, target int, begin int) {
	if target < 0 {
		return
	}
	if target == 0 {
		var tmp = make([]int, len(paths))
		copy(tmp, paths)
		results = append(results, tmp)
	}
	for i := begin; i < len(candidates); i++ {
		if len(paths) != 0 && target < candidates[i] {
			break //结束本次迭代
		}
		paths = append(paths, candidates[i])
		backtracking(paths, candidates, target-candidates[i], i)
		paths = paths[0 : len(paths)-1]
	}
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
	fmt.Println(combinationSum([]int{2}, 1))
}
