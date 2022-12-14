package main

import "fmt"

type MyQueue struct {
	list []int
	size int
}

func Constructor() MyQueue {
	var queue MyQueue
	queue.list = make([]int, 0, 32)
	return queue
}

func (this *MyQueue) Push(x int) {
	this.list = append(this.list, x)
	this.size++
}

func (this *MyQueue) Pop() int {
	val := this.Peek()
	if val != -1 {
		this.list = this.list[1:]
		this.size--
	}
	return val
}

func (this *MyQueue) Peek() int {
	if this.size > 0 {
		return this.list[0]
	}
	return -1
}

func (this *MyQueue) Empty() bool {
	return this.size == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
func main() {
	var myQueue = Constructor()
	myQueue.Push(1) // queue is: [1]
	fmt.Println(myQueue)
	myQueue.Push(2) // queue is: [1, 2] (leftmost is front of the queue)
	fmt.Println(myQueue)
	fmt.Println(myQueue.Peek()) // return 1
	fmt.Println(myQueue.Pop())  // return 1, queue is [2]
	fmt.Println(myQueue)
	fmt.Println(myQueue.Empty()) // return false

}

func (this MyQueue) String() string {
	return fmt.Sprint(this.list)
}
