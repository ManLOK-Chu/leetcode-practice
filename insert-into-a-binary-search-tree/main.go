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
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	var cur = root
	var pre *TreeNode
	for cur != nil {
		pre = cur
		if val > cur.Val {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}
	if val > pre.Val {
		pre.Right = &TreeNode{Val: val}
	}
	if val < pre.Val {
		pre.Left = &TreeNode{Val: val}
	}
	return root
}

func main() {
	fmt.Println(insertIntoBST(arrayToTreeNode([]int{4, 2, 7, 1, 3}), 5))
	fmt.Println(insertIntoBST(arrayToTreeNode([]int{40, 20, 60, 10, 30, 50, 70}), 25))
	fmt.Println(insertIntoBST(arrayToTreeNode([]int{4, 2, 7, 1, 3, null, null, null, null, null, null}), 5))
}

func arrayToTreeNode(arr []int) *TreeNode {
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
