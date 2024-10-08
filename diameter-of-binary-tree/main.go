package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans int

func diameterOfBinaryTree(root *TreeNode) int {
	ans = 1
	depth(root)
	return ans - 1
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := depth(node.Left)
	right := depth(node.Right)
	ans = max(ans, right+left+1)
	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
