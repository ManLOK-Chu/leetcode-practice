package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		size := queue.Len()
		var pre *Node
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*Node)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			if pre != nil {
				pre.Next = node
			}
			pre = node
		}
	}
	return root
}

func main() {
	fmt.Println(connect(arrayToNode([]int{1, 2, 3, 4, 5, 6, 7})))
	fmt.Println(connect(arrayToNode([]int{})))
}

func arrayToNode(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}
	var root = &Node{Val: arr[0]}
	queue := list.New()
	queue.PushBack(root)
	var isLeft = true
	for i := 1; i < len(arr); i++ {
		element := queue.Front() //队头出队
		node := element.Value.(*Node)
		if isLeft {
			node.Left = &Node{Val: arr[i]}
			queue.PushBack(node.Left) //队尾进队
			isLeft = false
		} else {
			node.Right = &Node{Val: arr[i]}
			queue.PushBack(node.Right) //队尾进队
			queue.Remove(element)
			isLeft = true
		}
	}
	return root
}
