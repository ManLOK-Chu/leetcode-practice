package main

import (
	"container/list"
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sum, carry int
	var cur1, cur2 = l1, l2
	var queue = list.New()
	for cur1 != nil || cur2 != nil {
		sum = carry
		if cur1 != nil {
			sum += cur1.Val
			cur1 = cur1.Next
		}
		if cur2 != nil {
			sum += cur2.Val
			cur2 = cur2.Next
		}
		queue.PushBack(sum % 10)
		carry = sum / 10
	}
	if carry != 0 {
		queue.PushBack(carry)
	}
	var dummyHead = &ListNode{}
	var cur = dummyHead
	for queue.Len() > 0 {
		val := queue.Remove(queue.Front()).(int)
		cur.Next = &ListNode{Val: val}
		cur = cur.Next
	}
	return dummyHead.Next
}

func main() {
	fmt.Println(addTwoNumbers(initListNode([]int{2, 4, 3}), initListNode([]int{5, 6, 4})))
	fmt.Println(addTwoNumbers(initListNode([]int{0}), initListNode([]int{0})))
	fmt.Println(addTwoNumbers(initListNode([]int{9, 9, 9, 9, 9, 9, 9}), initListNode([]int{9, 9, 9, 9})))
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
