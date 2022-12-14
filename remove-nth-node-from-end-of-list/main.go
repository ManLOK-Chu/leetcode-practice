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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//stack := list.New()
	//var dummyHead = &ListNode{Next: head}
	//var cur = dummyHead.Next
	//stack.PushBack(dummyHead)
	//for cur != nil {
	//	stack.PushBack(cur)
	//	cur = cur.Next
	//}
	//var pre, next *ListNode
	//for n > 0 {
	//	next = stack.Remove(stack.Back()).(*ListNode)
	//	n--
	//}
	//pre = stack.Back().Value.(*ListNode)
	//pre.Next = next.Next
	//return dummyHead.Next
	dummyHead := &ListNode{}
	dummyHead.Next = head
	cur := head
	prev := dummyHead
	i := 1
	for cur != nil {
		cur = cur.Next
		if i > n {
			prev = prev.Next
		}
		i++
	}
	prev.Next = prev.Next.Next
	return dummyHead.Next
}

func main() {
	fmt.Println(removeNthFromEnd(initListNode([]int{1, 2, 3, 4, 5}), 2))
	fmt.Println(removeNthFromEnd(initListNode([]int{1}), 1))
	fmt.Println(removeNthFromEnd(initListNode([]int{1, 2}), 1))
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
