package main

import (
	"container/list"
	"fmt"
)

//s只存在'('和')'
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
