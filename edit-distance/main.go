package main

import "fmt"

//自顶向下的递归
func minDistance(word1 string, word2 string) int {
	var length1, length2 = len(word1), len(word2)
	var distance func(int, int) int
	var cache = make([][]int, length1+1)
	for i := 0; i <= length1; i++ {
		cache[i] = make([]int, length2+1)
		for j := 0; j <= length2; j++ {
			cache[i][j] = -1
		}
	}
	distance = func(len1 int, len2 int) (val int) {
		defer func() {
			cache[len1][len2] = val //缓存
		}()
		if cache[len1][len2] != -1 {
			return cache[len1][len2]
		}
		if len1 == 0 || len2 == 0 { //递归出口
			return max(len1, len2) //只做插入操作就能把空串变成目标字符串（同理，对目标字符串只进行删除操作也能变成空串
		}
		if word1[len1-1] == word2[len2-1] { //最后一个字符相同，只需比较前面的字符串
			return distance(len1-1, len2-1)
		}
		return 1 + min( //否则把插入，删除，替换操作都尝试一遍，找到三种方案的最小值再加一就是当前问题的解
			distance(len1-1, len2-1),
			distance(len1, len2-1),
			distance(len1-1, len2),
		)
	}
	return distance(length1, length2)
}

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
	fmt.Println(minDistance("dinitrophenylhydrazine",
		"benzalphenylhydrazone"))
}
