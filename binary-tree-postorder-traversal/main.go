package main

import (
	"container/list"
	"fmt"
	"math"
)

const NULL = math.MinInt

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var stack = list.New()
	if root == nil {
		return []int{}
	}
	var result = make([]int, 0, 32)
	stack.PushBack(root)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		result = append(result, node.Val)
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
	}
	length := len(result)
	for i := 0; i < length/2; i++ {
		result[i], result[length-1-i] = result[length-1-i], result[i]
	}
	return result
}

func main() {
	fmt.Println(postorderTraversal(arrayToTreeNode([]int{1, NULL, 2, 3})))
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
			if arr[i] != NULL {
				node.Left = &TreeNode{Val: arr[i]}
				queue.PushBack(node.Left) //队尾进队
			}
			isLeft = false
		} else {
			if arr[i] != NULL {
				node.Right = &TreeNode{Val: arr[i]}
				queue.PushBack(node.Right) //队尾进队
			}
			queue.Remove(element)
			isLeft = true
		}
	}
	return root
}
