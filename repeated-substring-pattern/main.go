package main

import (
	"fmt"
)

//func repeatedSubstringPattern(s string) bool {
//	var length = len(s)
//	var subStr string
//	var contain bool
//	for i := 1; i < length; i++ {
//		if length%i != 0 {
//			continue
//		}
//		subStr = s[:i]
//		contain = true
//		for k := i; k < length; k += i {
//			if subStr != s[k:k+i] {
//				contain = false
//				break
//			}
//		}
//		if contain {
//			return true
//		}
//	}
//	return false
//}

//func repeatedSubstringPattern(s string) bool {
//	var length = len(s)
//	var t = s[1:] + s[:length-1]
//	return strings.Contains(t, s)
//}

func repeatedSubstringPattern(s string) bool {
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
	// next[length-1]  最长相同前后缀的长度
	if next[length-1] != 0 && length%(length-next[length-1]) == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(repeatedSubstringPattern("bb"))
	fmt.Println(repeatedSubstringPattern("abab"))
	fmt.Println(repeatedSubstringPattern("aba"))
	fmt.Println(repeatedSubstringPattern("abcabcabcabc"))
}
