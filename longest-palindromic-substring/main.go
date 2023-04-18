package main

import (
	"fmt"
)

//动态规划
//dp[i][j] = (s[i] == s[j]) && dp[i+1][j-1]
//边界条件：（j-1)-(i+1) + 1 < 2 整理后 j - i < 3 即s[i:j]的长度为2或者3，不用检查子串是否为回文子串
func longestPalindrome(s string) string {
	var length = len(s)
	if length < 2 {
		return s
	}
	//一个回文子串去除了首尾后剩下的仍然还是回文子串
	//dp[i][j]表示s[i:j]是否为回文子串(左闭右闭
	var dp = make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
		dp[i][i] = true //长度为1的子串必为回文子串
	}
	var maxLen = 1
	var begin = 0
	for l := 2; l <= length; l++ { //l为子串长度，从长度较短的字符串向长度较长的字符串进行转移
		// 枚举左边界，左边界的上限设置可以宽松一些
		for i := 0; i < length; i++ {
			//从子串长度l和i确定右边边界
			j := l + i - 1
			if j >= length { //右边界越界退出当前循环
				break
			}
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if (j-1)-(i+1)+1 < 2 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			// 只要 dp[i][l] == true 成立，就表示子串 s[i..L] 是回文，此时记录回文长度和起始位置
			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}

//暴力法O(n^3)
//func longestPalindrome(s string) string {
//	var length = len(s)
//	if length < 2 {
//		return s
//	}
//	var maxLen = 1
//	var begin = 0
//	for i := 0; i < length-1; i++ {
//		for j := i + 1; j < length; j++ {
//			if j-i+1 > maxLen && validPalindromic(s, i, j) {
//				maxLen = j - i + 1
//				begin = i
//			}
//		}
//	}
//	return s[begin : begin+maxLen]
//}
//
//func validPalindromic(s string, left, right int) bool {
//	for left < right {
//		if s[left] != s[right] {
//			return false
//		}
//		left++
//		right--
//	}
//	return true
//}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
}
