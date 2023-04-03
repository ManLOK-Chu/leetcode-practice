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

func (n *ListNode) String() string {
	node := n
	b := strings.Builder{}
	for node != nil {
		b.WriteString(strconv.Itoa(node.Val))
		b.WriteByte(' ')
		node = node.Next
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
func reverseList(head *ListNode) *ListNode {
	var cur, pre, tmp *ListNode = head, nil, nil
	for cur != nil {
		tmp = cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

func main() {
	fmt.Println(reverseList(initListNode([]int{5, 4, 3, 2, 1})))
	fmt.Println(reverseList(initListNode([]int{1, 2})))
	fmt.Println(reverseList(initListNode([]int{})))
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
