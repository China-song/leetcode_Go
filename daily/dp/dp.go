package dp

import (
	"sort"
	"strconv"
)

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

	// sell: 第i天结束后 无股票 所获最大利润
	// buy:  第i天结束后 有股票 所获最大利润
	sell, buy := 0, -prices[0]

	for i := 1; i < n; i++ {
		// 由于第i天结束时持有的利润只与前一天有关 所以只需要几个变量
		tmp := sell // sell可能会变化
		sell = max(sell, buy+prices[i]-fee)
		buy = max(buy, tmp-prices[i])
	}
	return sell
}

/*
MaxSatisfaction1402
1402. Reducing Dishes
*/
func MaxSatisfaction1402(satisfaction []int) int {
	n := len(satisfaction)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	sort.Ints(satisfaction)
	res := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i][j] = dp[i-1][j-1] + satisfaction[i-1]*j
			if j < i {
				dp[i][j] = max(dp[i][j], dp[i-1][j])
			}
			res = max(res, dp[i][j])
		}
	}
	return res
}

/*
NumRollsToTarget1155

1155. Number of Dice Rolls With Target Sum

Middle Dynamic Programming
*/
func NumRollsToTarget1155(n int, k int, target int) int {
	// 方法一， 使用二维数组
	// 分治 f(n, target) = f(n-1, target-1) + f(n-1, target-2) + ... + f(n-1, target-k) (target-x>=0)
	//f := make([][]int, n+1)
	//for i := 0; i <= n; i++ {
	//	f[i] = make([]int, target+1)
	//}
	//f[0][0] = 1 // no need to roll die f[1][x] = f[0][x-1] + f[0][x-2] + .. f[0][0] ( =1 if x<=k else =0)
	//for i := 1; i <= n; i++ {
	//	for j := 1; j <= target; j++ {
	//		// 分治 f(i, j) = f(i-1, j-1) + f(n-1, j-2) + ... + f(n-1, j-k) (j-x>=0)
	//		for x := 1; x <= k; x++ {
	//			if j-x >= 0 {
	//				f[i][j] = (f[i][j] + f[i-1][j-x]) % (1e9 + 7)
	//			}
	//		}
	//	}
	//}
	//return f[n][target]

	// 方法二，使用两个一维数组
	//mod := int(1e9 + 7)
	//f := make([]int, target+1)
	//f[0] = 1bC
	//for i := 1; i <= n; i++ {
	//	g := make([]int, target+1)
	//	for j := 1; j <= target; j++ {
	//		for x := 1; x <= k; x++ {
	//			if j-x >= 0 {
	//				g[j] = (g[j] + f[j-x]) % mod
	//			}
	//		}
	//	}
	//	f = g
	//}
	//return f[target]

	// 方法三，只用一个一维数组
	mod := int(1e9 + 7)
	f := make([]int, target+1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		for j := target; j >= 1; j-- {
			f[j] = 0
			for x := 1; x <= k; x++ {
				if j-x >= 0 {
					f[j] = (f[j] + f[j-x]) % mod
				}
			}
		}
	}
	return f[target]
}

/*
find i, 1<=i<=n
*/
func punishmentNumber(n int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		if isPunish(i, i*i) {
			ans += i
		}
	}
	return ans
}

func isPunish(i, square int) bool {
	squareStr := strconv.Itoa(square)
	return dfs(i, 0, 0, squareStr, 0)
}

func dfs(i int, start int, end int, squareStr string, sum int) bool {
	if start >= len(squareStr) {
		if sum == i {
			return true
		} else {
			return false
		}
	}

	// 是否将[start, end]作为一个数字 加入到sum中
	// 是
	num, _ := strconv.Atoi(squareStr[start : end+1])
	if dfs(i, end+1, end+1, squareStr, sum+num) {
		return true
	}
	// 否
	if dfs(i, start, end+1, squareStr, sum) {
		return true
	}
	return false
}
