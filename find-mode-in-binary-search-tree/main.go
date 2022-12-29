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
func findMode(root *TreeNode) []int {
	stack := list.New()
	stack.PushBack(root)
	var result []int
	var count, maxCount = 0, 0 // 统计频率，最大频率
	var preV int               //前值
	var flag bool              //开始标志位
	for stack.Len() > 0 {
		node, ok := stack.Remove(stack.Back()).(*TreeNode)
		if ok {
			if node.Right != nil {
				stack.PushBack(node.Right)
			}
			stack.PushBack(node)
			stack.PushBack(nil)
			if node.Left != nil {
				stack.PushBack(node.Left)
			}
		} else {
			node = stack.Remove(stack.Back()).(*TreeNode)
			if flag && preV == node.Val { // 与前值相同
				count++
			} else { // 第一个节点或与前值不同
				count = 1
			}
			if count == maxCount { // 如果和最大值相同，放进result中
				result = append(result, node.Val)
			}
			if count > maxCount { // 如果计数大于最大频率
				maxCount = count         // 更新最大频率
				result = []int{node.Val} //重新生成当前结果集
			}
			preV = node.Val
			flag = true
		}
	}
	return result
}

func main() {
	fmt.Println(findMode(arrayToTreeNode([]int{1, null, 2, 2})))
	fmt.Println(findMode(arrayToTreeNode([]int{0})))
	fmt.Println(findMode(arrayToTreeNode([]int{1, null, 2})))
	fmt.Println(findMode(arrayToTreeNode([]int{1, 0, 1, 0, 0, 1, 1, 0})))
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
