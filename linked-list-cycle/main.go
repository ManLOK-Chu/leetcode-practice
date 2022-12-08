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
func hasCycle(head *ListNode) bool {
	//var m = make(map[*ListNode]bool)
	//var cur = head
	//for cur != nil {
	//	if _, exist := m[cur]; exist {
	//		return true
	//	}
	//	m[cur] = true
	//	cur = cur.Next
	//}
	//return false
	var slow, fast = head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(hasCycle(initCycleListNode([]int{3, 2, 0, -4}, 1)))
	fmt.Println(hasCycle(initCycleListNode([]int{1, 2}, 0)))
	fmt.Println(hasCycle(initCycleListNode([]int{1}, -1)))
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
