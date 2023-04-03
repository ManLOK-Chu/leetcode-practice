package main

import (
	"fmt"
	"math/rand"
)

//快排
func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	mid := getPartitionIndex(nums, left, right)
	quickSort(nums, left, mid-1)
	quickSort(nums, mid+1, right)
}

func getPartitionIndex(nums []int, left, right int) int {
	cur := left + rand.Intn(right-left+1) // 随机选取"支点"
	nums[left], nums[cur] = nums[cur], nums[left]
	for left < right {
		for nums[left] < nums[right] && left < right { // 修改原来的 nums[right] >= nums[left]，增加交换频率
			right--
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		for nums[left] < nums[right] && left < right { // 修改原来的 nums[right] >= nums[left]，增加交换频率
			left++
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		}
	}
	return left
}

//冒泡排序
//func sortArray(nums []int) []int {
//	length := len(nums)
//	var flag bool
//	for i := 0; i < length-1; i++ { //一共要排序size-1次
//		flag = false                      //每次遍历标志位都要先置为false，才能判断后面的元素是否发生了交换
//		for j := 0; j < length-1-i; j++ { //选出该趟排序的最大值往后移动
//			if nums[j] > nums[j+1] {
//				//swap
//				nums[j+1], nums[j] = nums[j], nums[j+1]
//				flag = true //只要有发生了交换，flag就置为true
//			}
//		}
//		if !flag { //判断标志位是否为false，如果为false，说明后面的元素已经有序，提前结束
//			break
//		}
//	}
//	return nums
//}

func main() {
	fmt.Println(sortArray([]int{5, 2, 3, 1}))
	fmt.Println(sortArray([]int{5, 1, 1, 2, 0, 0}))
}
