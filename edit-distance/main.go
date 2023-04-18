package main

import "fmt"

//自底向上的动态规划
func minDistance(word1 string, word2 string) int {
	var length1, length2 = len(word1), len(word2)
	if length1 == 0 || length2 == 0 { //递归出口
		return max(length1, length2) //只做插入操作就能把空串变成目标字符串（同理，对目标字符串只进行删除操作也能变成空串
	}
	var dp = make([][]int, length1+1)
	for i := 0; i <= length1; i++ {
		dp[i] = make([]int, length2+1)
		dp[i][0] = i
	}
	for j := 0; j <= length2; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= length1; i++ {
		for j := 1; j <= length2; j++ {
			left := dp[i-1][j] + 1
			bottom := dp[i][j-1] + 1
			leftBottom := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				leftBottom += 1
			}
			dp[i][j] = min(left, bottom, leftBottom)
		}
	}
	return dp[length1][length2]
}

//自顶向下的递归
//func minDistance(word1 string, word2 string) int {
//	var length1, length2 = len(word1), len(word2)
//	var distance func(int, int) int
//	var cache = make([][]int, length1+1)
//	for i := 0; i <= length1; i++ {
//		cache[i] = make([]int, length2+1)
//		for j := 0; j <= length2; j++ {
//			cache[i][j] = -1
//		}
//	}
//	distance = func(len1 int, len2 int) (val int) {
//		defer func() {
//			cache[len1][len2] = val //缓存
//		}()
//		if cache[len1][len2] != -1 {
//			return cache[len1][len2]
//		}
//		if len1 == 0 || len2 == 0 { //递归出口
//			return max(len1, len2) //只做插入操作就能把空串变成目标字符串（同理，对目标字符串只进行删除操作也能变成空串
//		}
//		if word1[len1-1] == word2[len2-1] { //最后一个字符相同，只需比较前面的字符串
//			return distance(len1-1, len2-1)
//		}
//		return 1 + min( //否则把插入，删除，替换操作都尝试一遍，找到三种方案的最小值再加一就是当前问题的解
//			distance(len1-1, len2-1),
//			distance(len1, len2-1),
//			distance(len1-1, len2),
//		)
//	}
//	return distance(length1, length2)
//}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b, c int) int {
	var minVal = a
	if minVal > b {
		minVal = b
	}
	if minVal > c {
		minVal = c
	}
	return minVal
}

func main() {
	fmt.Println(minDistance("horse", "ros"))
	fmt.Println(minDistance("intention", "execution"))
	fmt.Println(minDistance("", ""))
	fmt.Println(minDistance("a", "a"))
	fmt.Println(minDistance("sea", "ate"))
	fmt.Println(minDistance("dinitrophenylhydrazine",
		"benzalphenylhydrazone"))
}
