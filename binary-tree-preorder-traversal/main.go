package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var stack = list.New()
	if root == nil {
		return []int{}
	}
	var result = make([]int, 0, 32)
	stack.PushBack(root)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		result = append(result, node.Val)
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
	return result
}

func main() {

}
