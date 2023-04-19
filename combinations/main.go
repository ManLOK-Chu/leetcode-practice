package main

import "fmt"

//给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
//你可以按 任何顺序 返回答案。
func combine(n int, k int) [][]int {
	var results [][]int
	var track []int
	var backtracking func(int)
	backtracking = func(start int) {
		if len(track) == k {
			temp := make([]int, k)
			copy(temp, track)
			results = append(results, temp)
		}
		if len(track)+n-start+1 < k {
			return
		}
		for i := start; i <= n; i++ {
			track = append(track, i)
			backtracking(i + 1)
			track = track[:len(track)-1]
		}
	}
	backtracking(1)
	return results
}

func main() {
	fmt.Println(combine(4, 2))
	fmt.Println(combine(1, 1))
}
