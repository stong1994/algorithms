package classify

import "math"

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
