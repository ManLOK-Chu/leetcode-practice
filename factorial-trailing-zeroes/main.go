package main

import "fmt"

/*
给定一个整数 n ，返回 n! 结果中尾随零的数量。
提示 n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1
*/
func trailingZeroes(n int) int {
	// n! 尾零的数量即为 n! 中因子 10 的个数，而 10=2×5，因此转换成求 n! 中质因子 2 的个数和质因子 5 的个数的较小值。
	// 由于质因子 5 的个数不会大于质因子 2 的个数，我们可以仅考虑质因子 5 的个数。
	var res int
	for i := 5; i <= n; i += 5 {
		for x := i; x%5 == 0; x /= 5 {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(trailingZeroes(3))
	fmt.Println(trailingZeroes(5))
	fmt.Println(trailingZeroes(0))
}
