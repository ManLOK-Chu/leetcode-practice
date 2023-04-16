package main

import "fmt"

//时间复杂度 O(n^3)
//空间复杂度：O(n)
func wordBreak(s string, wordDict []string) bool {
	var n = len(s)
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}

// 转化为 求装满背包s的前几位字符的方式有几种
//func wordBreak(s string, wordDict []string) bool {
//	var n = len(s)
//	// 装满背包s的前几位字符的方式有几种
//	dp := make([]int, n+1)
//	dp[0] = 1
//	for i := 0; i <= len(s); i++ { // 背包
//		for j := 0; j < len(wordDict); j++ { // 物品
//			if i >= len(wordDict[j]) && wordDict[j] == s[i-len(wordDict[j]):i] {
//				dp[i] += dp[i-len(wordDict[j])]
//			}
//		}
//	}
//
//	return dp[n] > 0
//}

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}
