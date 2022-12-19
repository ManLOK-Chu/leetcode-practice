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

func main() {
	var root = arrayToTreeNode([]int{5, 4, 6, 1, 2})
	fmt.Println(preorderTraversal(root))
	fmt.Println(inorderTraversal(root))
	fmt.Println(postorderTraversal(root))
}

func preorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return []int{}
	}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() != 0 {
		element := stack.Back()
		node, ok := element.Value.(*TreeNode)
		if ok {
			stack.Remove(element)
			if node.Right != nil { //右
				stack.PushBack(node.Right)
			}
			if node.Left != nil { //左
				stack.PushBack(node.Left)
			}
			stack.PushBack(node) //中
			stack.PushBack(nil)
		} else {
			stack.Remove(stack.Back())
			node = stack.Remove(stack.Back()).(*TreeNode)
			result = append(result, node.Val)
		}
	}
	return result
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return []int{}
	}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() != 0 {
		element := stack.Back()
		node, ok := element.Value.(*TreeNode)
		if ok {
			stack.Remove(element)  // 将该节点弹出，避免重复操作，下面再将右中左节点添加到栈中
			if node.Right != nil { // 添加右节点（空节点不入栈）
				stack.PushBack(node.Right)
			}
			stack.PushBack(node)  // 添加中节点
			stack.PushBack(nil)   // 中节点访问过，但是还没有处理，加入空节点做为标记。
			if node.Left != nil { // 添加左节点（空节点不入栈）
				stack.PushBack(node.Left)
			}
		} else { // 只有遇到空节点的时候，才将下一个节点放进结果集
			stack.Remove(stack.Back())                    // 将空节点弹出
			node = stack.Remove(stack.Back()).(*TreeNode) // 重新取出栈中元素
			result = append(result, node.Val)             // 加入到结果集
		}
	}
	return result
}

func postorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return []int{}
	}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() != 0 {
		element := stack.Back()
		node, ok := element.Value.(*TreeNode)
		if ok {
			stack.Remove(element)
			stack.PushBack(node) //中
			stack.PushBack(nil)
			if node.Right != nil { //右
				stack.PushBack(node.Right)
			}
			if node.Left != nil { //左
				stack.PushBack(node.Left)
			}
		} else {
			stack.Remove(stack.Back())
			node = stack.Remove(stack.Back()).(*TreeNode)
			result = append(result, node.Val)
		}
	}
	return result
}
