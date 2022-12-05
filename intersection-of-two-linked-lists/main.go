package main

import "fmt"

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
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var tailA, tailB = headA, headB
	var lengthA, lengthB int
	var dummyHeadA = &ListNode{Next: headA}
	var dummyHeadB = &ListNode{Next: headB}
	for tailA.Next != nil {
		lengthA++
		tailA = tailA.Next
	}
	for tailB.Next != nil {
		lengthB++
		tailB = tailB.Next
	}
	var curA = dummyHeadA.Next
	var curB = dummyHeadA.Next
	if lengthA < lengthB {
		for
	}
}

func main() {
	fmt.Println(getIntersectionNode(initListNode([]int{4, 1, 8, 4, 5}), initListNode([]int{5, 6, 1, 8, 4, 5})))
	fmt.Println(getIntersectionNode(initListNode([]int{1, 9, 1, 2, 4}), initListNode([]int{3, 2, 4})))
	fmt.Println(getIntersectionNode(initListNode([]int{2, 6, 4}), initListNode([]int{1, 5})))
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
