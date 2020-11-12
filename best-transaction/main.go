package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	dp := make([][2]int, len(prices))
	dp[0][0] = 0

	// last price to check
	dp[0][1] = -prices[0]

	fmt.Println(dp[0])
	for i := 1; i < len(prices); i++ {
		// dp[i][0]: value gained till last check
		// if dp[i-1][1] + prices[i] > 0, means go up
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])

		//
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])

		fmt.Println(dp[i])
	}

	return dp[len(prices)-1][0]
}

func maxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	total := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			total += prices[i] - prices[i-1]
		}
	}

	return total
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	strs := []int{7, 1, 1, 3, 6, 6}

	fmt.Println(maxProfit(strs))

	fmt.Println(maxProfit2(strs))
}
