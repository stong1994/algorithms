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

// 2. 给定一个数组 prices ，其中prices[i] 是一支给定股票第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii
func maxProfit2(prices []int) int {
	// 定义重叠子问题：第i天交易的最大收益取决于前i-1天交易的最大收益
	// 每天有两种状态：买入和卖出。令dp[i][0]表示第i天买入的最大收益，dp[i][1]表示第i天卖出的最大收益
	// 如果第i天为买入状态，则收益可为第i天买入或前i-1天买入，选择其最大值，即
	//		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i-1])
	// 如果第i天为卖出状态，则收益为第i天卖出或前i-1天卖出，选择其最大值，即
	//		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i-1])
	if len(prices) == 0 {
		return 0
	}
	dp := make([][2]int, len(prices)+1)
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i <= len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i-1])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i-1])
	}
	return dp[len(prices)][1]
}

// 3. 给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成两笔交易。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii
func maxProfit3_(prices []int) int {
	// 对于第i天的交易，有四种状态：第一次买入、第一次卖出、第二次买入、第二次卖出
	// 分别设为dp[i][0]、dp[i][2]、dp[i][3]、dp[i][4]
	// 对于第一次买入状态，可以选择第i天买入，或者前i-1天买入。dp[i][0] = max(-prices[i-1], dp[i-1][0])
	// 对于第一次卖出状态，可以选择第i天卖出，或者前i-1天卖出。dp[i][1] = max(dp[i-1][0]+prices[i-1], dp[i-1][1])
	// 对于第二次买入状态，可以选择第i天买入，或者前i-1天买入。dp[i][2] = max(dp[i-1][1]-prices[i-1], dp[i-1][2])
	// 对于第二次卖出状态，可以选择第i天卖出，或者前i-1天卖出。dp[i][3] = max(dp[i-1][2]+prices[i-1], dp[i-1][3])
	if len(prices) == 0 {
		return 0
	}
	dp := make([][4]int, len(prices)+1)
	dp[0][0] = -prices[0]
	dp[0][2] = -prices[0]
	for i := 1; i <= len(prices); i++ {
		dp[i][0] = max(-prices[i-1], dp[i-1][0])
		dp[i][1] = max(dp[i-1][0]+prices[i-1], dp[i-1][1])
		dp[i][2] = max(dp[i-1][1]-prices[i-1], dp[i-1][2])
		dp[i][3] = max(dp[i-1][2]+prices[i-1], dp[i-1][3])
	}
	return max(dp[len(prices)][1], dp[len(prices)][3])
}

// 4. 给定一个整数数组prices ，它的第 i 个元素prices[i] 是一支给定的股票在第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv
// 注意：一次买+一次卖算一笔交易
func maxProfit4_(k int, prices []int) int {
	// 对于第i天的交易，有2k个状态：
	// 	对于第k次交易中的买，其状态对应的数字为2k
	// 	对于第k次交易中的卖，其状态对应的数字为2k+1
	// 第i天为第j个买入状态，则当天的收益为当天买入（第i-1天为第j-1次卖出状态）或者前i-1天第j次买入
	// 		dp[i][2*j] = max(dp[i-1][2*j], dp[i-1][2*(j-1)+1] - prices[i-1])
	// 第i天为第j个卖出状态，则当天的收益为当天卖出(前i-1天为第j次买入状态)或者前i-1天第j次卖出
	//		dp[i][2*j+1] = max(dp[i-1][2*j+1], dp[i-1][2*j] + prices[i-1])
	if len(prices) <= 1 {
		return 0
	}
	dp := make([][]int, len(prices)+1)
	for i := range dp {
		dp[i] = make([]int, 2*k+2)
	}
	for j := 0; j <= k; j++ {
		dp[0][j*2] = -prices[0] // 初始第0天的买入状态, 防止后续计算第1天买入时max后为0
	}
	for i := 1; i <= len(prices); i++ {
		for j := 1; j <= k; j++ {
			dp[i][2*j] = max(dp[i-1][2*j], dp[i-1][2*(j-1)+1]-prices[i-1])
			dp[i][2*j+1] = max(dp[i-1][2*j+1], dp[i-1][2*j]+prices[i-1])
		}
	}
	var result int
	for j := 1; j <= k; j++ {
		result = max(result, dp[len(prices)][j*2+1])
	}
	return result
}

// 5. 给定一个整数数组，其中第i个元素代表了第i天的股票价格 。
// 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
// 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
// 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown
func maxProfit5_(prices []int) int {
	// 对于第i天，有3种状态：买入、卖出和冷冻期（冷冻期影响的是买入，即第i天买入时，第i-1天不能卖出）
	// 如果在第i天为买入状态，则可以在第i天买入（受冷冻期影响，此时收益为前i-2天卖出的收益-第i天加个）或者在前i-1天买入
	// 		即 dp[i][0] = max(dp[i-1][0], dp[i-2][1]-prices[i-1])
	// 如果在第i天为卖出状态，则可以在第i天卖出或者在前i-1天卖出
	// 		即 dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i-1])
	if len(prices) <= 1 {
		return 0
	}
	if len(prices) == 2 {
		return max(0, prices[1]-prices[0])
	}

	dp := make([][2]int, len(prices)+1)
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	dp[1][0] = -prices[0]
	dp[1][1] = 0
	for i := 2; i <= len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-2][1]-prices[i-1])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i-1])
	}
	return dp[len(prices)][1]
}

// 6. 给定一个整数数组prices，其中第i个元素代表了第i天的股票价格 ；整数fee 代表了交易股票的手续费用。
// 你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。
// 返回获得利润的最大值。
// 注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee
func maxProfit6_(prices []int, fee int) int {
	// 对于第i天，有两个状态:买入和卖出
	// 如果第i天为买入状态，那么可以选择第i天买入或者前i-1天买入，即
	//		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i-1])
	// 如果第i天为卖出状态，那么可以选择第i天卖出或者前i-1天卖出，即
	// 		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i-1]-fee)
	if len(prices) <= 1 {
		return 0
	}
	dp := make([][2]int, len(prices)+1)
	dp[0][0] = -prices[0]
	for i := 1; i <= len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i-1])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i-1]-fee)
	}
	return dp[len(prices)][1]
}
