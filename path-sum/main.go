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
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	var stack = list.New()
	var pathStack = list.New()
	stack.PushBack(root)
	pathStack.PushBack(root.Val)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		pathSum := pathStack.Remove(pathStack.Back()).(int)
		if node.Left == nil && node.Right == nil && pathSum == targetSum {
			return true
		}
		if node.Right != nil {
			stack.PushBack(node.Right)
			pathStack.PushBack(pathSum + node.Right.Val)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
			pathStack.PushBack(pathSum + node.Left.Val)
		}
	}
	return false
}

func main() {
	fmt.Println(hasPathSum(arrayToTreeNode([]int{5, 4, 8, 11, null, 13, 4, 7, 2, null, null, null, 1}), 22))
	fmt.Println(hasPathSum(arrayToTreeNode([]int{1, 2, 3}), 5))
	fmt.Println(hasPathSum(arrayToTreeNode([]int{}), 0))
	fmt.Println(hasPathSum(arrayToTreeNode([]int{1}), 1))
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
