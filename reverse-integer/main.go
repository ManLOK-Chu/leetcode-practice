package main

import (
	"fmt"
	"math"
)

/*
func reverse(x int) int {
	var result int32
	var val = x
	var neg, skipZero bool
	var stack = list.New()
	if x < 0 {
		val = -x
		neg = true
	}
	stack.PushBack(val % 10)
	val /= 10
	for val != 0 {
		stack.PushBack(val % 10)
		val /= 10
	}
	for stack.Len() != 0 {
		v := stack.Remove(stack.Front()).(int)
		if v == 0 && skipZero {
			continue
		} else if v != 0 {
			skipZero = false
		}
		result = 10*result + int32(v)
	}
	if result < 0 {
		return 0
	}
	if neg {
		return -int(result)
	}
	//if result > math.MaxInt32 || result < math.MinInt32 {
	//	return 0
	//}
	return int(result)
}
*/

func reverse(x int) int {
	var result int
	result = 0
	for x != 0 {
		temp := x % 10
		x /= 10
		result *= 10
		result += temp
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-321))
	fmt.Println(reverse(120))
	fmt.Println(reverse(1534236469))
}
