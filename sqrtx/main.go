package main

import "fmt"

func mySqrt(x int) int {
	i := 0
	for ; i <= x/2+1; i++ {
		result := i * i
		if result == x {
			return i
		}
		if result > x {
			break
		}
	}
	return i - 1
}

func main() {
	fmt.Println(mySqrt(1))
	fmt.Println(mySqrt(2))
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
}
