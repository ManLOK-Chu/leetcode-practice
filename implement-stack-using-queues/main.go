package main

import "fmt"

type MyStack struct {
	list []int
	size int
}

func Constructor() MyStack {
	return MyStack{
		list: make([]int, 0, 32),
	}
}

func (this *MyStack) Push(x int) {
	this.list = append(this.list, x)
	this.size++
}

func (this *MyStack) Pop() int {
	val := this.Top()
	if val != -1 {
		this.size--
		this.list = this.list[:this.size]
	}
	return val
}

func (this *MyStack) Top() int {
	if this.size == 0 {
		return -1
	}
	return this.list[this.size-1]
}

func (this *MyStack) Empty() bool {
	return this.size == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
func main() {
	var myStack = Constructor()
	myStack.Push(1)
	fmt.Println(myStack)
	myStack.Push(2)
	fmt.Println(myStack)
	fmt.Println(myStack.Top()) // return 2
	fmt.Println(myStack.Pop()) // return 2
	fmt.Println(myStack)
	fmt.Println(myStack.Empty()) // return False
}

func (this MyStack) String() string {
	return fmt.Sprint(this.list[:this.size])
}
