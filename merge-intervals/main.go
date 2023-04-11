package main

import (
	"fmt"
	"sort"
)

type IntArrArray [][]int

func (arr IntArrArray) Len() int {
	return len(arr)
}

func (arr IntArrArray) Less(i, j int) bool {
	if arr[i][0] < arr[j][0] {
		return true
	}
	return false
}

func (arr IntArrArray) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// 时间复杂度 O(n Log n)
// 空间复杂度 O(log n)
func merge(intervals [][]int) [][]int {
	length := len(intervals)
	if length == 0 {
		return nil
	}
	var intArr IntArrArray = intervals
	sort.Sort(intArr) // 先按照区间起始位置排序
	var merged [][]int
	merged = append(merged, intArr[0])
	for i := 1; i < length; i++ {
		left := intArr[i][0]
		right := intArr[i][1]
		last := merged[len(merged)-1]
		// 如果结果数组是空的，或者当前区间的起始位置 > 结果数组中最后区间的终止位置，
		// 则不合并，直接将当前区间加入结果数组。
		if last[1] < left {
			merged = append(merged, []int{left, right})
		} else { // 反之将当前区间合并至结果数组的最后区间
			last[1] = max(last[1], right)
		}
	}
	return merged
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
}
