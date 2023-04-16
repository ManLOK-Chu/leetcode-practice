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

//动态规划
func rob(root *TreeNode) int {
	var robTree func(*TreeNode) [2]int
	robTree = func(root *TreeNode) [2]int { //返回长度为2的数组，0：不偷，1：偷
		if root == nil {
			return [2]int{0, 0}
		}
		left := robTree(root.Left)
		right := robTree(root.Right)
		// 偷root，那么就不能偷左右节点
		val1 := root.Val + left[0] + right[0]
		// 不偷root，那么可以偷也可以不偷左右节点，则取较大的情况
		val2 := max(left[0], left[1]) + max(right[0], right[1])
		return [2]int{val2, val1}
	}
	var result = robTree(root)
	return max(result[0], result[1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//记忆化递推 + 使用哈希表
//时间复杂度O（n)
//空间复杂度O(log n)+递归栈空间
//func rob(root *TreeNode) int {
//	var cached = make(map[*TreeNode]int)
//	var robRange func(*TreeNode) int
//	robRange = func(root *TreeNode) int {
//		if root == nil {
//			return 0
//		}
//		if v, ok := cached[root]; ok { //命中缓存，提前返回
//			return v
//		}
//		if root.Left == nil && root.Right == nil {
//			return root.Val
//		}
//		val1 := root.Val //偷父节点
//		if root.Left != nil {
//			val1 += robRange(root.Left.Left) + robRange(root.Left.Right) //跳过左子节点
//		}
//		if root.Right != nil {
//			val1 += robRange(root.Right.Left) + robRange(root.Right.Right) //跳过右子节点
//		}
//
//		val2 := 0 //不偷父节点，偷子节点
//		val2 = robRange(root.Left) + robRange(root.Right)
//		if val2 > val1 { //求max
//			cached[root] = val2
//			return val2
//		}
//		cached[root] = val1
//		return val1
//	}
//	return robRange(root)
//}

func main() {
	fmt.Println(rob(arrayToTreeNode([]int{3, 2, 3, null, 3, null, 1})))
	fmt.Println(rob(arrayToTreeNode([]int{3, 4, 5, 1, 3, null, 1})))
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
