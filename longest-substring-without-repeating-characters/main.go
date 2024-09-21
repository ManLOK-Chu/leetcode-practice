package main

import "fmt"

//滑动窗口
func lengthOfLongestSubstring(s string) int {
	var length = len(s)
	if length == 0 {
		return 0
	}
	var tmp = make(map[byte]int)
	var start, result int
	for i := 0; i < length; i++ {
		if index, ok := tmp[s[i]]; ok && start <= index {
			start = index + 1
		} else {
			result = max(result, i-start+1)
		}
		tmp[s[i]] = i
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("babad"))
	fmt.Println(lengthOfLongestSubstring("cbbd"))
}
