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
func buildTree(inorder []int, postorder []int) *TreeNode {
	var pLength = len(postorder)
	if pLength == 0 {
		return nil
	}
	val := postorder[pLength-1]
	root := &TreeNode{Val: val}
	if pLength == 1 {
		return root
	}
	rootIndex := 0
	for ; rootIndex < pLength; rootIndex++ { //inorder数组与postorder等长
		if inorder[rootIndex] == val {
			break
		}
	}
	var inLeft, inRight = inorder[:rootIndex], inorder[rootIndex+1:]
	var postLeft, postRight = postorder[:len(inLeft)], postorder[len(inLeft) : pLength-1]
	root.Left = buildTree(inLeft, postLeft)
	root.Right = buildTree(inRight, postRight)
	return root
}

func main() {
	fmt.Println(buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}))
	fmt.Println(buildTree([]int{-1}, []int{-1}))
}
