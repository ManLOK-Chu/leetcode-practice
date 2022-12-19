package main

import (
	"container/list"
	"fmt"
	"math"
)

type Node struct {
	Val      int
	Children []*Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	var results [][]int
	if root == nil {
		return results
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		size := queue.Len()
		values := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*Node)
			if len(node.Children) != 0 {
				for j := range node.Children {
					queue.PushBack(node.Children[j])
				}
			}
			values[i] = node.Val
		}
		results = append(results, values)
	}
	return results
}

func main() {
	fmt.Println(levelOrder(arrayToNode([]int{1, null, 3, 2, 4, null, 5, 6})))
	fmt.Println(levelOrder(arrayToNode([]int{1, null, 2, 3, 4, 5, null, null, 6, 7, null, 8, null, 9, 10, null, null, 11, null, 12, null, 13, null, null, 14})))
	fmt.Println(levelOrder(arrayToNode([]int{})))
}

const null = math.MinInt32

func arrayToNode(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}
	var root = &Node{Val: arr[0]}
	var midNode = root
	var nodeList []*Node
	var listIndex int
	for i := 1; i < len(arr); i++ {
		if arr[i] == null {
			if listIndex+1 < len(nodeList) {
				midNode = nodeList[listIndex]
				listIndex++
			} else {
				nodeList = nil
			}
			continue
		}
		node := &Node{Val: arr[i]}
		nodeList = append(nodeList, node)
		midNode.Children = append(midNode.Children, node)
	}
	return root
}
