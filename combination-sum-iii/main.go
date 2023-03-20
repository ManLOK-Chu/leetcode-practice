package main

import "fmt"

var path []int
var results [][]int

func combinationSum3(k int, n int) [][]int {
	path = nil
	results = nil
	backtracking(n, k, 0, 1)
	return results
}

func backtracking(target, k, sum, start int) {
	if len(path) == k {
		if sum == target {
			results = append(results, path)
		}
		return
	}
	for i := start; i <= 9; i++ {
		sum += i
		path = append(path, i)
		backtracking(target, k, sum, i+1)
		sum -= i
		path = path[:len(path)-1]
	}
}

func main() {
	fmt.Println(combinationSum3(3, 7))
	fmt.Println(combinationSum3(3, 9))
	fmt.Println(combinationSum3(4, 1))
}
