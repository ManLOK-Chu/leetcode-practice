package main

import "fmt"

//需要用一个变量 cnt 来记录连续1的个数吧，那么当遇到了0的时候怎么处理呢，因为有一次0变1的机会，所以遇到0了还是要累加 cnt，
//然后此时需要用另外一个变量 cur 来保存当前 cnt 的值，然后 cnt 重置为0，以便于让 cnt 一直用来统计纯连续1的个数，然后我们每次都用用 cnt+cur 来更新结果 max
func findMaxConsecutiveOnes(nums []int) int {
	var max, count, lastCount int
	for i := 0; i < len(nums); i++ {
		count++
		if nums[i] == 0 {
			lastCount = count
			count = 0
		}
		if max < lastCount+count {
			max = lastCount + count
		}
	}
	return max
}

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0}))
}
