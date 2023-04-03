package main

import "fmt"

func numIslands(grid [][]byte) int {
	var nr = len(grid)
	var nc = len(grid[0])
	var num int
	for row := 0; row < nr; row++ {
		for column := 0; column < nc; column++ {
			if grid[row][column] == '1' {
				num++
				dfs(grid, row, column)
			}
		}
	}
	return num
}

func dfs(grid [][]byte, row, column int) {
	var nr = len(grid)
	var nc = len(grid[0])

	if row < 0 || column < 0 || row >= nr || column >= nc || grid[row][column] == '0' {
		return
	}

	grid[row][column] = '0'
	dfs(grid, row-1, column)
	dfs(grid, row+1, column)
	dfs(grid, row, column-1)
	dfs(grid, row, column+1)
}

func main() {
	fmt.Println(numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'}}))

	fmt.Println(numIslands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'}}))
}
