package main

import (
	"fmt"
	"strings"
)

var result []string
var path []string

func restoreIpAddresses(s string) []string {
	result = nil
	path = nil
	backtracking(s, 0)
	return result
}

func backtracking(s string, startIndex int) {
	if len(path) == 4 { // 够四段后就不再继续往下递归
		// 判断第四段子字符串是否合法，如果合法就放进result中
		if startIndex == len(s) {
			result = append(result, strings.Join(path, "."))
		}
		return
	}
	for i := startIndex; i < len(s); i++ {
		if isValid(s, startIndex, i) {
			path = append(path, s[startIndex:i+1])
			backtracking(s, i+1)
			path = path[:len(path)-1]
		} else { // 如果不满足条件，再往后也不可能满足条件，直接退出
			break
		}
	}
}

func isValid(s string, start, end int) bool {
	if s[start] == '0' && start != end {
		return false
	}
	var num int
	for i := start; i <= end; i++ {
		if s[i] > '9' || s[i] < '0' {
			return false
		}
		num = num*10 + int(s[i]-'0')
		if num > 255 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(restoreIpAddresses("0000"))
	fmt.Println(restoreIpAddresses("101023"))
}
