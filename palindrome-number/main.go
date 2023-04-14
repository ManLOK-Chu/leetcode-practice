package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false //排除负数
	}
	if x < 10 {
		return true //个位数必为
	}
	var cur, num = 0, x
	for num != 0 {
		cur = cur*10 + num%10
		num /= 10
	}
	return cur == x
}

//func isPalindrome(x int) bool {
//	if x < 0 {
//		return false //排除负数
//	}
//	if x < 10 {
//		return true //个位数必为
//	}
//	var arr []int
//	for x != 0 {
//		arr = append(arr, x%10)
//		x = x / 10
//	}
//	var i, j = 0, len(arr) - 1
//	for i <= j {
//		if arr[i] != arr[j] {
//			return false
//		}
//		i++
//		j--
//	}
//	return true
//}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(11))
	fmt.Println(isPalindrome(12121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(10))
}
