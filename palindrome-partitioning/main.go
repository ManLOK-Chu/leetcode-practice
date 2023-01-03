package main

import "fmt"

var results [][]string
var path []string

func partition(s string) [][]string {
	results = nil
	path = nil
	backtracking(s, 0)
	return results
}

func backtracking(s string, startIndex int) {
	if startIndex == len(s) {
		tmp := make([]string, len(path))
		copy(tmp, path)
		results = append(results, tmp)
		return
	}
	for i := startIndex; i < len(s); i++ {
		if isPalindrome(s, startIndex, i) {
			path = append(path, s[startIndex:i+1])
		} else {
			continue
		}
		backtracking(s, i+1)
		path = path[:len(path)-1]
	}
}

func isPalindrome(s string, start, end int) bool {
	for i, j := start, end; i < j; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(partition("aab"))
	fmt.Println(partition("a"))
}
