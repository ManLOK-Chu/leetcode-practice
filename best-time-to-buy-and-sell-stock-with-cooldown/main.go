package main

import "fmt"

/* 每天都有两种可能性，要么卖，要么啥都不干。 当然这天的最大收益则是这两个方案中大的那个。 按照DP的思想，不如开两个数组，一个表示今天要卖，另一个表示今天啥也不干。
* profitDoNothing[i]比较容易想，因为今天啥也不干，所以今天就继承昨天的结果，昨天的结果又有两种可能 --- 卖、啥也不干，当然要继承它俩中大的那个，所以
* profitDoNothing[i] = Math.max(profitSell[i-1], profitDoNothing[i-1])。
* 重点来了，profitSell[i]怎么算。其实也有两种情况，如果day[i-1]我买入了（或者卖了，因为卖完了可以再买回来），则day[i]可以直接卖，收益是:
* profitSell[i] = profitSell[i-1] + prices[i] - prices[i-1]。但是还有一种情况，就是day[i-1]是啥也没干的，所以day[i]需要把自己先买回来，再卖，收益是：
* profitSell[i] = profitDoNothing[i-1] + prices[i] - prices[i] = profitDoNothing[i-1]。最终取这两个大的就行了。
 */
func maxProfit(prices []int) int {
	var n = len(prices)
	if n <= 1 {
		return 0
	}
	var profitDoNothing = make([]int, n+1) //当天没有买卖
	var profitSell = make([]int, n+1)      //当天有买卖
	profitDoNothing[0] = 0
	profitSell[0] = 0
	profitDoNothing[1] = 0
	profitSell[1] = 0
	for i := 2; i <= n; i++ {
		profitDoNothing[i] = max(profitDoNothing[i-1], profitSell[i-1])
		profitSell[i] = max(profitSell[i-1]+prices[i-1]-prices[i-2], profitDoNothing[i-1])
	}
	return max(profitSell[n], profitDoNothing[n])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
	fmt.Println(maxProfit([]int{1}))
}
