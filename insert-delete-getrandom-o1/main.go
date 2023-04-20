package main

import (
	"fmt"
	"math/rand"
	"time"
)

type RandomizedSet struct {
	nums   []int
	length int
	localM map[int]int
}

func Constructor() RandomizedSet {
	rand.Seed(time.Now().UnixNano())
	return RandomizedSet{
		nums:   []int{},
		length: 0,
		localM: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.localM[val]; ok {
		return false
	}
	this.localM[val] = this.length
	this.nums = append(this.nums, val)
	this.length++
	return true

}

func (this *RandomizedSet) Remove(val int) bool {
	id, ok := this.localM[val]
	if !ok {
		return false
	}
	this.length--
	this.nums[id] = this.nums[this.length]
	this.localM[this.nums[id]] = id
	this.nums = this.nums[:this.length]
	delete(this.localM, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

func main() {
	randomizedSet := Constructor()
	//fmt.Println(randomizedSet.Insert(10))  // 向集合中插入 1 。返回 true 表示 1 被成功地插入。
	//fmt.Println(randomizedSet.Remove(20))  // 返回 false ，表示集合中不存在 2 。
	//fmt.Println(randomizedSet.Insert(20))  // 向集合中插入 2 。返回 true 。集合现在包含 [1,2] 。
	//fmt.Println(randomizedSet.GetRandom()) // getRandom 应随机返回 1 或 2 。
	//fmt.Println(randomizedSet.Remove(10))  // 从集合中移除 1 ，返回 true 。集合现在包含 [2] 。
	//fmt.Println(randomizedSet.Insert(20))  // 2 已在集合中，所以返回 false 。
	//fmt.Println(randomizedSet.GetRandom()) // 由于 2 是集合中唯一的数字，getRandom 总是返回 2 。
	fmt.Println(randomizedSet.Insert(0)) //{0}
	fmt.Println(randomizedSet.Insert(1)) //{0,1}
	fmt.Println(randomizedSet)
	fmt.Println(randomizedSet.Remove(0)) //{1}
	fmt.Println(randomizedSet)
	fmt.Println(randomizedSet.Insert(2)) //{1,2}
	fmt.Println(randomizedSet)
	fmt.Println(randomizedSet.Remove(1)) //{2}
	fmt.Println(randomizedSet)
	fmt.Println(randomizedSet.GetRandom())
}
