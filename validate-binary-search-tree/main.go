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
func isValidBST(root *TreeNode) bool {
	stack := list.New()
	stack.PushBack(root)
	var preV int          //前值
	var flag bool         //开始比对的标志位
	for stack.Len() > 0 { //中序遍历
		element := stack.Back()
		node, ok := element.Value.(*TreeNode)
		if ok {
			stack.Remove(element)
			if node.Right != nil { //右
				stack.PushBack(node.Right)
			}
			stack.PushBack(node) //中
			stack.PushBack(nil)
			if node.Left != nil { //左
				stack.PushBack(node.Left)
			}
		} else {
			stack.Remove(element)
			node = stack.Remove(stack.Back()).(*TreeNode)
			if flag && preV >= node.Val {
				return false
			} else {
				preV = node.Val
				flag = true
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isValidBST(arrayToTreeNode([]int{2, 1, 3})))
	fmt.Println(isValidBST(arrayToTreeNode([]int{2, 2, 2})))
	fmt.Println(isValidBST(arrayToTreeNode([]int{5, 1, 4, null, null, 3, 6})))
	fmt.Println(isValidBST(arrayToTreeNode([]int{0, -1})))
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
