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
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		element := stack.Back()
		node, ok := element.Value.(*TreeNode)
		if ok {
			stack.Remove(element)
			if node.Right != nil {
				stack.PushBack(node.Right)
			}
			if node.Left != nil {
				stack.PushBack(node.Left)
			}
			stack.PushBack(node)
			stack.PushBack(nil)
		} else {
			stack.Remove(element)
			node = stack.Remove(stack.Back()).(*TreeNode)
			node.Left, node.Right = node.Right, node.Left
		}
	}
	return root
}

//func invertTree(root *TreeNode) *TreeNode {
//	if root ==nil{
//		return nil
//	}
//	root.Left,root.Right=root.Right,root.Left//交换
//	invertTree(root.Left)
//	invertTree(root.Right)
//	return root
//}

func main() {
	fmt.Println(invertTree(arrayToTreeNode([]int{4, 2, 7, 1, 3, 6, 9})))
	fmt.Println(invertTree(arrayToTreeNode([]int{2, 1, 3})))
	fmt.Println(invertTree(arrayToTreeNode([]int{})))
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
