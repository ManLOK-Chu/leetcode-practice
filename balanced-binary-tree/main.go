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
func isBalanced(root *TreeNode) bool {
	var f func(root *TreeNode) (int, bool)
	f = func(root *TreeNode) (int, bool) {
		if root == nil {
			return 0, true
		}

		leftHeight, balanced := f(root.Left)
		if !balanced {
			return 0, false
		}

		rightHeight, balanced := f(root.Right)
		if !balanced {
			return 0, false
		}

		if abs(leftHeight-rightHeight) > 1 {
			return 0, false
		}

		return 1 + max(leftHeight, rightHeight), true
	}

	_, balanced := f(root)
	return balanced
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

//func isBalanced(root *TreeNode) bool {
//	if root == nil {
//		return true
//	}
//	if !isBalanced(root.Left) || !isBalanced(root.Right) { //它的左子树应该是一棵平衡二叉树,右子树应该是一棵平衡二叉树
//		return false
//	}
//	leftHeight := maxDepth(root.Left) + 1
//	rightHeight := maxDepth(root.Right) + 1
//	if leftHeight-rightHeight > 1 || rightHeight-leftHeight > 1 { //任意节点左右子树高度之差绝对值小于1
//		return false
//	}
//	return true
//}
//
//func maxDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
//}
//

func main() {
	fmt.Println(isBalanced(arrayToTreeNode([]int{3, 9, 20, null, null, 15, 7})))
	fmt.Println(isBalanced(arrayToTreeNode([]int{1, 2, 2, 3, 3, null, null, 4, 4})))
	fmt.Println(isBalanced(arrayToTreeNode([]int{})))
	fmt.Println(isBalanced(arrayToTreeNode([]int{1, 2, 3, 4, 5, 6, null, 8})))
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
