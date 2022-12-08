package main

import (
	"fmt"
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
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var cur = headA
	var lengthA, lengthB = 1, 1
	for cur.Next != nil {
		lengthA++
		cur = cur.Next
	}
	cur = headB
	for cur.Next != nil {
		lengthB++
		cur = cur.Next
	}
	var longCur, shortCur = headA, headB
	var gap = lengthA - lengthB
	if lengthB > lengthA {
		longCur, shortCur = headB, headA
		gap = lengthB - lengthA
	}
	for gap > 0 {
		longCur = longCur.Next
		gap--
	}
	for longCur != nil {
		if longCur == shortCur {
			return longCur
		}
		longCur = longCur.Next
		shortCur = shortCur.Next
	}
	return nil
	//var m = make(map[*ListNode]bool)
	//var cur = headA
	//for cur != nil {
	//	m[cur] = true
	//	cur = cur.Next
	//}
	//cur = headB
	//for cur != nil {
	//	if m[cur] {
	//		return cur
	//	}
	//	cur = cur.Next
	//}
	//return nil
}

func main() {
	fmt.Println(getIntersectionNode(initListNode([]int{4, 1, 8, 4, 5}, []int{5, 6, 1, 8, 4, 5}, 2, 3)))
	fmt.Println(getIntersectionNode(initListNode([]int{1, 9, 1, 2, 4}, []int{3, 2, 4}, 3, 1)))
	fmt.Println(getIntersectionNode(initListNode([]int{2, 6, 4}, []int{1, 5}, 3, 2)))
}

func initListNode(arrA, arrB []int, skipA, skipB int) (*ListNode, *ListNode) {
	var dummyHeadA, dummyHeadB = &ListNode{}, &ListNode{}
	var listA = make([]*ListNode, len(arrA))
	var cur = dummyHeadA
	for i := 0; i < len(arrA); i++ {
		cur.Next = &ListNode{Val: arrA[i]}
		cur = cur.Next
		listA[i] = cur
	}
	cur = dummyHeadB
	for i := 0; i < skipB; i++ {
		cur.Next = &ListNode{Val: arrB[i]}
		cur = cur.Next
	}
	if skipB < len(arrB) {
		cur.Next = listA[skipA]
	}
	return dummyHeadA.Next, dummyHeadB.Next
}
