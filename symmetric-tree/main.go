package main

import (
	"container/list"
	"fmt"
	"math"
)

const NULL = math.MinInt

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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	stack := list.New()
	stack.PushFront(root.Left)
	stack.PushFront(root.Right)
	var l, r *list.Element
	for stack.Len() > 0 {
		if l = stack.Front(); l != nil {
			stack.Remove(l)
		}
		if r = stack.Front(); r != nil {
			stack.Remove(r)
		}
		left := l.Value.(*TreeNode)
		right := r.Value.(*TreeNode)
		if left == nil && nil == right {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		stack.PushFront(left.Left)
		stack.PushFront(right.Right)
		stack.PushFront(left.Right)
		stack.PushFront(right.Left)
	}
	return true
}

//func isSymmetric(root *TreeNode) bool {
//	if root == nil {
//		return true
//	}
//	return compare(root.Left, root.Right)
//}
//
//func compare(left *TreeNode, right *TreeNode) bool {
//	if left == nil && right == nil {
//		return true
//	} else if left != nil && right == nil {
//		return false
//	} else if left == nil && right != nil {
//		return false
//	} else if left.Val != right.Val {
//		return false
//	}
//	return compare(left.Left, right.Right) && compare(left.Right, right.Left)
//}

func main() {
	root := arrayToTreeNode([]int{1, 2, 2, 3, 4, 4, 3})
	fmt.Println(isSymmetric(root))
	root = arrayToTreeNode([]int{1, 2, 2, NULL, 3, NULL, 3})
	fmt.Println(isSymmetric(root))
	root = arrayToTreeNode([]int{1, 2, 2, 2, NULL, 2})
	fmt.Println(isSymmetric(root))
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
			if arr[i] != NULL {
				node.Left = &TreeNode{Val: arr[i]}
				queue.PushBack(node.Left) //队尾进队
			}
			isLeft = false
		} else {
			if arr[i] != NULL {
				node.Right = &TreeNode{Val: arr[i]}
				queue.PushBack(node.Right) //队尾进队
			}
			queue.Remove(element)
			isLeft = true
		}
	}
	return root
}
