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
	var cur = node
	var b strings.Builder
	for cur != nil {
		b.WriteString(strconv.Itoa(cur.Val))
		b.WriteByte(' ')
		cur = cur.Next
	}
	return b.String()
}

/*
*
给你两个链表 list1 和 list2 ，它们包含的元素分别为 n 个和 m 个。
请你将 list1 中下标从 a 到 b 的全部节点都删除，并将list2 接在被删除节点的位置。
*/
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	/*在实现上，我们使用两个指针 p 和 q，初始时均指向链表 list1 的头节点。
	然后我们向后移动指针 p 和 q，直到指针 p 指向链表 list1 中第 a 个节点的前一个节点，指针 q 指向链表 list1 中第 b 个节点。
	此时我们将 p 的 next 指针指向链表 list2 的头节点，将链表 list2 的尾节点的 next 指针指向 q 的 next 指针指向的节点，即可完成题目要求。
	*/
	p, q := list1, list1
	for i := a; i > 1; i-- {
		p = p.Next
	}
	for i := b; i > 0; i-- {
		q = q.Next
	}
	p.Next = list2
	for p.Next != nil {
		p = p.Next
	}
	p.Next = q.Next
	q.Next = nil
	return list1
}

func main() {
	fmt.Println(mergeInBetween(
		prepareLists([]int{10, 1, 13, 6, 9, 5}),
		3, 4,
		prepareLists([]int{1000000, 1000001, 1000002})))
}

func prepareLists(arr []int) *ListNode {
	var dummyHead = &ListNode{}
	var cur = dummyHead
	for j := 0; j < len(arr); j++ {
		cur.Next = &ListNode{Val: arr[j]}
		cur = cur.Next
	}
	return dummyHead.Next
}
