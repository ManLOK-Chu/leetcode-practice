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
	if root == nil {
		return true
	}
	stack := list.New()
	var pre, cur *TreeNode = nil, root
	for cur != nil && stack.Len() > 0 {
		for cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		}
		node := stack.Back().Value.(*TreeNode)
		if node.Right == nil || node.Right == pre {
			leftHeight := getHeight(node.Left)
			rightHeight := getHeight(node.Right)
			if leftHeight-rightHeight > 1 || rightHeight-leftHeight > 1 {
				return false
			}
			stack.Remove(stack.Back())
			pre = node
			cur = nil
		} else {
			cur = cur.Right
		}
	}
	return true
}

//getHeight 层序遍历，求结点的高度
func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var height int
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		size := queue.Len()
		height++
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return height
}

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
