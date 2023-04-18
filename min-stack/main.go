package main

// MinStack 最小栈
//设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
//MinStack() 初始化堆栈对象。
//void push(int val) 将元素val推入堆栈。
//void pop() 删除堆栈顶部的元素。
//int top() 获取堆栈顶部的元素。
//int getMin() 获取堆栈中的最小元素。
type MinStack struct {
	stack    []int
	minStack []int
	size     int
}

func Constructor() MinStack {
	return MinStack{
		stack:    make([]int, 0, 16),
		minStack: make([]int, 0, 16),
		size:     0,
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if this.size > 0 {
		top := this.minStack[this.size-1]
		if top < val {
			val = top
		}
	}
	this.minStack = append(this.minStack, val)
	this.size++
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:this.size-1]
	this.minStack = this.minStack[:this.size-1]
	this.size--
}

func (this *MinStack) Top() int {
	return this.stack[this.size-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[this.size-1]
}
