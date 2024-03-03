package 数组_字符串

import (
	"fmt"
	"testing"
)

// 给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。
//
// 在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。
//
// 返回 你能获得的 最大 利润 。

// 输入：prices = [7,1,5,3,6,4]
// 输出：7
// 解释：在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4 。
// 随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6 - 3 = 3 。
// 总利润为 4 + 3 = 7 。

// 我的思路：
// 无非是记录我买卖过程中，总共能赚多少钱，核心思想就是记录收入和支出
func maxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	// 用一个二维数组记录收入和支出
	dp := make([][2]int, len(prices)) // dp[i][j], j = 0是持有现金，j=1是支出（支出也可以是买股票赚钱后的）

	dp[0][0] = 0          // 第一天没现金为0 ==》卖
	dp[0][1] = -prices[0] // 第一天买股票的支出 ==》买

	// 从index=1开始，第一天没收入不用遍历
	for i := 1; i < len(prices); i++ {
		// 第i天的现金 = 第i-1天的买股票支出+卖掉股票的收入（现金），如果发现第i天买，导致现金还不如之前的，那么就不买，现金还是保持上一天的
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		// 第i天的支出 = 第i-1的持有现金-买股票的钱，如果发现第i天买，导致总支出大于前一天(也就是dp[i-1][1]比dp[i][1]还小)，就不买
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]) // 支出大会导致dp[i][1]变小
		fmt.Printf("cash:%d,stock:%d \n", dp[i][0], dp[i][1])
	}
	return dp[len(prices)-1][0]
}

func maxProfit2_(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	cash := 0           // 持有现金 ==》卖
	stock := -prices[0] // 股票支出 ==》买

	for _, item := range prices {
		cash = max(cash, stock+item)  // 股票支出 + 当日价格 = 现金
		stock = max(stock, cash-item) // 现金 - 当日价格 = 当前股票支出后的收益
	}
	return cash
}

// 贪心
// 其实就是若相邻两数是升高，则累加其升高的值
func maxProfit__(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i-1] < prices[i] {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}

func Test(t *testing.T) {
	data := []int{7, 1, 5, 7, 6, 7}
	fmt.Println(maxProfit2(data))
}
