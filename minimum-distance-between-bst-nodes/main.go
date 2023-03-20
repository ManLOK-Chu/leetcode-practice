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
func minDiffInBST(root *TreeNode) int {
	stack := list.New()
	stack.PushBack(root)
	var min = math.MaxInt
	var flag bool //开始计算标志位
	var preV int  //前值
	for stack.Len() > 0 {
		node, ok := stack.Remove(stack.Back()).(*TreeNode)
		if ok {
			if node.Right != nil { //右
				stack.PushBack(node.Right)
			}
			stack.PushBack(node) //中
			stack.PushBack(nil)
			if node.Left != nil { //左
				stack.PushBack(node.Left)
			}
		} else {
			node = stack.Remove(stack.Back()).(*TreeNode)
			if flag && node.Val-preV < min {
				min = node.Val - preV
			} else {
				flag = true
			}
			preV = node.Val
		}
	}
	return min
}

func main() {
	fmt.Println(minDiffInBST(arrayToTreeNode([]int{4, 2, 6, 1, 3})))
	fmt.Println(minDiffInBST(arrayToTreeNode([]int{1, 0, 48, null, null, 12, 49})))
	fmt.Println(minDiffInBST(arrayToTreeNode([]int{90, 69, null, 49, 89, null, 52})))
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
