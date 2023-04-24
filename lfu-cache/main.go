package main

import "fmt"

// 请你为 最不经常使用（LFU）缓存算法设计并实现数据结构
// 双哈希表
type LFUCache struct {
	cache    map[int]*LFUCacheNode
	freqMap  map[int]*LFUCacheList
	minFreq  int // 最小频次
	capacity int
}

type LFUCacheNode struct {
	pre   *LFUCacheNode
	next  *LFUCacheNode
	freq  int //频次
	key   int
	value int
}

type LFUCacheList struct {
	head   *LFUCacheNode //虚拟头节点
	tail   *LFUCacheNode //虚拟尾节点
	length int           //链表长度
}

func newLFUCacheList() *LFUCacheList {
	list := &LFUCacheList{
		head: &LFUCacheNode{},
		tail: &LFUCacheNode{},
	}
	list.head.next = list.tail
	list.tail.pre = list.head
	return list
}

func (list *LFUCacheList) removeNode(node *LFUCacheNode) bool {
	node.remove()
	list.length--
	// 查询长度，若小于等于0，即该链表没有元素，通知删除该链表
	if list.length <= 0 {
		return true
	}
	return false
}

// 将节点添加到双向链表前面
func (list *LFUCacheList) addToHead(node *LFUCacheNode) {
	node.moveToHead(list.head)
	list.length++
}

// 获取链表的最后一个节点
func (list *LFUCacheList) getLast() *LFUCacheNode {
	return list.tail.pre
}

//移除节点
func (node *LFUCacheNode) remove() {
	node.pre.next = node.next
	node.next.pre = node.pre
	node.pre = nil
	node.next = nil
}

// 插入最前面
func (n *LFUCacheNode) moveToHead(head *LFUCacheNode) {
	head.next.pre = n
	n.next = head.next
	n.pre = head
	head.next = n
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cache:    make(map[int]*LFUCacheNode, capacity),
		freqMap:  make(map[int]*LFUCacheList),
		capacity: capacity,
		minFreq:  0,
	}
}

func (this *LFUCache) Get(key int) int {
	if this.capacity < 1 {
		return -1
	}
	// 查询节点
	node, ok := this.cache[key]
	// 不存在直接返回-1
	if !ok {
		return -1
	}
	this.updateNode(node) // 更新频次链表
	return node.value
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity < 1 {
		return
	}
	// 查询节点
	node, ok := this.cache[key]
	// 节点存在，更新节点
	if ok {
		node.value = value
		this.updateNode(node) // 更新频次链表
		return
	}
	//不存在就创建
	node = &LFUCacheNode{
		key:   key,
		value: value,
	}
	//当前长度小于设计容量时，直接插入并更新节点
	if len(this.cache) < this.capacity {
		this.cache[key] = node
		this.updateNode(node)
		return
	}
	//长度大于等于容量时，删除最不常用的节点
	list := this.freqMap[this.minFreq]
	lastNode := list.getLast()
	isDelete := list.removeNode(lastNode)
	if isDelete {
		delete(this.freqMap, this.minFreq)
	}
	delete(this.cache, lastNode.key)
	//插入
	this.cache[key] = node
	this.updateNode(node)
	return
}

func (this *LFUCache) updateNode(node *LFUCacheNode) {
	//获取原频次链表
	list, ok := this.freqMap[node.freq]
	if !ok {
		// 原频次不存在，说明是新节点，频次必定为最小，更新频次
		this.minFreq = node.freq + 1
	} else {
		//原频次存在，说明是旧节点，需要删除
		isDelete := list.removeNode(node)
		if isDelete {
			delete(this.freqMap, node.freq)
			// 当需要删除时，说明该频次已无节点，如果最小频次为该节点，则需更新最小频次
			if this.minFreq >= node.freq {
				this.minFreq = node.freq + 1
			}
		}
	}
	//更新频次
	node.freq++
	// 查询新频次链表
	newList, exist := this.freqMap[node.freq]
	if !exist { // 若不存在，创建链表并插入
		newList = newLFUCacheList()
		this.freqMap[node.freq] = newList
	}
	newList.addToHead(node)
	return
}

func main() {
	lfu := Constructor(2)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	fmt.Println(lfu.Get(1))
	lfu.Put(3, 3)
	fmt.Println(lfu.Get(2))
	fmt.Println(lfu.Get(3))
	lfu.Put(4, 4)
	fmt.Println(lfu.Get(1))
	fmt.Println(lfu.Get(3))
	fmt.Println(lfu.Get(3))
}
