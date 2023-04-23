package main

import (
	"container/list"
	"fmt"
)

//给定一个二进制数组 nums 和一个整数 k，如果可以翻转最多 k 个 0 ，则返回 数组中连续 1 的最大个数

//维护一个窗口 [left,right] 来容纳至少k个0。当遇到了0，就累加 zero 的个数，然后判断如果此时0的个数大于k，则右移左边界left，
//如果移除掉的 nums[left] 为0，那么 zero 自减1。如果不大于k，则用窗口中数字的个数来更新 res
//因为 nums[left] 需要访问之前的数字。可以将遇到的0的位置全都保存下来，这样需要移动 left 的时候就知道移到哪里了
func longestOnes(nums []int, k int) int {
	var res, left int
	queue := list.New()
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			queue.PushBack(right)
		}
		if queue.Len() > k {
			left = queue.Remove(queue.Front()).(int) + 1
		}
		res = max(res, right-left+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))
}
