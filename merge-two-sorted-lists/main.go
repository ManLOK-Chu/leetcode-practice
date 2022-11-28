package main

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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	return nil
}

func main() {
	initListNode("124")
}

func initListNode(str string) *ListNode {
	var begin, current *ListNode
	for i := range str {
		current.Val = int(str[i]) - 32
	}
	return begin
}
