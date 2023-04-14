package main

import (
	"container/list"
	"fmt"
)

//单调栈
func dailyTemperatures(temperatures []int) []int {
	var n = len(temperatures)
	if n == 1 {
		return []int{0}
	}
	var result = make([]int, n)
	stack := list.New()
	for i := 0; i < n; i++ {
		for stack.Len() > 0 {
			front := stack.Back().Value.(int)
			if temperatures[front] > temperatures[i] { //维护单调递减栈
				break
			}
			stack.Remove(stack.Back()) //出栈
			result[front] = i - front
		}
		stack.PushBack(i) //入栈
	}
	return result
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	fmt.Println(dailyTemperatures([]int{30, 40, 50, 60}))
	fmt.Println(dailyTemperatures([]int{30, 60, 90}))
}
