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
//func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
//	var dummyHead = &ListNode{}
//	var node1, node2, cur = list1, list2, dummyHead
//	for node1 != nil && node2 != nil {
//		var l1, l2 int
//		if node1 == nil {
//			l1 = 101
//			l2 = node2.Val
//		} else if node2 == nil {
//			l1 = node1.Val
//			l2 = 101
//		} else {
//			l1 = node1.Val
//			l2 = node2.Val
//		}
//		if l1 < l2 {
//			cur.Next = &ListNode{Val: l1}
//			if node1 != nil {
//				node1 = node1.Next
//			}
//		} else {
//			cur.Next = &ListNode{Val: l2}
//			if node2 != nil {
//				node2 = node2.Next
//			}
//		}
//		cur = cur.Next
//	}
//	if node1 != nil {
//		cur.Next = node1
//	}
//	if node2 != nil {
//		cur.Next = node2
//	}
//	return dummyHead.Next
//}
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var currA = list1
	var currB = list2
	var dummy = &ListNode{}
	var cur = dummy
	for currA != nil && currB != nil {
		valA := currA.Val
		valB := currB.Val
		if valA > valB {
			currB = currB.Next
			cur.Next = &ListNode{Val: valB}
		} else {
			currA = currA.Next
			cur.Next = &ListNode{Val: valA}
		}
		cur = cur.Next
	}
	if currA != nil {
		cur.Next = currA
	}
	if currB != nil {
		cur.Next = currB
	}
	return dummy.Next
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
	fmt.Println(mergeTwoLists(initListNode([]int{1, 2, 4}), initListNode([]int{1, 3, 4})))
	fmt.Println(mergeTwoLists(initListNode([]int{}), initListNode([]int{})))
	fmt.Println(mergeTwoLists(initListNode([]int{}), initListNode([]int{0})))
	fmt.Println(mergeTwoLists(initListNode([]int{5}), initListNode([]int{1, 2, 4})))
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
