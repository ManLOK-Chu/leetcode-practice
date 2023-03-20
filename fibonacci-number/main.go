package main

import "fmt"

func fib(n int) int {
	if n < 2 {
		return n
	}
	var dp = make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
	fmt.Println(fib(30))
	fmt.Println(fib(0))
	fmt.Println(fib(1))
}
