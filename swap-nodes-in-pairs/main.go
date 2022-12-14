package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) String() string {
	cur := node
	var b strings.Builder
	for cur != nil {
		b.WriteString(strconv.Itoa(cur.Val))
		b.WriteByte(' ')
		cur = cur.Next
	}
	return b.String()
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	var dummyHead = &ListNode{Next: head}
	var first, second, third, cur *ListNode = nil, nil, nil, dummyHead
	for cur.Next != nil && cur.Next.Next != nil {
		first, second, third = cur.Next, cur.Next.Next, cur.Next.Next.Next
		cur.Next = second
		second.Next = first
		first.Next = third
		cur = first
	}
	return dummyHead.Next
}

func main() {
	fmt.Println(swapPairs(initListNode([]int{1, 2, 3, 4})))
	fmt.Println(swapPairs(initListNode([]int{1})))
}

func initListNode(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	var dummyHead = &ListNode{}
	var cur = dummyHead
	for i := range arr {
		cur.Next = &ListNode{Val: arr[i]}
		cur = cur.Next
	}
	return dummyHead.Next
}
