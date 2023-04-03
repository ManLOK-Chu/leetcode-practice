package main

import (
	"fmt"
	"math"
)

//func maxProfit(prices []int) int {
//	//min := math.MaxInt
//	//var res int
//	//for i := 0; i < len(prices); i++ {
//	//	if prices[i] < min {
//	//		min = prices[i]
//	//	}
//	//	if prices[i]-min > res {
//	//		res = prices[i] - min
//	//	}
//	//}
//	//return res
//	var max int
//	var slowIndex, fastIndex, length = 0, 1, len(prices)
//	for slowIndex < length-1 {
//		if gap := prices[fastIndex] -
//			prices[slowIndex]; gap > max {
//			max = gap
//		}
//		fastIndex++
//		if fastIndex >= length {
//			slowIndex++
//			fastIndex = slowIndex + 1
//		}
//	}
//	return max
//}
func maxProfit(prices []int) int {
	buy := prices[0]
	sell := prices[0]
	max := math.MinInt
	for _, v := range prices {
		if buy > v {
			buy = v
			sell = buy
		} else if buy < v {
			sell = v
			gap := sell - buy
			if gap > max {
				max = gap
			}
		}
	}
	if max > 0 {
		return max
	}
	return 0
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit([]int{0}))
	fmt.Println(maxProfit([]int{3, 2, 6, 5, 0, 3}))
}
