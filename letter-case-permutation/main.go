package main

import (
	"fmt"
)

func letterCasePermutation(s string) []string {
	var results []string
	var length = len(s)
	var backtracking func([]byte, int)
	backtracking = func(paths []byte, begin int) {
		results = append(results, string(paths))
		for i := begin; i < length; i++ {
			if paths[i] < 'A' {
				continue
			}
			paths[i] ^= 0b100000 //只有第六位不同，通过^32可以将 0（1）改变 1（0）
			backtracking(paths, i+1)
			paths[i] ^= 0b100000
		}
	}
	backtracking([]byte(s), 0)
	return results
}

func main() {
	fmt.Println(letterCasePermutation("a1b2"))
}
