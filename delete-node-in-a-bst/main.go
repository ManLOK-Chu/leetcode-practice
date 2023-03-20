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
func deleteNode(root *TreeNode, key int) *TreeNode {
	// 特殊情况处理
	if root == nil {
		return root
	}
	cur := root
	var pre *TreeNode
	for cur != nil {
		if cur.Val == key {
			break
		}
		pre = cur
		if cur.Val > key {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	if pre == nil {
		return deleteOneNode(cur)
	}
	// pre 要知道是删除左孩子还有右孩子
	if pre.Left != nil && pre.Left.Val == key {
		pre.Left = deleteOneNode(cur)
	}
	if pre.Right != nil && pre.Right.Val == key {
		pre.Right = deleteOneNode(cur)
	}
	return root
}

func deleteOneNode(target *TreeNode) *TreeNode {
	if target == nil {
		return target
	}
	if target.Right == nil {
		return target.Left
	}
	cur := target.Right
	for cur.Left != nil {
		cur = cur.Left
	}
	cur.Left = target.Left
	return target.Right
}

func main() {
	fmt.Println(deleteNode(arrayToTreeNode([]int{5, 3, 6, 2, 4, null, 7}), 3))
	fmt.Println(deleteNode(arrayToTreeNode([]int{5, 3, 6, 2, 4, null, 7}), 0))
	fmt.Println(deleteNode(arrayToTreeNode([]int{}), 0))
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
