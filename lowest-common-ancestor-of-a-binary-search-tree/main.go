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
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var cur = root
	for cur != nil {
		if cur.Val > p.Val && cur.Val > q.Val {
			cur = cur.Left
		} else if cur.Val < p.Val && cur.Val < q.Val {
			cur = cur.Right
		} else {
			return cur
		}
	}
	return nil
}

func main() {
	fmt.Println(lowestCommonAncestor(arrayToTreeNode([]int{6, 2, 8, 0, 4, 7, 9, null, null, 3, 5}, 2, 8)))
	fmt.Println(lowestCommonAncestor(arrayToTreeNode([]int{6, 2, 8, 0, 4, 7, 9, null, null, 3, 5}, 2, 4)))
	fmt.Println(lowestCommonAncestor(arrayToTreeNode([]int{2, 1}, 2, 1)))
}

func arrayToTreeNode(arr []int, pVal, qVal int) (*TreeNode, *TreeNode, *TreeNode) {
	var root, p, q *TreeNode
	root = &TreeNode{Val: arr[0]}
	queue := list.New()
	queue.PushBack(root)
	var isLeft = true
	if arr[0] == pVal {
		p = root
	} else if arr[0] == qVal {
		q = root
	}
	for i := 1; i < len(arr); i++ {
		element := queue.Front() //队头出队
		node := element.Value.(*TreeNode)
		tmp := &TreeNode{Val: arr[i]}
		if isLeft {
			if arr[i] != null {
				node.Left = tmp
				queue.PushBack(node.Left) //队尾进队
			}
			isLeft = false
		} else {
			if arr[i] != null {
				node.Right = tmp
				queue.PushBack(node.Right) //队尾进队
			}
			queue.Remove(element)
			isLeft = true
		}
		if arr[i] == pVal {
			p = tmp
		} else if arr[i] == qVal {
			q = tmp
		}
	}
	return root, p, q
}
