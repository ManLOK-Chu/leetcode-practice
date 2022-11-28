package main

import "fmt"

//var resultMap = make(map[int]int, 45)
//
//func climbStairs(n int) int {
//	if n <= 2 {
//		return n
//	}
//	if val, exist := resultMap[n]; exist {
//		return val
//	} else {
//		result := climbStairs(n-1) + climbStairs(n-2)
//		resultMap[n] = result
//		return result
//	}
//}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	var res, fn1, fn2 = 0, 1, 2
	for i := 3; i <= n; i++ {
		res = fn1 + fn2
		fn1 = fn2
		fn2 = res
	}
	return fn2
}

func main() {
	fmt.Println(climbStairs(3))
}
