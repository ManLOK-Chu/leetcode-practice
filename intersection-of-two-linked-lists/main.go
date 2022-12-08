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
	var dummyHeadA, dummyHeadB = &ListNode{Next: headA}, &ListNode{Next: headB}
	for tailA.Next != nil {
		lengthA++
		tailA = tailA.Next
	}
	for tailB.Next != nil {
		lengthB++
		tailB = tailB.Next
	}
	var curA, curB = dummyHeadA.Next, dummyHeadB.Next

}

func main() {
	fmt.Println(getIntersectionNode(initListNode([]int{4, 1, 8, 4, 5}, []int{5, 6, 1, 8, 4, 5})))
	fmt.Println(getIntersectionNode(initListNode([]int{1, 9, 1, 2, 4}, []int{3, 2, 4})))
	fmt.Println(getIntersectionNode(initListNode([]int{2, 6, 4}, []int{1, 5})))
}

func initListNode(arrA []int, arrB []int) (*ListNode, *ListNode) {
	var dummyHeadA, dummyHeadB = &ListNode{}, &ListNode{}
	var cur = dummyHeadA
	for i := 0; i < len(arrA); i++ {
		cur.Next = &ListNode{Val: arrA[i]}
		cur = cur.Next
	}
	cur = dummyHeadB
	for i := 0; i < len(arrB); i++ {
		cur.Next = &ListNode{Val: arrB[i]}
		cur = cur.Next
	}
	return dummyHeadA.Next, dummyHeadB.Next
}
