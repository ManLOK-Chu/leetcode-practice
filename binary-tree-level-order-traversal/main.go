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
//func levelOrder(root *TreeNode) [][]int {
//	if root == nil {
//		return [][]int{}
//	}
//	var result [][]int
//	var queue = list.New()
//	var cur = root
//	queue.PushBack(cur)
//	for queue.Len() != 0 {
//		size := queue.Len()
//		values := make([]int, size)
//		for i := 0; i < size; i++ {
//			node := queue.Remove(queue.Front()).(*TreeNode)
//			values[i] = node.Val
//			if node.Left != nil {
//				queue.PushBack(node.Left)
//			}
//			if node.Right != nil {
//				queue.PushBack(node.Right)
//			}
//		}
//		result = append(result, values)
//	}
//	return result
//}
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var result [][]int
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		size := queue.Len()
		values := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			values[i] = node.Val
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		result = append(result, values)
	}
	return result
}

func main() {
	fmt.Println(levelOrder(arrayToTreeNode([]int{3, 9, 20, null, null, 15, 7})))
	fmt.Println(levelOrder(arrayToTreeNode([]int{1})))
	fmt.Println(levelOrder(arrayToTreeNode([]int{})))
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
