package main

import "fmt"

var results [][]int

func combine(n int, k int) [][]int {
	results = [][]int{}
	backtracking(n, k, 1, []int{})
	return results
}

func backtracking(n, k, startIndex int, track []int) {
	if len(track) == k {
		temp := make([]int, k)
		copy(temp, track)
		results = append(results, temp)
	}
	if len(track)+n-startIndex+1 < k {
		return
	}
	for i := startIndex; i <= n; i++ {
		track = append(track, i)
		backtracking(n, k, i+1, track)
		track = track[:len(track)-1]
	}
}

func main() {
	fmt.Println(combine(4, 2))
	fmt.Println(combine(1, 1))
}
