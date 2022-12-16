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
func repeatedSubstringPattern(s string) bool {
	var length = len(s)
	var t = s[1:] + s[:length-1]
	var newLength = (length << 1) - 2
	for i := 0; i+length <= newLength; i++ {
		if t[i:i+length] == s {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(repeatedSubstringPattern("bb"))
	fmt.Println(repeatedSubstringPattern("abab"))
	fmt.Println(repeatedSubstringPattern("aba"))
	fmt.Println(repeatedSubstringPattern("abcabcabcabc"))
}
