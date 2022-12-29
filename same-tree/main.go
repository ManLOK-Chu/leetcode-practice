package main

import (
	"container/list"
	"fmt"
	"math"
)

const null = math.MinInt32

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p != nil && q == nil) || (p == nil && q != nil) {
		return false
	}
	stack := list.New()
	stack.PushBack(p)
	stack.PushBack(q)
	for stack.Len() > 0 {
		node1 := stack.Remove(stack.Back()).(*TreeNode)
		node2 := stack.Remove(stack.Back()).(*TreeNode)
		if node1.Val != node2.Val {
			return false
		}
		if (node1.Right != nil && node2.Right == nil) || (node1.Right == nil && node2.Right != nil) {
			return false
		} else if node1.Right != nil && node2.Right != nil {
			stack.PushBack(node1.Right)
			stack.PushBack(node2.Right)
		}
		if (node1.Left != nil && node2.Left == nil) || (node1.Left == nil && node2.Left != nil) {
			return false
		} else if node1.Left != nil && node2.Left != nil {
			stack.PushBack(node1.Left)
			stack.PushBack(node2.Left)
		}
	}
	return true
}

//func isSameTree(p *TreeNode, q *TreeNode) bool {
//	if p == nil && q == nil {
//		return true
//	} else if p != nil && q == nil {
//		return false
//	} else if p == nil && q != nil {
//		return false
//	}
//	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(q.Right, p.Right)
//}

func main() {
	fmt.Println(isSameTree(arrayToTreeNode([]int{1, 2, 3}), arrayToTreeNode([]int{1, 2, 3})))
	fmt.Println(isSameTree(arrayToTreeNode([]int{1, 2}), arrayToTreeNode([]int{1, null, 2})))
	fmt.Println(isSameTree(arrayToTreeNode([]int{1, 2, 1}), arrayToTreeNode([]int{1, 1, 2})))
	fmt.Println(isSameTree(arrayToTreeNode([]int{}), arrayToTreeNode([]int{})))
}

func arrayToTreeNode(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	var root = &TreeNode{Val: arr[0]}
	queue := list.New()
	queue.PushBack(root)
	var isLeft = true
	for i := 1; i < len(arr); i++ {
		element := queue.Front() //队头出队
		node := element.Value.(*TreeNode)
		if isLeft {
			if arr[i] != null {
				node.Left = &TreeNode{Val: arr[i]}
				queue.PushBack(node.Left) //队尾进队
			}
			isLeft = false
		} else {
			if arr[i] != null {
				node.Right = &TreeNode{Val: arr[i]}
				queue.PushBack(node.Right) //队尾进队
			}
			queue.Remove(element)
			isLeft = true
		}
	}
	return root
}
