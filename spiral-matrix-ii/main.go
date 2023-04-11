package main

import "fmt"

func generateMatrix(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}
	if n == 1 {
		return [][]int{{1}}
	}
	var result = make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	var left, right, top, bottom = 0, n - 1, 0, n - 1
	var num, total = 1, n * n
	for num <= total {
		for i := left; i <= right; i++ { // left to right.
			result[top][i] = num
			num++
		}
		top++
		for i := top; i <= bottom; i++ { // top to bottom.
			result[i][right] = num
			num++
		}
		right--
		for i := right; i >= left; i-- { //right to left
			result[bottom][i] = num
			num++
		}
		bottom--
		for i := bottom; i >= top; i-- { //bottom to top
			result[i][left] = num
			num++
		}
		left++
	}
	return result
}

func main() {
	fmt.Println(generateMatrix(3))
	fmt.Println(generateMatrix(1))
}
