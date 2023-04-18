package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//时间复杂度: O(n logn)
//空间复杂度：O(logn)
func sortList(head *ListNode) *ListNode {
	var mergeSort func(*ListNode, *ListNode) *ListNode
	mergeSort = func(head, tail *ListNode) *ListNode {
		if head == nil {
			return head
		}
		if head.Next == tail {
			head.Next = nil
			return head
		}
		slow, fast := head, head
		for fast != tail {
			slow = slow.Next
			fast = fast.Next
			if fast != tail {
				fast = fast.Next
			}
		}

		mid := slow
		return merge(mergeSort(head, mid), mergeSort(mid, tail))
	}
	return mergeSort(head, nil)
}

func merge(head1, head2 *ListNode) *ListNode {
	var dummyHead = &ListNode{}
	cur, cur1, cur2 := dummyHead, head1, head2
	for cur1 != nil && cur2 != nil {
		if cur1.Val <= cur2.Val {
			cur.Next = cur1
			cur1 = cur1.Next
		} else {
			cur.Next = cur2
			cur2 = cur2.Next
		}
		cur = cur.Next
	}
	if cur1 != nil {
		cur.Next = cur1
	} else if cur2 != nil {
		cur.Next = cur2
	}
	return dummyHead.Next
}

func main() {

}
