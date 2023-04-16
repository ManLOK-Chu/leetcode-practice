package main

import "fmt"

//递推公式：dp[i] = max(dp[i - 2] + nums[i], dp[i - 1])
//由于不能连续偷两家, 那么, 第dp[i]则有两种可能, (1)偷当前这家, (2)不偷当前这家, 由于小偷希望偷到最大值, 则最优解为(1)和(2)中大的那个,
//既, dp[i]的值即为nums[i]+dp[i-2]和dp[i-1]中大的那个
func rob(nums []int) int {
	var n = len(nums)
	if n <= 1 {
		return nums[0]
	}
	var dp = make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	//从前到后遍历
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
	fmt.Println(rob([]int{1}))
	fmt.Println(rob([]int{1, 2}))
}
