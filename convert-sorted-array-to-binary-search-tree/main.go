package main

import (
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
func sortedArrayToBST(nums []int) *TreeNode {
	var length = len(nums)
	if length == 0 {
		return nil
	}
	index := length / 2
	rootVal := nums[index]
	var left, right []int
	left = nums[:index]
	right = nums[index+1:]
	var root = &TreeNode{Val: rootVal}
	root.Left = sortedArrayToBST(left)
	root.Right = sortedArrayToBST(right)
	return root

}

func main() {
	fmt.Println(sortedArrayToBST([]int{-10, -3, 0, 5, 9}))
	fmt.Println(sortedArrayToBST([]int{1, 3}))
}
