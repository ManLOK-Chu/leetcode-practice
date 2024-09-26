package main

import "fmt"

/*
*
给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。
*/
func longestCommonSubsequence(text1 string, text2 string) int {
	//二维动态规划
	m, n := len(text1), len(text2)
	//dp[i][j] 表示text1[0:i]和text2[0:j]到最长公共子序列长度
	var dp = make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	//空字符串和任意字符串的最长公共子序列必为0
	dp[0][0] = 0
	// 当text1[i-1]==text2[j-1]时，相同的字符称为公共字符，再增加一个字符即可得到最长公共子序列，
	// 所以 dp[i][j] = dp[i-1][j-1] + 1
	// 当text1[i-1]!=text2[j-1]时，取两项中长度较大的一项
	// dp[i][j] = max(dp[i-1][j], dp[i][j-1])
	for i, c1 := range text1 {
		for j, c2 := range text2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestCommonSubsequence("abcde", "ace")) //3
	fmt.Println(longestCommonSubsequence("abc", "abc"))   //3
	fmt.Println(longestCommonSubsequence("abc", "def"))   //0
}
