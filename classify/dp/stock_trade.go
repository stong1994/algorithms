package dp

/** 股票交易类问题**/

// 1. 给定一个数组 prices ，它的第i 个元素prices[i] 表示一支给定股票第 i 天的价格。
// 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
// 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock
func maxProfit(prices []int) int {
	// dp做法：确定子问题——前i天卖出股票的最大收益为dp[i][0],买入股票的最大收益为dp[i][1]
	// 对于每天的交易有两种状态：买入和卖出
	// 如果前i天买入，则买入的最大收益为要么为前i-1天的最大买入收益，要么为今天买入，
	//		则dp[i][1] = max(-prices[i-1], dp[i-1][1])
	// 如果前i天卖出，则卖出的最大收益为要么前i-1天的最大卖出收益，要么为今天卖出，
	//		则dp[i][0] = max(dp[i-1][1] + prices[i-1], dp[i-1][0])
	if len(prices) == 0 {
		return 0
	}
	dp := make([][2]int, len(prices)+1)
	dp[0][1] = -prices[0]
	for i := 1; i <= len(prices); i++ {
		dp[i][1] = max(-prices[i-1], dp[i-1][1])
		dp[i][0] = max(dp[i-1][1]+prices[i-1], dp[i-1][0])
	}
	return dp[len(prices)][0]
}
