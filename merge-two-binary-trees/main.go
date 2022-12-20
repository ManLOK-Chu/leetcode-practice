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
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	queue := list.New()
	queue.PushBack(root1)
	queue.PushBack(root2)
	var root = &TreeNode{}
	queue.PushBack(root)
	for queue.Len() > 0 {
		node1 := queue.Remove(queue.Front()).(*TreeNode)
		node2 := queue.Remove(queue.Front()).(*TreeNode)
		cur := queue.Remove(queue.Front()).(*TreeNode)
		cur.Val = node1.Val + node2.Val
		if node1.Left != nil && node2.Left != nil {
			queue.PushBack(node1.Left)
			queue.PushBack(node2.Left)
			cur.Left = &TreeNode{}
			queue.PushBack(cur.Left)
		}
		if node1.Right != nil && node2.Right != nil {
			queue.PushBack(node1.Right)
			queue.PushBack(node2.Right)
			cur.Right = &TreeNode{}
			queue.PushBack(cur.Right)
		}
		if node1.Left != nil && node2.Left == nil {
			cur.Left = node1.Left
		}
		if node1.Left == nil && node2.Left != nil {
			cur.Left = node2.Left
		}
		if node1.Right != nil && node2.Right == nil {
			cur.Right = node1.Right
		}
		if node1.Right == nil && node2.Right != nil {
			cur.Right = node2.Right
		}
	}
	return root
}

//func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
//	if root1 == nil {
//		return root2
//	}
//	if root2 == nil {
//		return root1
//	}
//	root := &TreeNode{Val: root1.Val + root2.Val}
//	root.Left = mergeTrees(root1.Left, root2.Left)
//	root.Right = mergeTrees(root1.Right, root2.Right)
//	return root
//}

func main() {
	fmt.Println(mergeTrees(arrayToTreeNode([]int{1, 3, 2, 5}), arrayToTreeNode([]int{2, 1, 3, null, 4, null, 7})))
	fmt.Println(mergeTrees(arrayToTreeNode([]int{1}), arrayToTreeNode([]int{1, 2})))
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
