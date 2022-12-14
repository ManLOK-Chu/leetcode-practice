package main

import (
	"container/list"
	"fmt"
)

func evalRPN(tokens []string) int {
	stack := list.New()
	var length = len(tokens)
	for i := 0; i < length; i++ {
		switch tokens[i] {
		case "+":
			a := stack.Remove(stack.Back()).(int)
			b := stack.Remove(stack.Back()).(int)
			stack.PushBack(a + b)
		case "-":
			a := stack.Remove(stack.Back()).(int)
			b := stack.Remove(stack.Back()).(int)
			stack.PushBack(b - a)
		case "*":
			a := stack.Remove(stack.Back()).(int)
			b := stack.Remove(stack.Back()).(int)
			stack.PushBack(a * b)
		case "/":
			a := stack.Remove(stack.Back()).(int)
			b := stack.Remove(stack.Back()).(int)
			stack.PushBack(b / a)
		default:
			stack.PushBack(stringToInt(tokens[i]))
		}
	}
	return stack.Back().Value.(int)
}

func stringToInt(str string) int {
	var i, result int
	var length = len(str)
	if str[0] == '-' {
		i = 1
	}
	for i < length {
		result = 10*result + int(str[i]-'0')
		i++
	}
	if str[0] == '-' {
		return -result
	}
	return result
}

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
