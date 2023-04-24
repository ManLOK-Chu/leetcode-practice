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

func (node *TreeNode) Values() (results []int) {
	var cur = node
	var n int
	for n == 0 || cur != node {
		results = append(results, cur.Val)
		cur = cur.Right
		n++
	}
	return
}

//将一个二叉搜索树就地转化为一个已排序的双向循环链表。可以将左右孩子指针作为双向循环链表的前驱和后继指针
// 这道题目实际上也是将二叉搜索树序列化。注意转换规则：左指针指向前继节点，右指针指向后继节点。实际上就是中序遍历。
// 使用栈来帮助我们进行中序遍历。在出栈时改变它的左指针，指向，前继节点prev；prev的右指针指向当前出栈的节点。
// 需要注意的是：在最后需要连接头指针和尾指针
func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var cur = root
	stack := list.New()
	var head, pre *TreeNode = nil, nil
	for cur != nil || stack.Len() > 0 {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		} else {
			cur = stack.Remove(stack.Back()).(*TreeNode)
			if head == nil {
				head = cur
			}
			if pre != nil {
				pre.Right = cur
				cur.Left = pre
			}
			pre = cur
			cur = cur.Right
		}
	}
	head.Left = pre
	pre.Right = head
	return head
}

func main() {
	fmt.Println(treeToDoublyList(arrayToTreeNode([]int{4, 2, 5, 1, 3})).Values())
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
	}
	return root
}
