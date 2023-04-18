package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) String() string {
	var cur = node
	var b strings.Builder
	for cur != nil {
		b.WriteString(strconv.Itoa(cur.Val))
		b.WriteByte(' ')
		cur = cur.Next
	}
	return b.String()
}

//给你一个链表数组，每个链表都已经按升序排列。
//请你将所有链表合并到一个升序链表中，返回合并后的链表。
func mergeKLists(lists []*ListNode) *ListNode {
	var n = len(lists)
	if n == 0 {
		return nil
	}
	var dummyHead = &ListNode{}
	var cur = dummyHead
	var empty = make([]bool, n)
	for {
		var minIndex, minVal = -1, math.MaxInt
		for i := 0; i < n; i++ {
			if !empty[i] && lists[i] != nil { //当第i个链表不为空时
				if minVal > lists[i].Val {
					minVal = lists[i].Val
					minIndex = i
				}
			} else {
				empty[i] = true
			}
		}
		if minIndex == -1 {
			break
		}
		cur.Next = &ListNode{Val: minVal}
		cur = cur.Next
		lists[minIndex] = lists[minIndex].Next //前进一步
	}
	return dummyHead.Next
}

func main() {
	fmt.Println(mergeKLists(prepareKLists([][]int{
		{1, 4, 5}, {1, 3, 4}, {2, 6},
	})))
	fmt.Println(mergeKLists(prepareKLists([][]int{})))
	fmt.Println(mergeKLists(prepareKLists([][]int{
		{},
	})))
}

func prepareKLists(arr [][]int) []*ListNode {
	var results []*ListNode
	for i := 0; i < len(arr); i++ {
		results = append(results, prepareLists(arr[i]))
	}
	return results
}

func prepareLists(arr []int) *ListNode {
	var dummyHead = &ListNode{}
	var cur = dummyHead
	for j := 0; j < len(arr); j++ {
		cur.Next = &ListNode{Val: arr[j]}
		cur = cur.Next
	}
	return dummyHead.Next
}
