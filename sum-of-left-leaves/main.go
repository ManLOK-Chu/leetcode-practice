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
func sumOfLeftLeaves(root *TreeNode) int {
	var sum int
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode) //取出节点，中
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			sum += node.Left.Val
		}
		if node.Right != nil { //右
			stack.PushBack(node.Right)
		}
		if node.Left != nil { //左
			stack.PushBack(node.Left)
		}
	}
	return sum
}

func main() {
	fmt.Println(sumOfLeftLeaves(arrayToTreeNode([]int{3, 9, 20, null, null, 15, 7})))
	fmt.Println(sumOfLeftLeaves(arrayToTreeNode([]int{1})))
	fmt.Println(sumOfLeftLeaves(arrayToTreeNode([]int{1, 2, 3, 4, 5})))
	fmt.Println(sumOfLeftLeaves(arrayToTreeNode([]int{0, 2, 4, 1, null, 3, -1, 5, 1, null, 6, null, 8})))
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
