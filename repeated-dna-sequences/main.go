package main

import "fmt"

//DNA序列 由一系列核苷酸组成，缩写为 'A', 'C', 'G' 和 'T'.
//给定一个表示 DNA序列 的字符串 s ，返回所有在 DNA 分子中出现不止一次的 长度为 10 的序列(子字符串)。你可以按 任意顺序 返回答案。
//时间复杂度： O(nL)，L为子串长度
//空间复杂度: O(nL)
func findRepeatedDnaSequences(s string) []string {
	//统计 s 所有长度为 10 的子串的出现次数，返回所有出现次数超过10的子串
	const length = 10
	var n = len(s)
	if n <= 10 {
		return []string{}
	}
	var cnt = make(map[string]int)
	var result []string
	for i := 0; i+length <= n; i++ {
		subStr := s[i : i+length]
		cnt[subStr]++
		if cnt[subStr] == 2 { //只统计当前出现次数为 2 的子串
			result = append(result, subStr)
		}
	}
	return result
}

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAAAA"))
}
