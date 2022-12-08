package main

import (
	"container/list"
	"fmt"
	"math"
)

func myAtoi(s string) int {
	var result int
	var neg bool
	var i, length = 0, len(s)
	queue := list.New()
	for i < length {
		if !isNumber(s[i]) {
			i++
			continue
		}
		if i != 0 && s[i-1] == '-' {
			neg = true
		} else if i != 0 && s[i-1] == ' ' {
			return 0
		}
		break
	}
	queue.PushBack(s[i] - '0')
	i++
	for i < length {
		if !isNumber(s[i]) {
			break
		}
		queue.PushBack(s[i] - '0')
		i++
	}
	for queue.Len() > 0 {
		val := queue.Remove(queue.Front()).(byte)
		result = 10*result + int(val)
	}
	if neg {
		result = -result
	}
	if result > math.MaxInt32 {
		return math.MaxInt32
	} else if result < math.MinInt32 {
		return math.MinInt32
	}
	return result
}

func isNumber(c byte) bool {
	return c >= '0' && c <= '9'
}

func main() {
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("   -42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("words and 987"))
}
