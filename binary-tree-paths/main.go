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

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	var result []string
	treeStack := list.New()
	pathStack := list.New()
	treeStack.PushBack(root)
	pathStack.PushBack(strconv.Itoa(root.Val))
	for treeStack.Len() > 0 {
		node := treeStack.Remove(treeStack.Back()).(*TreeNode)
		paths := pathStack.Remove(pathStack.Back()).(string)
		if node.Left == nil && node.Right == nil {
			result = append(result, paths)
		}
		if node.Right != nil {
			treeStack.PushBack(node.Right)
			pathStack.PushBack(paths + "->" + strconv.Itoa(node.Right.Val))
		}
		if node.Left != nil {
			treeStack.PushBack(node.Left)
			pathStack.PushBack(paths + "->" + strconv.Itoa(node.Left.Val))
		}
	}
	return result
}

func main() {
	fmt.Println(binaryTreePaths(arrayToTreeNode([]int{1, 2, 3, null, 5})))
	fmt.Println(binaryTreePaths(arrayToTreeNode([]int{1})))
	fmt.Println(binaryTreePaths(arrayToTreeNode([]int{6, 1, null, null, 3, 2, 5, null, null, 4})))
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
