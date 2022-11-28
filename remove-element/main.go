package main

import "fmt"

/*
func removeElement(nums []int, val int) int {
	length := len(nums)
	var count int
	for i := 0; i < length; i++ {
		if nums[i] == val {
			nums[i] = 101
			count++
		}
	}
	bubbleSort(nums)
	return length - count
}

func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		var changed bool
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j+1] < nums[j] && nums[j] == 101 {
				changed = true
				nums[j+1], nums[j] = nums[j], nums[j+1]
			}
		}
		if !changed {
			break
		}
	}
}*/

func removeElement(nums []int, val int) int { //双指针
	length := len(nums)
	var slowIndex, fastIndex int
	for ; fastIndex < length; fastIndex++ {
		if nums[fastIndex] != val {
			nums[slowIndex] = nums[fastIndex]
			slowIndex++
		}
	}
	return slowIndex
}

func main() {
	var input = []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(input, 2))
	fmt.Println(input)
}
