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
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := list.New()
	stack.PushBack(root)
	var preV int
	for stack.Len() > 0 { //反中序遍历
		node, ok := stack.Remove(stack.Back()).(*TreeNode)
		if ok {
			if node.Left != nil {
				stack.PushBack(node.Left)
			}
			stack.PushBack(node)
			stack.PushBack(nil)
			if node.Right != nil {
				stack.PushBack(node.Right)
			}
		} else {
			node = stack.Remove(stack.Back()).(*TreeNode)
			fmt.Println(node.Val)
			node.Val += preV
			preV = node.Val
		}
	}
	return root
}

func main() {
	fmt.Println(convertBST(arrayToTreeNode([]int{4, 1, 6, 0, 2, 5, 7, null, null, null, 3, null, null, null, 8})))
	fmt.Println(convertBST(arrayToTreeNode([]int{0, null, 1})))
	fmt.Println(convertBST(arrayToTreeNode([]int{})))
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
		tmp := &TreeNode{Val: arr[i]}
		if isLeft {
			if arr[i] != null {
				node.Left = tmp
				queue.PushBack(node.Left) //队尾进队
			}
			isLeft = false
		} else {
			if arr[i] != null {
				node.Right = tmp
				queue.PushBack(node.Right) //队尾进队
			}
			queue.Remove(element)
			isLeft = true
		}
	}
	return root
}
