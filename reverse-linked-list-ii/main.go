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
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var dummyHead = &ListNode{Next: head}
	var cur = dummyHead
	var pre, last, front *ListNode
	for i := 1; i <= left-1; i++ {
		cur = cur.Next
	}
	pre = cur
	last = cur.Next
	for i := left; i <= right; i++ {
		cur = pre.Next
		pre.Next = cur.Next
		cur.Next = front
		front = cur
	}
	cur = pre.Next
	pre.Next = front
	last.Next = cur
	return dummyHead.Next
}

func main() {
	fmt.Println(reverseBetween(initListNode([]int{1, 2, 3, 4, 5}), 2, 4))
	fmt.Println(reverseBetween(initListNode([]int{5}), 1, 1))
}

func initListNode(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	var head = &ListNode{Val: arr[0]}
	var cur = head
	for i := 1; i < len(arr); i++ {
		cur.Next = &ListNode{Val: arr[i]}
		cur = cur.Next
	}
	return head
}
