package main

import (
	"container/list"
	"fmt"
	"math"
	"strconv"
)

const null = math.MinInt32

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) String() string {
	return strconv.Itoa(n.Val)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	var stack = list.New()
	var pathStack = list.New()
	stack.PushBack(root)
	pathStack.PushBack([]int{root.Val})
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		paths := pathStack.Remove(pathStack.Back()).([]int)
		if node.Left == nil && node.Right == nil {
			if sum(paths) == targetSum {
				result = append(result, paths)
			}
		}
		if node.Right != nil {
			stack.PushBack(node.Right)
			var tmp = make([]int, len(paths))
			copy(tmp, paths)
			pathStack.PushBack(append(tmp, node.Right.Val))
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
			var tmp = make([]int, len(paths))
			copy(tmp, paths)
			pathStack.PushBack(append(tmp, node.Left.Val))
		}
	}
	return result
}

func sum(a []int) int {
	var s int
	for i := range a {
		s += a[i]
	}
	return s
}

func equals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(pathSum(arrayToTreeNode([]int{5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1}), 22))
	fmt.Println(pathSum(arrayToTreeNode([]int{1, 2, 3}), 5))
	fmt.Println(pathSum(arrayToTreeNode([]int{1, 2}), 0))
	fmt.Println(pathSum(arrayToTreeNode([]int{}), 0))
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
