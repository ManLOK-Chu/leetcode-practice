package main

import (
	"container/heap"
	"fmt"
)

type Pair struct {
	Key int
	Val int
}

type PairList []*Pair

func (p PairList) Len() int {
	return len(p)
}

func (p PairList) Less(i, j int) bool {
	return p[i].Val < p[j].Val
}

func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PairList) Push(x interface{}) {
	*p = append(*p, x.(*Pair))
}

func (p *PairList) Pop() interface{} {
	list := *p
	var length = len(list)
	val := list[length-1]
	*p = list[:length-1]
	return val
}

//func topKFrequent(nums []int, k int) []int {
//	var length = len(nums)
//	var m = make(map[int]int)
//	for i := 0; i < length; i++ {
//		m[nums[i]]++
//	}
//	length = len(m)
//	var pairs = make(PairList, 0, length)
//	for key, v := range m {
//		pairs = append(pairs, &Pair{Key: key, Val: v})
//	}
//	sort.Sort(pairs)
//	var results = make([]int, k)
//	for k > 0 {
//		k--
//		results[k] = pairs[length-1-k].Key
//	}
//	return results
//}

func topKFrequent(nums []int, k int) []int {
	var length = len(nums)
	var m = make(map[int]int)
	for i := 0; i < length; i++ {
		m[nums[i]]++
	}
	var result = make([]int, k)
	var pairs = new(PairList)
	heap.Init(pairs)
	for key, v := range m {
		heap.Push(pairs, &Pair{Key: key, Val: v})
		if pairs.Len() > k {
			heap.Pop(pairs)
		}
	}
	for k > 0 {
		k--
		p := heap.Pop(pairs).(*Pair)
		result[k] = p.Key
	}
	return result
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{7, 9, 7, 1, 2, 3, 7, 1, 2, 7, 4, 9, 8, 3, 4, 7, 7, 2, 7, 4, 8, 3}, 3))
	fmt.Println(topKFrequent([]int{1}, 1))
}
