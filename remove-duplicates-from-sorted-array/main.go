package main

import "fmt"

func removeDuplicates(nums []int) int {
	length := len(nums)
	var count int
	for i := 0; i < length-1; i++ {
		if nums[i] == nums[i+1] {
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
			if nums[j+1] < nums[j] {
				changed = true
				nums[j+1], nums[j] = nums[j], nums[j+1]
			}
		}
		if !changed {
			break
		}
	}
}

func main() {
	var input = []int{98, 99, 99, 100}
	fmt.Println(removeDuplicates(input))
	fmt.Println(input)
}
