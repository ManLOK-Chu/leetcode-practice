package main

import "fmt"

//func spiralOrder(matrix [][]int) []int {
//	if len(matrix) == 0 || len(matrix[0]) == 0 {
//		return []int{}
//	}
//	m := len(matrix)
//	n := len(matrix[0])
//	visited := make([][]bool, m)
//	for i := 0; i < m; i++ {
//		visited[i] = make([]bool, n)
//	}
//	total := m * n
//	var row, col int
//	var directions = [][]int{
//		{0, 1},
//		{1, 0},
//		{0, -1},
//		{-1, 0},
//	}
//	var directionIndex = 0
//	var result = make([]int, total)
//	for i := 0; i < total; i++ {
//		result[i] = matrix[row][col]
//		visited[row][col] = true
//		nextRow, nextCol := row+directions[directionIndex][0], col+directions[directionIndex][1]
//		if nextRow < 0 || nextRow >= m || nextCol < 0 || nextCol >= n || visited[nextRow][nextCol] {
//			directionIndex = (directionIndex + 1) % 4
//		}
//		row += directions[directionIndex][0]
//		col += directions[directionIndex][1]
//
//	}
//	return result
//}
func spiralOrder(matrix [][]int) []int {
	var left, right, top, bottom = 0, len(matrix[0]) - 1, 0, len(matrix) - 1
	var total = len(matrix[0]) * len(matrix)
	var result = make([]int, 0, total)
	for total > 0 {
		//left to right
		for i := left; total > 0 && i <= right; i++ {
			result = append(result, matrix[top][i])
			total--
		}
		top++
		//top to bottom
		for i := top; total > 0 && i <= bottom; i++ {
			result = append(result, matrix[i][right])
			total--
		}
		right--
		//right to left
		for i := right; total > 0 && i >= left; i-- {
			result = append(result, matrix[bottom][i])
			total--
		}
		bottom--
		//bottom to left
		for i := bottom; total > 0 && i >= top; i-- {
			result = append(result, matrix[i][left])
			total--
		}
		left++
	}
	return result
}

func main() {
	fmt.Println(spiralOrder(prepare(3, 3)))
	fmt.Println(spiralOrder(prepare(3, 4)))
}

func prepare(x, y int) [][]int {
	var results [][]int
	for i := 0; i < x; i++ {
		arr := make([]int, y)
		for j := 0; j < y; j++ {
			arr[j] = i*y + j + 1
		}
		results = append(results, arr)
	}
	return results
}
