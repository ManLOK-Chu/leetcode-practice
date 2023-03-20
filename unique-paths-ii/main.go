package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var m = len(obstacleGrid)
	var n = len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 { //起始或终点位置有障碍，直接返回0
		return 0
	}
	var dp = make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for j := 0; j < n; j++ { //初始化第一行，遇到障碍则停止
		if obstacleGrid[0][j] == 1 {
			break
		}
		dp[0][j] = 1
	}
	for i := 0; i < m; i++ { //初始化第一列，遇到障碍则停止
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0}}))
	fmt.Println(uniquePathsWithObstacles([][]int{
		{0, 1},
		{0, 0}}))
}
