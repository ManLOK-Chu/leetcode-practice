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
func removeElements(head *ListNode, val int) *ListNode {
	var dummyHead = &ListNode{Next: head}
	var cur = dummyHead
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}

func main() {
	fmt.Println(removeElements(NewLinkedList([]int{1, 2, 6, 3, 4, 5, 6}), 6))
	fmt.Println(removeElements(NewLinkedList([]int{}), 1))
	fmt.Println(removeElements(NewLinkedList([]int{7, 7, 7, 7}), 7))
	fmt.Println(removeElements(NewLinkedList([]int{1, 2, 2, 1}), 2))
}

func NewLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	var head = &ListNode{Val: arr[0]}
	var pNode = head
	for i := 1; i < len(arr); i++ {
		pNode.Next = &ListNode{Val: arr[i]}
		pNode = pNode.Next
	}
	return head
}
