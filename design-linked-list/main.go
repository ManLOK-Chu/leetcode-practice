package main

type MyLinkedList struct {
	dummyHead *LinkedListNode
	size      int
}

type LinkedListNode struct {
	val  int
	next *LinkedListNode
}

func Constructor() MyLinkedList {
	return MyLinkedList{dummyHead: &LinkedListNode{}}
}

func (this *MyLinkedList) Get(index int) int {
	cur := this.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &LinkedListNode{val: val, next: this.dummyHead.next}
	this.dummyHead.next = node
	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	cur := this.dummyHead.next
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = &LinkedListNode{val: val}
	this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	cur := this.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	cur := this.dummyHead
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.next = cur.next.next
	this.size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func main() {
	myLinkedList := Constructor()
	myLinkedList.AddAtHead(1)
	myLinkedList.AddAtTail(3)
	myLinkedList.AddAtIndex(1, 2) // linked list becomes 1->2->3
	myLinkedList.Get(1)           // return 2
	myLinkedList.DeleteAtIndex(1) // now the linked list is 1->3
	myLinkedList.Get(1)           // return 3
}
