package dp

import (
	"math"
)

// 0-1 背包
// 有一个容量为 N 的背包，要用这个背包装下物品的价值最大，这些物品有两个属性：体积 w 和价值 v。
// W背包总体积，weights表示物品的体积数组，values表示物品的价值数组
func bag(W int, weights, values []int) int {
	// dp[i][j]表示前i个物品体积不超过j下能达到的最大价值
	// - 如果第i个物品放到背包中，那么dp[i][j] = dp[i-1][j-w]+v
	// - 如果第i个物品不放到背包，那么dp[i][j] = dp[i-1][j]
	// 需要进行比较来选择，所以 dp[i][j] = max(dp[i][j-w]+v, dp[i-1][j])
	var dp [][]int
	N := len(weights) + 1
	for i := 0; i <= N; i++ {
		dp = append(dp, make([]int, W+1))
	}
	for i := 1; i < N; i++ {
		for j := 1; j < W; j++ {
			w, v := weights[i-1], values[i-1]
			if j >= w {
				dp[i][j] = max(dp[i][j-w]+v, dp[i-1][j])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[N][W]
}

// 空间优化
// 状态转移方程 dp[i][j] = max(dp[i][j-w]+v, dp[i-1][j]) 中i只与i-1有关
// 可以优化为 dp[j] = max(dp[j-w]+v, dp[j])
func bagOpt(W int, weights, values []int) int {
	dp := make([]int, W+1)
	N := len(weights) + 1
	for i := 1; i <= N; i++ {
		w, v := weights[i-1], values[i-1]
		for j := W; j > 0; j-- {
			if j >= w {
				dp[j] = max(dp[j-w]+v, dp[j])
			}
		}
	}
	return dp[W]
}

// 分割等和子集*****
// 给你一个 只包含正整数 的 非空 数组nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
// 示例 1：
//	输入：nums = [1,5,11,5]
//	输出：true
//	解释：数组可以分割成 [1, 5, 5] 和 [11] 。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/partition-equal-subset-sum
func canPartition(nums []int) bool {
	// 简化问题：划分成两个和相等的子集=》存在子集使得其和为元素和的一半
	// 解法：1. 先根据特性判断特殊情况
	//      2. 利用动态规划对元素进行组合，一旦得到sum/2，则返回true
	// 动态规划：dp[i][j] 对于前i个数字能否得到j
	// - 如果j>=nums[i],则可选可不选，两者有一个为true即可
	//	- 如果不选：dp[i][j] = dp[i-1][j]
	//	- 如果选：dp[i][j] = dp[i-1][j-nums[i]]
	// - 如果j<nums[i],则不可选：dp[i][j] = dp[i-1][j]
	L := len(nums)
	if L < 2 {
		return false
	}
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum&1 == 1 {
		return false
	}
	halfSum := sum / 2
	var dp [][]bool
	for i := 0; i <= L; i++ {
		dp = append(dp, make([]bool, halfSum+1))
	}
	// 初始化：当j为0时，不选取的结果一定为0，所以d[i][0] = true
	for i := 0; i <= L; i++ {
		dp[i][0] = true
	}

	for i := 1; i <= L; i++ {
		for j := 1; j <= halfSum; j++ {
			if j >= nums[i-1] {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[L][halfSum]
}

func canPartitionOpt(nums []int) bool {
	// 由上知
	// 动态规划：dp[i][j] 对于前i个数字能否得到j
	// - 如果j>=nums[i],则可选可不选，两者有一个为true即可
	//	- 如果不选：dp[i][j] = dp[i-1][j]
	//	- 如果选：dp[i][j] = dp[i-1][j-nums[i]]
	// - 如果j<nums[i],则不可选：dp[i][j] = dp[i-1][j]
	// dp[i]至于dp[i-1]有关，因此可进行优化：
	// 动态规划：dp[j] = dp[j] || dp[j-nums[i]]
	L := len(nums)
	if L < 2 {
		return false
	}
	var (
		sum int
		max int
	)
	for _, num := range nums {
		sum += num
		if num > max {
			max = num
		}
	}
	if sum&1 == 1 {
		return false
	}
	halfSum := sum / 2
	if max > halfSum {
		return false
	}

	dp := make([]bool, halfSum+1)
	// 初始化：当j为0时，不选取的结果一定为0，所以d[0] = true
	dp[0] = true

	for i := 0; i < L; i++ {
		v := nums[i]
		// 逆循环
		for j := halfSum; j >= v; j-- {
			dp[j] = dp[j] || dp[j-v]
		}
	}
	return dp[halfSum]
}

// 目标和*****
// 给你一个整数数组 nums 和一个整数 target 。
// 向数组中的每个整数前添加'+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
// 例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
// 返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/target-sum
func findTargetSumWays(nums []int, target int) int {
	// 简化问题：
	// 1. target可以为负数，但是只要将所有的符合正负替换，就能够得到正数，因此可将target视为正数
	// 2. nums的总和sum为最大值，sum必须大于target，那么需要去掉和为neg的数。sum - neg = target + neg，即neg = (sum-target)/2
	// 则问题转换为在nums找到几个元素，使其和为neg，求有几种组合方式=>回溯法或者dp
	// dp[i][j]: 前i个元素和为j的组合方式
	// 对于第i个元素
	//	- 如果nums[i]>j, 则dp[i][j] = dp[i-1][j]
	//	- 如果nums[i]<=j
	//		- 如果使用第i个元素，则dp[i][j] = dp[i-1][j-nums[i]]
	//		- 如果不使用第i个元素，则dp[i][j] = dp[i-1][j]
	if target < 0 {
		target = -target
	}
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum < target {
		return 0
	}
	residue := sum - target
	if residue&1 == 1 {
		return 0
	}
	neg := residue / 2

	var (
		dp [][]int
		L  = len(nums)
	)
	for i := 0; i < L+1; i++ {
		dp = append(dp, make([]int, neg+1))
	}
	// 当target为0时，如果没有选择任意元素时，其组合数为1
	dp[0][0] = 1
	for i := 1; i < L+1; i++ {
		num := nums[i-1]
		for j := 0; j <= neg; j++ {
			if num > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-num]
			}
		}
	}
	return dp[L][neg]
}

// 字符构成最多的字符串
// 给你一个二进制字符串数组 strs(strs仅有0和1组成)和两个整数 m 和 n 。
// 请你找出并返回 strs 的最大子集的长度，该子集中 最多 有 m 个 0 和 n 个 1 。
// 如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/ones-and-zeroes
func findMaxForm(strs []string, m int, n int) int {
	// dp[i][j][k]：前i个元素能够构成的最多j个0,k个1的最大子集的长度.count[i][0]表示strs[i]中0的数量，count[i][1]表示strs[i]中1的数量
	// 对于第i个元素：
	// 	- j > count[i][0]，或者 k>count[i][1]，那么不存放，此时那么dp[i][j][k] = dp[i-1][j][k]
	//  - 否则,那么可以存放，也可以不存放
	//      - 如果存放：dp[i][j][k] = dp[i-1][j-count[i][0]][k-count[i][1]]
	//		- 如果不存放：那么dp[i][j][k] = dp[i-1][j][k]
	//      - 取上边两种情况的最大值
	var (
		dp [][][]int
		L  = len(strs)
	)
	for i := 0; i < L+1; i++ {
		var tmp [][]int
		for j := 0; j < m+1; j++ {
			tmp = append(tmp, make([]int, n+1))
		}
		dp = append(dp, tmp)
	}

	// 处理边界，当i为0时，意味着没有任何可供选择的元素，那么最大子集的长度一定为0.则dp[0][j][k]都为0.但是已默认为0，因此无需处理
	for i := 1; i < L+1; i++ {
		num0, num1 := getNum0AndNum1(strs[i-1])
		for j := 0; j < m+1; j++ {
			for k := 0; k < n+1; k++ {
				if j < num0 || k < num1 {
					dp[i][j][k] = dp[i-1][j][k]
				} else {
					dp[i][j][k] = max(dp[i-1][j-num0][k-num1]+1, dp[i-1][j][k])
				}
			}
		}
	}
	return dp[L][m][n]
}

func getNum0AndNum1(str string) (num0 int, num1 int) {
	for _, v := range str {
		if v == '0' {
			num0++
		} else {
			num1++
		}
	}
	return
}

// 找零钱的最少硬币数
// 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
// 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回-1 。
// 你可以认为每种硬币的数量是无限的。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/coin-change
func coinChange(coins []int, amount int) int {
	// dp[i] 金额为i时使用的最少的硬币数量
	// dp[i] = min(dp[i-coins[0]], ... dp[i-coins[j]]) +1
	dp := make([]int, amount+1)

	// 初始化，将dp中的元素初始化为一个最大值
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// 找零钱的硬币数组合 ******
// 给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。
// 请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。
// 假设每一种面额的硬币有无限个。
// 题目数据保证结果符合 32 位带符号整数。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/coin-change-2
// 和爬楼梯问题的异同：https://leetcode-cn.com/problems/coin-change-2/solution/ling-qian-dui-huan-iihe-pa-lou-ti-wen-ti-dao-di-yo/
func change(amount int, coins []int) int {
	// 对于第i个硬币达到金额j的组合数，第i个硬币
	// 	- 如果coin大于金额j，则不放，那么dp[i][j] = dp[i-1][j]
	// 	- 否则，其组合数等于（加入coin后和为j的组合数+不加入coin和为j的组合数） 那么dp[i][j] = dp[i][j-coin] +dp[i-1][j]
	L := len(coins)
	dp := make([][]int, L+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
	}
	// 当没有硬币时，即i为0时，组合数为0
	// 对总和为0的任意硬币的组合数为1
	for i := range dp {
		dp[i][0] = 1
	}
	for i := 1; i <= L; i++ {
		coin := coins[i-1]
		for j := 1; j <= amount; j++ {
			if j >= coin {
				dp[i][j] = dp[i-1][j] + dp[i][j-coin]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	//fmt.Println(dp)
	return dp[L][amount]
}

func change2(amount int, coins []int) int {
	// coins中的硬币可以选取多次，求能够组合为总价值为amount的组合数
	// dp[i]为能够组合为i的选取的硬币的数量
	// - 对于i等于0时，任意coins组合为0的组合数都为1
	// - 对于coin<=j<=amount, 组合数为其子集（i-coin）的组合之和,即 d[i]+=d[i-coin]
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins { // 重点理解 两层循环的先后顺序，以及空闲优化后的内存循环逆序问题
		for j := coin; j <= amount; j++ {
			dp[j] += dp[j-coin]
		}
	}
	return dp[amount]
}

// 单词拆分
// 给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 。
// 注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/word-break
func wordBreak(s string, wordDict []string) bool {
	// dp[i]: s中前i个字符串能否被拼接
	existWord := make(map[string]bool)
	for _, word := range wordDict {
		existWord[word] = true
	}
	L := len(s)
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= L; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && existWord[s[j:i]] {
				dp[i] = true
			}
		}
	}
	return dp[L]
}

// 组合总和 Ⅳ
// 给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素组合的个数。
// 题目数据保证答案符合 32 位整数范围。
// 请注意，顺序不同的序列被视作不同的组合。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/combination-sum-iv
func combinationSum4(nums []int, target int) int {
	// 子问题：在nums中找到和为i的组合数
	// 遍历nums中的元素j，对于j<=i，有 dp[i] += dp[i-j]
	dp := make([]int, target+1)
	dp[0] = 1 // 对于目标为0的组合数，任意nums都只有一种组合方式
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}

// 最佳买卖股票时机含冷冻期
// 给定一个整数数组，其中第i个元素代表了第i天的股票价格 。
// 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
// 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
// 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown
func maxProfit3(prices []int) int {
	// 子问题：第i天的最大利益与前i天的最大收益有关
	// dp[i]: 第i天的最大收益
	// 如果第i天不买入/卖出，则第i天的最大收益等于第i-1天的最大收益
	// 如果第i天卖出，假设第j天买入（1<=j<i）则第i天的最大收益等于第j-2天（冷冻一天）的最大收益加上第i天卖出的收益prices[i-1]-prices[j-1]
	L := len(prices)
	dp := make([]int, L+1)
	dp[0] = 0
	for i := 1; i <= L; i++ {
		for j := 1; j < i; j++ {
			if j >= 2 {
				dp[i] = max(dp[i], max(dp[i-1], dp[j-2]+prices[i-1]-prices[j-1]))
			} else {
				dp[i] = max(dp[i], max(dp[i-1], dp[j-1]+prices[i-1]-prices[j-1]))
			}
		}
	}
	//fmt.Println(dp)
	return dp[L]
}

func maxProfit4(prices []int) int {
	// 子问题：第i天的最大利益与前i天的最大收益有关
	// dp[i]: 第i天的最大收益
	// 第i天的状态：持有股票、不持有股票、冷冻期(第i天卖出后，即第i+1天不能买入)，对应的收益分别记为dp[i][0]、dp[i][1]、dp[i][2]
	// 如果第i天持有股票，则分为两种情况
	//	- 前i-1天买入，此时最大收益dp[i][0] = dp[i-1][0]
	//	- 第i天买入，此时最大收益dp[i][0] = dp[i-1][1]-prices[i-1]
	// 如果第i天不持有股票，即处于观望期，此时最大收益为前i-1天的最大收益(分为两种状态，第i-1天不持有股票或者冷冻期)
	//		dp[i][1] = max(dp[i-1][1], dp[i-1][2])
	// 如果第i天处于冷冻期，此时卖出了前i-1天持有的股票 最大收益等于dp[i][2]= dp[i-1][0] + prices[i-1]
	if len(prices) == 0 {
		return 0
	}
	L := len(prices)
	dp := make([][3]int, L+1)
	dp[1][0] = -prices[0]
	dp[1][1] = 0
	dp[1][2] = 0
	for i := 2; i <= L; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i-1])
		dp[i][1] = max(dp[i-1][1], dp[i-1][2])
		dp[i][2] = dp[i-1][0] + prices[i-1]
	}

	//fmt.Println(dp)
	return max(dp[L][1], dp[L][2])
}

// 需要交易费用的股票交易
// 给定一个整数数组prices，其中第i个元素代表了第i天的股票价格 ；整数fee 代表了交易股票的手续费用。
// 你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。
// 返回获得利润的最大值。
// 注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee
func maxProfit5(prices []int, fee int) int {
	// 子问题：第i天的最大收益取决于前i-1天的最大收益
	// 两种状态：持有和未持有=》dp[i][1]、dp[i][0]
	// 如果第i天持有，分为两种：
	//	- 如果是第i天买入，那么其最大收益取决于第i-1天的未持有的收益-买入价格， 则dp[i][1] = dp[i-1][0] - prices[i-1]
	//	- 如果是前i-1天买入，那么其最大收益取决于第i-1天的已持有的收益， 则dp[i][1] = dp[i-1][1]
	// 如果第i天不持有，分为两种情况：
	// 	- 第i天卖出，那么最大收益为第i-1天持有的收益加上当天的价格再减去手续费,则dp[i][0] = dp[i-1][1]+prices[i-1]-fee
	// 	- 第i-1天已经未持有，那么最大收益与第i-1天相等，即dp[i][0] = dp[i-1][0]
	if len(prices) == 0 {
		return 0
	}
	L := len(prices)
	dp := make([][2]int, L+1)
	dp[1][0] = 0
	dp[1][1] = -prices[0]
	for i := 2; i <= L; i++ {
		dp[i][1] = max(dp[i-1][0]-prices[i-1], dp[i-1][1])
		dp[i][0] = max(dp[i-1][1]+prices[i-1]-fee, dp[i-1][0])
	}
	//fmt.Println(dp)
	return dp[L][0]
}

// 买卖股票的最佳时机 III
// 给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成两笔交易。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii
func maxProfit6(prices []int) int {
	// 与其他题目的区别在于限制了交易次数：最多两笔
	// 题目简化为：将prices划分最多两个区间，使得两个区间的最大值和最小值之差 相加最大
	// 对于分割索引i，0<=i<len(prices)
	var result int
	for i := 0; i < len(prices); i++ {
		diff := maxDiff(prices, 0, i) + maxDiff(prices, i+1, len(prices)-1)
		if diff > result {
			result = diff
		}
	}
	return result
}

func maxDiff(prices []int, start, end int) int {
	if start == end || start >= len(prices) {
		return 0
	}
	var (
		result  int
		tmpHigh = prices[start]
		tmpMin  = prices[start]
	)
	for i := start + 1; i <= end; i++ {
		if prices[i] >= tmpHigh {
			tmpHigh = prices[i]
		}
		if prices[i] < tmpMin {
			if result < tmpHigh-tmpMin {
				result = tmpHigh - tmpMin
			}
			tmpMin = prices[i]
			tmpHigh = prices[i]
		}
	}
	if result < tmpHigh-tmpMin {
		result = tmpHigh - tmpMin
	}
	return result
}

// 上述方法超时，使用动态规划
func maxProfit6Opt(prices []int) int {
	// 与其他题目的区别在于限制了交易次数：最多两笔
	// 子问题：第i次交易的最大收益取决于前i-1次交易
	// 5个状态：
	//	- 没有进行交易,没有持有股票 dp[i][0]
	//	- 没有进行交易，但持有股票 dp[i][1]
	//	- 进行第一次交易，没有持有股票（这次卖掉了股票） dp[i][2]
	//	- 进行第一次交易，有持有股票（之前卖掉了股票，此时又买入了股票） dp[i][3]
	//	- 进行了第二次交易，没有持有股票（第二次卖掉了股票） dp[i][4]

	// - 如果第i次交易之前没有进行交易，没有持有股票，则dp[i][0] = 0
	// - 如果第i次交易之前没有进行交易，但持有股票，则dp[i][1] = max(-prices[i-1], dp[i-1][1])
	// - 如果第i次交易进行了第一次交易，卖掉了股票，则dp[i][2] = max(dp[i-1][1] + prices[i-1], dp[i-1][2])
	// - 如果第i次之前进行了交易，此时持有股票 dp[i][3] = max(dp[i-1][2] - prices[i-1], dp[i-1][3])
	// - 如果第i次之前进行了第二次卖掉股票 dp[i][4] = max(dp[i-1][3] + prices[i-1], dp[i-1][4])
	L := len(prices)
	if L == 0 {
		return 0
	}
	dp := make([][5]int, L+1)
	dp[1][0] = 0
	dp[1][1] = -prices[0]
	dp[1][2] = 0
	dp[1][3] = -prices[0] // 假设一天能够进行重复交易，此时为买入第二笔
	dp[1][4] = 0
	for i := 2; i <= L; i++ {
		dp[i][0] = 0
		dp[i][1] = max(-prices[i-1], dp[i-1][1])
		dp[i][2] = max(dp[i-1][1]+prices[i-1], dp[i-1][2])
		dp[i][3] = max(dp[i-1][2]-prices[i-1], dp[i-1][3])
		dp[i][4] = max(dp[i-1][3]+prices[i-1], dp[i-1][4])
	}
	//fmt.Println(dp)
	return max(dp[L][2], dp[L][4])
}

// 买卖股票的最佳时机 IV
// 给定一个整数数组prices ，它的第 i 个元素prices[i] 是一支给定的股票在第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv
func maxProfit7(k int, prices []int) int {
	// 上题为k=2，有五种状态：未进行买、第一次买、第一次卖、第二次买、第二次卖
	//  如果允许进行k笔交易，则有2k+1种状态
	L := len(prices)
	if L == 0 {
		return 0
	}
	dp := make([][]int, L+1)
	for i := range dp {
		dp[i] = make([]int, 2*k+1)
	}
	// 初始化第一天的状态：买的收益为-prices[0]
	for i := 1; i < 2*k+1; i += 2 {
		dp[1][i] = -prices[0]
	}
	for i := 2; i <= L; i++ {
		for j := 1; j < 2*k+1; j++ { // 状态为0时，表示未买入，收益永远为0，不用考虑
			if j&1 == 1 { // 买操作
				dp[i][j] = max(dp[i-1][j-1]-prices[i-1], dp[i-1][j])
			} else { // 卖操作
				dp[i][j] = max(dp[i-1][j-1]+prices[i-1], dp[i-1][j])
			}
		}

	}
	//fmt.Println(dp)
	var result int
	for i := 2; i < 2*k+1; i += 2 {
		if result < dp[L][i] {
			result = dp[L][i]
		}
	}
	return result
}
