package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var length = len(s)
	var j = 0
	next := make([]int, length)
	next[0] = j
	for i := 1; i < length; i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
	fmt.Println(next)
	return 0
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("babad"))
	fmt.Println(lengthOfLongestSubstring("cbbd"))
}
