package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	list     *list.List    //双向链表
	maps     map[int]*Node //优化读取O(1)
	capacity int
	size     int
}

type Node struct {
	key  int
	val  int
	self *list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		list:     list.New(),
		maps:     make(map[int]*Node),
		capacity: capacity,
		size:     0,
	}
}

func (this *LRUCache) Get(key int) int {
	if val, exist := this.maps[key]; exist {
		this.update(key)
		return val.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, exist := this.maps[key]; exist {
		this.update(key)
		this.maps[key].val = value
		return
	}
	node := &Node{key: key, val: value}
	node.self = this.list.PushBack(node)
	this.maps[key] = node
	this.size++
	if this.size > this.capacity {
		node = this.list.Remove(this.list.Front()).(*Node)
		delete(this.maps, node.key)
	}
}

func (this *LRUCache) update(key int) {
	node := this.maps[key]
	this.list.Remove(node.self)
	node.self = this.list.PushBack(node)
	//this.list.MoveToBack(node.self)
	this.maps[key] = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	var lRUCache = Constructor(2)
	lRUCache.Put(1, 1) // 缓存是 {1=1}
	lRUCache.Print()
	lRUCache.Put(2, 2) // 缓存是 {1=1, 2=2}
	lRUCache.Print()
	lRUCache.Get(1)    // 返回 1
	lRUCache.Put(3, 3) // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	lRUCache.Print()
	lRUCache.Get(2)    // 返回 -1 (未找到)
	lRUCache.Put(4, 4) // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	lRUCache.Print()
	lRUCache.Get(1) // 返回 -1 (未找到)
	lRUCache.Get(3) // 返回 3
	lRUCache.Get(4) // 返回 4
}

func (this *LRUCache) Print() {
	cur := this.list.Front()
	for cur != nil {
		node := cur.Value.(*Node)
		fmt.Printf("%d=%d,", node.key, node.val)
		cur = cur.Next()
	}
	fmt.Println()
}
