package main

import (
	"container/heap"
	"fmt"
)

/*
给你一个 m x n 的矩阵，其中的值均为非负整数，代表二维高度图每个单元的高度，请计算图中形状最多能接多少体积的雨水。
*/
func trapRainWater(heightMap [][]int) int {
	//最小堆
	// 方块 (i,j) 的接水后的高度为：
	// water[i][j] = max(heightMap[i][j], min(water[i−1][j], water[i+1][j], water[i][j−1], water[i][j+1]))
	// 方块 (i,j) 实际接水的容量计算公式为 water[i][j]−heightMap[i][j]
	// 因为最外层的方块无法接水，因此最外层的方块 water[i][j]=heightMap[i][j]。
	var res int
	m, n := len(heightMap), len(heightMap[0])
	if m <= 2 || n <= 2 {
		return res
	}

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	h := &hp{}
	for i, row := range heightMap {
		for j, v := range row {
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				heap.Push(h, cell{v, i, j})
				visited[i][j] = true
			}
		}
	}

	dirs := []int{-1, 0, 1, 0, -1}
	for h.Len() > 0 {
		cur := heap.Pop(h).(cell)
		for k := 0; k < 4; k++ {
			nx, ny := cur.x+dirs[k], cur.y+dirs[k+1]
			if 0 <= nx && nx < m && 0 <= ny && ny < n && !visited[nx][ny] {
				if heightMap[nx][ny] < cur.h {
					res += cur.h - heightMap[nx][ny]
				}
				visited[nx][ny] = true
				heap.Push(h, cell{max(heightMap[nx][ny], cur.h), nx, ny})
			}
		}
	}
	return res
}

type cell struct{ h, x, y int }
type hp []cell

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].h < h[j].h }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(cell)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(trapRainWater([][]int{
		{1, 4, 3, 1, 3, 2}, {3, 2, 1, 3, 2, 4}, {2, 3, 3, 2, 3, 1},
	}))
	fmt.Println(trapRainWater([][]int{
		{3, 3, 3, 3, 3}, {3, 2, 2, 2, 3}, {3, 2, 1, 2, 3}, {3, 2, 2, 2, 3}, {3, 3, 3, 3, 3},
	}))
}
