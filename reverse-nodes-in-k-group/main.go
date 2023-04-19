package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) String() string {
	var res []int
	var cur = node
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	return fmt.Sprint(res)
}

//给你链表的头节点 head ，每k个节点一组进行翻转，请你返回修改后的链表。
//k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是k的整数倍，那么请将最后剩余的节点保持原有顺序。
//你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
func reverseKGroup(head *ListNode, k int) *ListNode {
	var dummyHead = &ListNode{Next: head}
	var pre, cur = dummyHead, head
	for cur != nil {
		tail := pre
		for i := 0; i < k; i++ { //检查剩余长度是否大于等于k
			tail = tail.Next
			if tail == nil {
				return dummyHead.Next
			}
			//if tail != nil && tail.Next == nil { //不足k个也要翻转
			//	break
			//}
		}
		//暂存子链表的下一个节点
		tmp := tail.Next
		cur, tail = reverseList(cur, tail) //翻转一个子链表，返回新的头尾节点
		//把翻转后的子链表拼接回原来的链表
		pre.Next = cur
		tail.Next = tmp
		pre = tail
		cur = tmp
	}
	return dummyHead.Next
}

func reverseList(head, tail *ListNode) (*ListNode, *ListNode) {
	var cur, pre, tmp *ListNode = head, nil, nil
	for pre != tail {
		tmp = cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return tail, head
}

func main() {
	fmt.Println(reverseKGroup(initListNode([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}), 4))
	fmt.Println(reverseKGroup(initListNode([]int{1, 2, 3, 4, 5}), 2))
	fmt.Println(reverseKGroup(initListNode([]int{1, 2, 3, 4, 5}), 3))
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
