package main

import "container/list"

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
func inorderTraversal(root *TreeNode) []int {
	var stack = list.New()
	if root == nil {
		return []int{}
	}
	var result = make([]int, 0, 32)
	var cur = root
	for cur != nil || stack.Len() > 0 {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
			continue
		}
		node := stack.Remove(stack.Back()).(*TreeNode)
		result = append(result, node.Val)
		cur = node.Right
	}
	return result
}

//func inorderTraversal(root *TreeNode) []int {
//	if root == nil {
//		return nil
//	}
//	var result []int
//	if leftRes := inorderTraversal(root.Left); leftRes != nil {
//		result = append(result, leftRes...)
//	}
//	result = append(result, root.Val)
//	if rightRes := inorderTraversal(root.Right); rightRes != nil {
//		result = append(result, rightRes...)
//	}
//	return result
//}

func main() {

}
