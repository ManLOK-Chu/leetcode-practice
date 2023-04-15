package main

import "fmt"

func getMaxMatrix(matrix [][]int) []int {
	var m, n = len(matrix), len(matrix[0])
	var max = matrix[0][0]
	var res = make([]int, 4)
	//构造列的前缀和
	var preSum = make([][]int, m+1)
	preSum[0] = make([]int, n)
	for i := 1; i < m+1; i++ {
		preSum[i] = make([]int, n)
		for j := 0; j < n; j++ {
			preSum[i][j] = preSum[i-1][j] + matrix[i-1][j]
		}
	}

	//合并行
	for top := 0; top < m; top++ {
		for bottom := top; bottom < m; bottom++ {
			//构建一维数组
			arr := make([]int, n)
			for i := 0; i < n; i++ {
				arr[i] = preSum[bottom+1][i] - preSum[top][i]
			}
			//求最大子数组
			start := 0
			sum := arr[0]
			for i := 1; i < n; i++ {
				if sum > 0 {
					sum += arr[i]
				} else {
					sum = arr[i]
					start = i
				}
				if sum > max {
					max = sum
					res[0] = top
					res[1] = start
					res[2] = bottom
					res[3] = i
				}
			}

		}
	}
	return res
}

func main() {
	fmt.Println(getMaxMatrix([][]int{{9, -8, 1, 3, -2}, {-3, 7, 6, -2, 4}, {6, -4, -4, 8, -7}}))
}
