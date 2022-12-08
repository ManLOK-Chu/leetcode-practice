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
	//stack := list.New()
	//var cur = head
	//for cur != nil {
	//	stack.PushBack(cur.Val)
	//	cur = cur.Next
	//}
	//var dummyHead = &ListNode{}
	//var resultCur = dummyHead
	//for stack.Len() > 0 {
	//	val := stack.Remove(stack.Back()).(int)
	//	resultCur.Next = &ListNode{Val: val}
	//	resultCur = resultCur.Next
	//}
	//return dummyHead.Next
	var temp, pre, cur *ListNode = nil, nil, head
	for cur != nil {
		temp = cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}

func main() {
	fmt.Println(reverseList(initListNode([]int{5, 4, 3, 2, 1})))
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
