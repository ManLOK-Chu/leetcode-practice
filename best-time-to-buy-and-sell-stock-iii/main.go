package main

import "fmt"

/*
 * 正确的思路是使用“局部最优和全局最优”的思想 --- 和cooldown那个很像， 利用两个动态规划数组local和global。 其中global[i][j]表示到第i天最多进行j次交易可以达到的最大利润，
 * 而local[i][j]表示到第i天最多进行j次交易，并且在第i天卖出可以达到的最大利润。
 * 则： global[i][j] = Math.max(local[i][j], global[i-1][j]) 也就是到今天为止的全局最好，是到今天为止的局部最好与到昨天的全局最好中大的那个
 * 因为今天的最大利润有两种情况，第一是几天卖出了（既local[i][j]），另一种是今天什么都没做，那今天的最好的情况就是昨天的全局最好情况
 * 而： local[i][j] = Math.max(global[i-1][j-1] + prices[i] - prices[i-1], local[i-1][j] + prices[i] - prices[i-1]) --- 这里是local[i-1][j]因为就算第i-1天卖了， 可以再买回来在第i天再卖，所以并不增加transaction
 * 这里就不太好理解了， 因为今天需要卖出，所以需要加prices[i] - prices[i-1], 两种加的可能性，（1） 和全局比，（2）和局部比
 *
 * 如果只是两次操作，相当于j = 2，而且空间上维护当天数据即可，不需要i
 */
func maxProfit(prices []int) int {
	var n = len(prices)
	if n <= 1 {
		return 0
	}
	var local = make([][3]int, n)
	var global = make([][3]int, n)
	for i := 1; i < n; i++ {
		profitToday := prices[i] - prices[i-1]
		for j := 2; j >= 1; j-- {
			local[i][j] = max(global[i-1][j-1]+profitToday, local[i-1][j]+profitToday)
			global[i][j] = max(local[i][j], global[i-1][j])
		}
	}
	return global[n-1][2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
}
