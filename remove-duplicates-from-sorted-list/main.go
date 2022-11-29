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

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	var dummyHead = &ListNode{Next: head, Val: 101}
	var cur = dummyHead
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == cur.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
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

func main() {
	fmt.Println(deleteDuplicates(initListNode([]int{0, 0, 2})))
	fmt.Println(deleteDuplicates(initListNode([]int{1, 1, 2})))
	fmt.Println(deleteDuplicates(initListNode([]int{1, 1, 2, 3, 3})))
}
