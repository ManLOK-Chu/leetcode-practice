package main

import "fmt"

func plusOne(digits []int) []int {
	var length = len(digits)
	digits[length-1]++
	for i := length - 1; i > 0; i-- {
		var process bool
		if digits[i] >= 10 {
			digits[i-1]++
			digits[i] -= 10
			process = true
		}
		if !process {
			break
		}
	}
	if digits[0] >= 10 {
		var result = make([]int, 0, length+1)
		result = append(result, 1)
		digits[0] -= 10
		result = append(result, digits...)
		return result
	}
	return digits
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{9}))
	fmt.Println(plusOne([]int{9, 9, 9, 9, 9, 9}))
}
