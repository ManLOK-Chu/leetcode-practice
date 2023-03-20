package main

import (
	"container/heap"
	"fmt"
)

type IntList []int

func (p IntList) Len() int {
	return len(p)
}

func (p IntList) Less(i, j int) bool { //大顶堆
	return p[i] > p[j]
}

func (p IntList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *IntList) Push(x interface{}) {
	*p = append(*p, x.(int))
}

func (p *IntList) Pop() interface{} {
	list := *p
	var length = len(list)
	val := list[length-1]
	*p = list[:length-1]
	return val
}

func lastStoneWeight(stones []int) int {
	var stoneList = new(IntList)
	heap.Init(stoneList)
	for _, v := range stones {
		heap.Push(stoneList, v)
	}
	for stoneList.Len() > 1 {
		y := heap.Pop(stoneList).(int)
		x := heap.Pop(stoneList).(int)
		if x != y {
			heap.Push(stoneList, y-x)
		}
	}
	if stoneList.Len() == 0 {
		return 0
	}
	return heap.Pop(stoneList).(int)
}

func main() {
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
	fmt.Println(lastStoneWeight([]int{1, 2, 3}))
}
