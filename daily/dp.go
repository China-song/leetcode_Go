package daily

// MaxProfit123
// 123. 买卖股票的最佳时机 III
func MaxProfit123(prices []int) int {
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0

	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return sell2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MaxProfit309
// 309. 最佳买卖股票时机含冷冻期
func MaxProfit309(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	// f[i][]表示第i天后所获最大利润
	// f[i][0]	持有股票
	// f[i][1]	不持有股票 冷却期
	// f[i][2]	不持有股票 无冷却
	f := make([][3]int, n)
	f[0][0] = -prices[0]

	for i := 1; i < n; i++ {
		// f[i-1][0]: 持有的是i-1天结束时的股票
		// f[i-1][2]-prices[i]: 第i天买了股票(i-1天结束时无股票且无冷却)
		f[i][0] = max(f[i-1][0], f[i-1][2]-prices[i])
		// f[i-1][0]+prices[i]: 将i-1天结束时持有的股票卖掉
		// f[i-1][2]: i-1天结束时无股票无冷却，第i天买了又卖掉
		f[i][1] = max(f[i-1][0]+prices[i], f[i-1][2])
		// 当天无任何操作且结束时无股票， 表示前一天无股票
		f[i][2] = max(f[i-1][1], f[i-1][2])
	}

	return max(f[n-1][1], f[n-1][2])
}

// MaxProfit714
// 714. Best Time to Buy and Sell Stock with Transaction Fee
func MaxProfit714(prices []int, fee int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	// f[i][0]: 第i天结束后有股票 所获最大利润
	// f[i][1]: 第i天结束后无股票 所获最大利润
	f := make([][2]int, n)
	f[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		f[i][0] = max(f[i-1][0], f[i-1][1]-prices[i])
		f[i][1] = max(f[i-1][0]+prices[i]-fee, f[i-1][1])
	}
	return f[n-1][1]
}
