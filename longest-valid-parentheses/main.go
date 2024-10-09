package main

import (
	"container/list"
	"fmt"
)

// 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号 子串 的长度。
// s只存在'('和')'
func longestValidParentheses(s string) int {
	var length = len(s)
	if length <= 1 {
		return 0
	}
	var result = 0
	stack := list.New()
	stack.PushBack(-1)
	for i := 0; i < length; i++ {
		if s[i] == '(' {
			stack.PushBack(i)
		} else {
			stack.Remove(stack.Back())
			if stack.Len() == 0 {
				stack.PushBack(i)
			} else {
				last := stack.Back().Value.(int)
				result = max(result, i-last)
			}
		}
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
	fmt.Println(longestValidParentheses("(()"))
	fmt.Println(longestValidParentheses(")()())"))
	fmt.Println(longestValidParentheses(")("))
	fmt.Println(longestValidParentheses("()(())"))
}
