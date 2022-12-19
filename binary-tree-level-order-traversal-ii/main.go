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
func levelOrderBottom(root *TreeNode) [][]int {
	var results [][]int
	if root == nil {
		return [][]int{}
	}
	queue := list.New()
	var cur = root
	queue.PushBack(cur)
	var length int
	for queue.Len() > 0 {
		size := queue.Len()
		values := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			values[i] = node.Val
		}
		results = append(results, values)
		length++
	}
	for i := 0; i < length/2; i++ {
		results[length-1-i], results[i] = results[i], results[length-1-i]
	}
	return results
}

func main() {
	fmt.Println(levelOrderBottom(arrayToTreeNode([]int{3, 9, 20, null, null, 15, 7})))
	fmt.Println(levelOrderBottom(arrayToTreeNode([]int{1})))
	fmt.Println(levelOrderBottom(arrayToTreeNode([]int{})))
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
