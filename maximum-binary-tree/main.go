package main

import "fmt"

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
func constructMaximumBinaryTree(nums []int) *TreeNode {
	var length = len(nums)
	if length == 0 {
		return nil
	}
	var max, index int
	for i := 0; i < length; i++ {
		if nums[i] > max {
			index = i
			max = nums[i]
		}
	}
	root := &TreeNode{Val: max}
	if length == 1 {
		return root
	}
	root.Left = constructMaximumBinaryTree(nums[:index])
	root.Right = constructMaximumBinaryTree(nums[index+1:])
	return root
}

func main() {
	fmt.Println(constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5}))
	fmt.Println(constructMaximumBinaryTree([]int{3, 2, 1}))
}
