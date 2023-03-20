package main

import "fmt"

var results [][]string
var solution [][]byte

func solveNQueens(n int) [][]string {
	if n == 0 {
		return nil
	}
	results = nil
	solution = make([][]byte, n)
	for i := 0; i < n; i++ {
		solution[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			solution[i][j] = '.'
		}
	}
	backtracking(n, 0)
	return results
}

func backtracking(n int, row int) {
	if row == n {
		var arr = make([]string, n)
		for i := range solution {
			arr[i] = string(solution[i])
		}
		results = append(results, arr)
		return
	}
	for col := 0; col < n; col++ {
		if isValid(row, col, n) {
			solution[row][col] = 'Q'
			backtracking(n, row+1)
			solution[row][col] = '.'
		}
	}
}

func isValid(row, col, n int) bool {
	// 检查列
	for i := 0; i < row; i++ { // 这是一个剪枝
		if solution[i][col] == 'Q' {
			return false
		}
	}
	// 检查 45度角是否有皇后
	for i, j := row-1, col-1; i >= 0 && j >= 0; {
		if solution[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}
	// 检查 135度角是否有皇后
	for i, j := row-1, col+1; i >= 0 && j < n; {
		if solution[i][j] == 'Q' {
			return false
		}
		i--
		j++
	}
	return true
}

func main() {
	fmt.Println(solveNQueens(4))
	fmt.Println(solveNQueens(1))
}
