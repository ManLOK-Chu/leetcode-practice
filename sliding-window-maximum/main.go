package main

import (
	"fmt"
)

type Queue struct {
	queue []int
	size  int
}

/*
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
返回 滑动窗口中的最大值
*/
func maxSlidingWindow(nums []int, k int) []int {
	var result []int
	if k == 1 {
		return nums
	} else {
		queue := NewQueue()
		for i := 0; i < k; i++ {
			queue.Push(nums[i])
		}
		result = append(result, queue.Front())
		for i := k; i < len(nums); i++ {
			queue.Pop(nums[i-k])
			queue.Push(nums[i])
			result = append(result, queue.Front())
		}
	}
	return result
}

func NewQueue() *Queue {
	q := new(Queue)
	q.queue = make([]int, 0, 32)
	return q
}

// Pop 每次弹出的时候，比较当前要弹出的数值是否等于队列出口元素的数值，如果相等则弹出。
// 同时pop之前判断队列当前是否为空。
func (q *Queue) Pop(val int) {
	if q.size > 0 && q.queue[0] == val {
		q.queue = q.queue[1:]
		q.size--
	}
}

// Push 如果push的数值大于入口元素的数值，那么就将队列后端的数值弹出，直到push的数值小于等于队列入口元素的数值为止。
// 这样就保持了队列里的数值是单调从大到小的了。
func (q *Queue) Push(val int) {
	for q.size > 0 && val > q.queue[q.size-1] {
		q.queue = q.queue[:q.size-1]
		q.size--
	}
	q.queue = append(q.queue, val)
	q.size++
}

func (q *Queue) Front() int {
	return q.queue[0]
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow([]int{1}, 1))
}
