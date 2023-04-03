package main

import (
	"fmt"
)

//暴力法O(n^3)
func longestPalindrome(s string) string {
	var length = len(s)
	if length < 2 {
		return s
	}
	var maxLen = 1
	var begin = 0
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if j-i+1 > maxLen && validPalindromic(s, i, j) {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}

func validPalindromic(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
}
