package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

var _exist struct{}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	//var cur = head
	//var m = make(map[*ListNode]struct{})
	//var ok bool
	//for cur != nil {
	//	if _, ok = m[cur]; ok {
	//		return cur
	//	}
	//	m[cur] = _exist
	//	cur = cur.Next
	//}
	//return nil
	var fast, slow = head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			var cur = head
			for slow != cur {
				slow = slow.Next
				cur = cur.Next
			}
			return cur
		}
	}
	return nil
}

func main() {
	fmt.Println(detectCycle(initCycleListNode([]int{3, 2, 0, -4}, 1)))
	fmt.Println(detectCycle(initCycleListNode([]int{1, 2}, 0)))
	fmt.Println(detectCycle(initCycleListNode([]int{1}, -1)))
}

func initCycleListNode(arr []int, pos int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	var nodes = make([]*ListNode, len(arr))
	var dummyHead = &ListNode{}
	var cur = dummyHead
	for i := 0; i < len(arr); i++ {
		cur.Next = &ListNode{Val: arr[i]}
		nodes[i] = cur.Next
		cur = cur.Next
	}
	if pos != -1 {
		cur.Next = nodes[pos]
	}
	return dummyHead.Next
}
