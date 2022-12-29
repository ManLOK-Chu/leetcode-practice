package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	if len(s) == 0 {
		return 0
	}
	sort.Ints(g)
	sort.Ints(s)
	index := len(s) - 1
	var result int
	for i := len(g) - 1; i >= 0 && index >= 0; i-- {
		if s[index] >= g[i] {
			result++
			index--
		}
	}
	return result
}

func main() {
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{1, 1}))
	fmt.Println(findContentChildren([]int{1, 2}, []int{1, 2, 3}))
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{3}))
}
