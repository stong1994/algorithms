package dp

import (
	"math"
	"sort"
)

// 爬楼梯
// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
// 注意：给定 n 是一个正整数。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/climbing-stairs
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	var (
		dp = make([]int, n+1) // 从 dp[i] = dp[i-1]+dp[i-2] 解释：爬到第i个台阶的方式=爬到第i-1个台阶的方式+爬到第i-2个台阶的方式
	)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func climbStairsFib(n int) int {
	var (
		dp = make([]int, n+1)
	)
	var fib func(m int) int
	fib = func(m int) int {
		if m <= 2 {
			return m
		}
		if dp[m] > 0 {
			return dp[m]
		}
		rst := fib(m-1) + fib(m-2)
		dp[m] = rst
		return rst
	}
	return fib(n)
}

// 爬梯子的回溯法。dp优于回溯就在于db能够利用之前的“经验”来加快计算，比如登两个台阶有两种方式，那么四个台阶就有2*2个方式
func climbStairs_backTrack(n int) int {
	var (
		result int
	)
	var backTrack func(int)
	backTrack = func(idx int) {
		if idx == n {
			result++
			return
		}
		if idx > n {
			return
		}
		backTrack(idx + 1)
		backTrack(idx + 2)
	}
	backTrack(0)
	return result
}

// 打家劫舍
// 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/house-robber
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums)) // dp[i] = max(dp[i-2]+ nums[i], dp[i-1])
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var fib func(int) int
	fib = func(idx int) int {
		if dp[idx] > 0 {
			return dp[idx]
		}
		if idx == 0 {
			return nums[0]
		}
		if idx == 1 {
			dp[idx] = max(nums[0], nums[1])
			return dp[idx]
		}
		dp[idx] = max(fib(idx-2)+nums[idx], fib(idx-1))
		return dp[idx]
	}
	return fib(len(nums) - 1)
}

func robNormal(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums)) // dp[i] = max(dp[i-2]+ nums[i], dp[i-1])
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

// 空间优化
func robNormalOpt(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	first := nums[0]
	second := max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		first, second = second, max(first+nums[i], second)
	}
	return second
}

// 打家劫舍 II
// 你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/house-robber-ii
func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	// 要么有第一个元素，没最后一个元素；要么有最后一个元素，没第一个元素
	a := robNormalOpt(nums[1:])
	b := robNormalOpt(nums[:len(nums)-1])
	if a > b {
		return a
	}
	return b
}

// 信件错排
// 题目描述：有N个信和信封，它们被打乱，求错误装信方式的数量。
func envelopeFalse(n int) int {
	// 第N个信可以放到N-1个位置，对于第N个信封，
	// 如果有1个信封和信，则没有错误装信的方式
	// 如果有2个信封和信，则有一种错误装信的方式
	// 用dp[i]来表示i个信和信封被打乱的错误方式。
	// 对于第n个信，可以将其放到第k个信封中（k有n-1种可能）
	// 对于第k个信，如果把它放到第n个信封中，那么此时剩余n-2个信和信封的错误装信方式为dp[n-2]，那么此时的错误数量为 dp[n-2]*(n-1)
	// 对于第k个信，如果不把它放到第n个信封中，那么此时第k个信和信封的错误方式为dp[n-1]，那么此时的错误数量为 dp[n-1]*(n-1),即dp[n-1]*(n-1)
	// 综上：dp[n] = dp[n-2]*(n-1) + dp[n-1]*(n-1)
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 0
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) * (i - 1)
	}
	return dp[n]
}

// 题目描述：假设农场中成熟的母牛每年都会生 1 头小母牛，并且永远不会死。
// 第一年有 1 只小母牛，从第二年开始，母牛开始生小母牛。每只小母牛 3 年之后成熟又可以生小母牛。
// 给定整数 N，求 N 年后牛的数量。
func cowNum(n int) int {
	// 第i年成熟母牛的数量为：
	// dp[i] = dp[i-3]+dp[i-1] 找到i-3年出生的小母牛，i年的数量等于i-3年小母牛的数量，再加上i-1年的成熟母牛数量
	// 第i年的牛的数量：
	// rst[i] = dp[i]*2 第i年牛的数量为第i年成熟母牛的数量*2，再加上未成熟的母牛数量，即dp[i-2]+dp[i-1]
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	dp[2] = 3
	for i := 3; i < n; i++ {
		dp[i] = dp[i-3] + dp[i-1]
	}
	totalCount := dp[n]*2 + dp[n-2] + dp[n-1]
	return totalCount
}

// 区域和检索 - 数组不可变
// 给定一个整数数组 nums，求出数组从索引i到j（i≤j）范围内元素的总和，包含i、j两点。
// 实现 NumArray 类：
// NumArray(int[] nums) 使用数组 nums 初始化对象
// int sumRange(int i, int j) 返回数组 nums 从索引i到j（i≤j）范围内元素的总和，包含i、j两点（也就是 sum(nums[i], nums[i + 1], ... , nums[j])）
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/range-sum-query-immutable
type NumArray struct {
	sums []int // 前缀和
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}
	return NumArray{
		sums: sums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.sums[right]
	}
	return this.sums[right] - this.sums[left-1]
}

// 等差数列划分*****
// 如果一个数列 至少有三个元素 ，并且任意两个相邻元素之差相同，则称该数列为等差数列。
// 例如，[1,3,5,7,9]、[7,7,7,7] 和 [3,-1,-5,-9] 都是等差数列。
// 给你一个整数数组 nums ，返回数组 nums 中所有为等差数组的 子数组 个数。
// 子数组 是数组中的一个连续序列。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/arithmetic-slices
func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	dp := make([]int, len(nums))
	// 对于[1,2,3,4,5]
	// i为2时，dp[2]为1，即[1,2,3]
	// i为3时，dp[3]=dp[2]+1 = 2, 代表的是新增的等差数组数量，即[2,3,4]与[1,2,3,4]
	// 如果nums[i]与nums[i-1]不能形成等差数组，就需要重新计数
	// 对于一个连续的数列，dp[i] = dp[i-1]+1, dp[i]是第i个元素能新增的子数组的数量, 每次多一个元素，就加一种组合。
	// nums[i] - nums[i-1]= nums[i-1] - nums[i-2] => nums[i] = nums[i-1]*2 - nums[i-2]
	for i := 2; i < len(nums); i++ {
		if nums[i] == nums[i-1]*2-nums[i-2] {
			dp[i] = dp[i-1] + 1
		}
	}
	var result int
	for _, v := range dp {
		result += v
	}
	return result
}

// 整数拆分*****
// 给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。
// 说明: 你可以假设n不小于 2 且不大于 58。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/integer-break
func integerBreak(n int) int {
	// 对于dp[i], 1<=j<i
	// 如果i-j可拆，那么此时最大值为j*dp[i-j]
	// 如果i-j不可拆，那么此时最大值为j*(i-j)
	// 不用考虑j是否可拆，因为j只需要判断三种情况：1,2,3.因为其他所有的数字都可以由2和3来相乘得到。
	// 如i等于11，当j为6时，组合为(6,5)可以拆分为(3,3,5)的组合，因此，当j为3时，就已经计算过这种情况了
	// 即(3,8)包含了(6,5)中的(3,3,5),(2,9)包含了(6,5)中的(2,2,2,5)、(2,4,5)
	// 即dp[i] = max(j*(i-j), j*dp[i-j])
	dp := make([]int, n+1)
	// 能够满足n<2的场景，不用做特殊处理
	for i := 2; i <= n; i++ {
		for j := 1; j < i && j <= 3; j++ {
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 完全平方数
// 给定正整数n，找到若干个完全平方数（比如1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
// 给你一个整数 n ，返回和为 n 的完全平方数的 最少数量 。
// 完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/perfect-squares
func numSquares2(n int) int {
	// 对于整数i至少需要f(i)个完全平方数，平方数j的范围在[1, i^0.5]，
	// 当我们寻找j的平方数个数时，需要在[1, i-j^2]中查找完全平方数之和，与原问题相似，可以构成状态转移方程
	// dp(i) = 1 + min( dp(1), dp(2),... dp(i-j^2) ) j的范围为[1, i^0.5]
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		min := math.MaxInt64
		for j := 1; j*j <= i; j++ {
			if min > dp[i-j*j] {
				min = dp[i-j*j]
			}
		}
		dp[i] = min + 1
	}
	return dp[n]
}

// 解码方法
// 一条包含字母A-Z 的消息通过以下映射进行了 编码 ：
// 'A' -> 1
// 'B' -> 2
// ...
// 'Z' -> 26
// 要 解码 已编码的消息，所有数字必须基于上述映射的方法，反向映射回字母（可能有多种方法）。例如，"11106" 可以映射为：
// "AAJF" ，将消息分组为 (1 1 10 6)
// "KJF" ，将消息分组为 (11 10 6)
// 注意，消息不能分组为 (1 11 06) ，因为 "06" 不能映射为 "F" ，这是由于 "6" 和 "06" 在映射中并不等价。
// 给你一个只含数字的 非空 字符串 s ，请计算并返回 解码 方法的 总数 。
// 题目数据保证答案肯定是一个 32 位 的整数。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/decode-ways
func numDecodings(s string) int {
	// df[i]: 前i个字符能够编码的方式
	// 对第i个字符进行编码有两种方式：
	// 1. 单个字符编码，只要s不为'0'即可，此时df[i] += df[i-1]
	// 2. 双字符编码编码，此时s[i-1:i+1]需要不大于'26' df[i] += df[i-2]
	df := make([]int, len(s)+1)
	df[0] = 1
	for i := 1; i <= len(s); i++ {
		if s[i-1] != '0' {
			df[i] += df[i-1]
		}
		if i >= 2 && s[i-2] != '0' && s[i-2:i] <= "26" && s[i-2:i] > "0" {
			df[i] += df[i-2]
		}
	}
	return df[len(s)]
}

// 最长递增子序列
// 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
// 子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
// 进阶：
//	你可以设计时间复杂度为 O(n2) 的解决方案吗？
//	你能将算法的时间复杂度降低到 O(n log(n)) 吗?
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/longest-increasing-subsequence
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums)) // dp[i]：第i个元素之前的最长子序列长度
	for i := 0; i < len(nums); i++ {
		maxVal := 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				maxVal = max(maxVal, dp[j]+1)
			}
		}
		dp[i] = maxVal
	}
	var result int
	for _, v := range dp {
		if result >= v {
			continue
		}
		result = v
	}
	return result
}

// 优化时间复杂度为N*log(N): 使用二分查找优化以下内容
// for j := 0; j < i; j++ {
//			if nums[i] > nums[j] {
//				maxVal = max(maxVal, dp[j]+1)
//			}
//		}
// 使用tails数组,tails[i]用来存储长度为i的子序列的最后一个元素，对于一个元素x
// - 如果x大于tails中的元素，那么将x append 到tails的末尾，说明最长子序列的长度+1
// - 如果tails[i-1]<x<tails[i],则将tails[i]替换为x
func lengthOfLISOpt(nums []int) int {
	var tails []int
	binarySearch := func(val int) int {
		lo, hi := 0, len(tails)-1
		for lo <= hi {
			mid := lo + (hi-lo)/2
			if tails[mid] == val {
				return mid
			}
			if tails[mid] < val {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		return lo
	}

	storeTails := func(val int) {
		if len(tails) == 0 || val > tails[len(tails)-1] {
			tails = append(tails, val)
			return
		}
		tails[binarySearch(val)] = val
	}

	for _, num := range nums {
		storeTails(num)
	}
	return len(tails)
}

// 一组整数对能够构成的最长链*****
// 给出n个数对。在每一个数对中，第一个数字总是比第二个数字小。
// 现在，我们定义一种跟随关系，当且仅当b < c时，数对(c, d)才可以跟在(a, b)后面。我们用这种形式来构造一个数对链。
// 给定一个数对集合，找出能够形成的最长数对链的长度。你不需要用到所有的数对，你可以以任何顺序选择其中的一些数对来构造。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/maximum-length-of-pair-chain
func findLongestChain(pairs [][]int) int {
	// 对pairs根据pair[0]进行排序
	// 遍历pairs，构造以pair结尾的最长链dp, 即dp[i]为以pairs[i]结尾的链的最大长度，对小于i的所有pair j进行比较，如果符合要求，则将j的长度+1作为i的长度
	// 对于i>j且pairs[i][0]>pairs[j][1],那么dp[i] = max(dp[j]+1, dp[i])
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})
	dp := make([]int, len(pairs))
	for i := range dp {
		dp[i] = 1
	}
	for i, pair := range pairs {
		for j := 0; j < i; j++ {
			if pair[0] > pairs[j][1] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
	}
	var result int
	for _, v := range dp {
		if result < v {
			result = v
		}
	}
	return result
}

// 最长摆动子序列
// 如果连续数字之间的差严格地在正数和负数之间交替，则数字序列称为 摆动序列 。第一个差（如果存在的话）可能是正数或负数。
// 仅有一个元素或者含两个不等元素的序列也视作摆动序列。
// 例如，[1, 7, 4, 9, 2, 5] 是一个 摆动序列 ，因为差值 (6, -3, 5, -7, 3)是正负交替出现的。
// 相反，[1, 4, 7, 2, 5]和[1, 7, 4, 5, 5] 不是摆动序列，第一个序列是因为它的前两个差值都是正数，第二个序列是因为它的最后一个差值为零。
// 子序列 可以通过从原始序列中删除一些（也可以不删除）元素来获得，剩下的元素保持其原始顺序。
// 给你一个整数数组 nums ，返回 nums 中作为 摆动序列 的 最长子序列的长度 。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/wiggle-subsequence
func wiggleMaxLength(nums []int) int {
	// 第i个元素的最长子序列是由i之前的能够形成摆动序列的那个元素决定的，但是两个元素不一定相邻
	// 对于i>j，如果nums[i]能够和nums[j]形成序列，那么dp[i] = dp[j]+1; 又由于0<=j<i，所以要找到dp[j]+1的最大值作为dp[i]
	// 摆动序列存在方向，即dp[i]需要一个元素标识当前方向，令dp[i][1]表示向上（比前一个大），dp[i][0]表示向下（比前一个小）
	dp := make([][2]int, len(nums))
	for i := 0; i < len(nums); i++ {
		tmp1, tmp2 := 0, 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && tmp1 < dp[j][1] {
				tmp1 = dp[j][1]
			}
			if nums[i] < nums[j] && tmp2 < dp[j][0] {
				tmp2 = dp[j][0]
			}
		}
		dp[i][0] = tmp1 + 1
		dp[i][1] = tmp2 + 1
	}
	var result int
	for _, v := range dp {
		if result < v[0] {
			result = v[0]
		}
		if result < v[1] {
			result = v[1]
		}
	}
	return result
}

//  最长公共子序列
// 给定两个字符串text1 和text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
// 一个字符串的子序列是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
// 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
// 两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/longest-common-subsequence
func longestCommonSubsequence(text1 string, text2 string) int {
	// dp[i,j]表示以text1第i个字符和text2第j个字符的公共子序列的长度
	// 如果text1[i]==text[j]，那么dp[i,j] = dp[i-1,j-1]+1
	// 如果text1[i] != text2[j], 那么dp[i,j] = max(dp[i,j-1], dp[i-1, j])
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}
	if len(text1) < len(text2) {
		text1, text2 = text2, text1
	}
	var dp [][]int
	for i := 0; i < len(text1); i++ {
		dp = append(dp, make([]int, len(text2)))
	}
	if text1[0] == text2[0] {
		dp[0][0] = 1
	}
	for i := 1; i < len(text1); i++ {
		if text1[i] == text2[0] {
			dp[i][0] = 1
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}
	for i := 1; i < len(text2); i++ {
		if text2[i] == text1[0] {
			dp[0][i] = 1
		} else {
			dp[0][i] = dp[0][i-1]
		}
	}

	for i := 1; i < len(text1); i++ {
		for j := 1; j < len(text2); j++ {
			if text1[i] == text2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
				continue
			}
			dp[i][j] = max(dp[i][j-1], dp[i-1][j])
		}
	}
	return dp[len(text1)-1][len(text2)-1]
}

// 删除两个字符串的字符使它们相等
// 给定两个单词word1和word2，找到使得word1和word2相同所需的最小步数，每步可以删除任意一个字符串中的一个字符。
// 提示：
//	给定单词的长度不超过500。
//	给定单词中的字符只含有小写字母。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/delete-operation-for-two-strings
func minDistance2(word1 string, word2 string) int {
	// 简化问题：删除两个字符串的字符使他们相等=》找到两个字符串的最大公共字串
	// 子问题：dp[i][j]对于字符串word1的前i个字符中，在word2前j个字符匹配的个数
	// 如果word1[i]等于word2[j]，则添加第j个字符，
	//		此时最大公共字串的个数为word1前i-1个字符匹配word2前j-1个字符的最大个数。那么dp[i][j]=dp[i-1][j-1]+1
	// 如果word1[i]不等于word2[j]，则不能添加第j个字符，此时最大公共字串的个数为word1的前i个字符匹配word2前j-1个字符的最大个数 或者
	//		word1前i-1个字符匹配word2前j个字符的最大个数，那么dp[i][j]=max(dp[i][j-1], dp[i-1][j])
	L1, L2 := len(word1), len(word2)
	dp := make([][]int, L1+1)
	for i := range dp {
		dp[i] = make([]int, L2+1)
	}
	for i := 1; i <= L1; i++ {
		for j := 1; j <= L2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	//for _, v := range dp {
	//	fmt.Println(v)
	//}
	return L1 + L2 - dp[L1][L2] - dp[L1][L2]
}

// 编辑距离
// 给你两个单词word1 和word2，请你计算出将word1转换成word2 所使用的最少操作数。
// 你可以对一个单词进行如下三种操作：
// 插入一个字符
// 删除一个字符
// 替换一个字符
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/edit-distance
func minDistance(word1 string, word2 string) int {
	// 子问题：word1前i个字符转换为word2前j个字符与之前的字符串有关
	// dp[i][j]: 对于word1前i个字符转换为word2前j个字符所使用的的最少操作数
	// 如果word1[i]==word2[j]，那么无需做转换操作, 即dp[i][j] = dp[i-1][j-1]
	// 如果word1[i]!=word2[j]
	// - 如果需要向word1插入字符：此时的操作数为前i个word1字符转换为前j-1个word2字符所需操作数+1 // abc => abcd
	// 		即 dp[i][j] = dp[i][j-1]+1
	// - 如果需要向word1删除字符：此时的操作数为前i-1个word1字符转换为前j个word2字符所需操作数+1 // abcd => abc
	// 		即 dp[i][j] = dp[i-1][j]+1
	// - 如果需要向word1替换字符：此时的操作数为前i-1个word1字符转换为前j-1个word2字符所需操作数+1 // abc => abd
	// 		即 dp[i][j] = dp[i-1][j-1]+1
	L1, L2 := len(word1), len(word2)
	dp := make([][]int, L1+1)
	// 初始化
	for i := range dp {
		dp[i] = make([]int, L2+1)
	}
	// 处理边界 i=0和j=0
	for i := 0; i <= L1; i++ {
		dp[i][0] = i // j为0时，word1转换为空字符串，都为删除操作
	}
	for i := 0; i <= L2; i++ {
		dp[0][i] = i // 当空字符串转换为word2时，都为插入操作
	}
	for i := 1; i <= L1; i++ { // word1转换为word2和word2转换为word1是相同的，所以两层循环的先后顺序可以互换
		for j := 1; j <= L2; j++ {
			if word1[i-1] != word2[j-1] { // i,j为word1和word2的索引，要减1
				dp[i][j] = min(dp[i][j-1]+1, min(dp[i-1][j]+1, dp[i-1][j-1]+1))
			} else {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[L1][L2]
}

// 复制粘贴字符
// 最初记事本上只有一个字符 'A' 。你每次可以对这个记事本进行两种操作：
// Copy All（复制全部）：复制这个记事本中的所有字符（不允许仅复制部分字符）。
// Paste（粘贴）：粘贴 上一次 复制的字符。
// 给你一个数字n ，你需要使用最少的操作次数，在记事本上输出 恰好n个 'A' 。返回能够打印出n个 'A' 的最少操作次数。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/2-keys-keyboard
func minSteps(n int) int {
	// 输出i个字符，找到整除数j，然后获取到获得j个字符的操作数dp[j]，然后进行复制-粘贴(粘贴次数为i/j-1),即dp[i] = dp[j] + i/j
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = math.MaxInt32
		for j := 1; j*j <= i; j++ {
			if i%j == 0 {
				dp[i] = min(dp[i], dp[j]+i/j)
				dp[i] = min(dp[i], dp[i/j]+j)
			}
		}
	}
	return dp[n]
}

func minSteps2(n int) int {
	// 质数分解
	// 复制-粘贴过程就是一个不断相乘的过程，即n=a*b*c*d...
	// 对于任意的数i，如果其为素数，那么无法继续进行拆分；如果为合数，那么对于任意一种拆分成两个大于1的整数的方式
	//	i = j1 *j2
	// 复制-粘贴的过程是相乘的过程，而代价最大为j1*j2，最小为j1+j2
	// 在数字大于1时，j1+j2<=j1*j2
	// 因此找到几个整数相乘等于n的过程中，最小代价找到素数之积为n，素数之和即为最小代价
	var result int
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			n /= i
			result += i
		}
	}
	if n > 1 { // 最后剩余的素数
		result += n
	}
	return result
}

// 最长有效括号
// 给你一个只包含 '('和 ')'的字符串，找出最长有效（格式正确且连续）括号子串的长度。
//示例 1：
// 	输入：s = ")()())"
//	输出：4
//	解释：最长有效括号子串是 "()()"
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-valid-parentheses
func longestValidParentheses(s string) int {
	// 对于一个字符串中的括号，左括号的数量不能比右括号的数量少，否则就要舍弃多余的左括号。
	// 所以问题可以转换：在目标字符串中str中，左括号和右括号的数量相等，并且对于任意前i个字符中，左括号要不少于右括号
	// 令dp[i][3] 为字符串s中前i个字符中左括号(dp[i][0])和右括号(dp[i][1])的数量,和有效括号的数量(dp[i][2])
	// - 如果第i-1个字符中左括号的数量大于右括号，那么对于第i个字符，
	//		- 如果为左括号，将第i个字符的左括号的数量+1，即
	//			dp[i][0] = dp[i-1][0]+1
	//			dp[i][1] = dp[i-1][1]
	//			此时有效括号的数量没有增加 dp[i][2] = 0
	//		- 如果为右括号，将第i个字符的右括号的数量+1，即
	//			dp[i][0] = dp[i-1][0]
	//			dp[i][1] = dp[i-1][1]+1
	//			此时有效括号的数量增加，其计算方式为dp[i-1]+1
	// - 如果第i-1个字符中左括号的数量等于右括号，那么对于第i个字符，
	// 		- 如果为左括号，将第i个字符的左括号的数量+1，即
	//			dp[i][0] = dp[i-1][0]+1
	//			dp[i][1] = dp[i-1][1]
	//		- 如果为右括号，则不能将第i个字符的长度计入。从i+1个字符重新开始计算，并将dp[i][0]比较-标记为临时最大值
	var (
		result int
		dp     = make([][2]int, len(s)+1)
	)
	for i := 1; i < len(s)+1; i++ {
		if dp[i-1][0] > dp[i-1][1] {
			if s[i-1] == '(' {
				dp[i][0] = dp[i-1][0] + 1
				dp[i][1] = dp[i-1][1]
			} else {
				dp[i][0] = dp[i-1][0]
				dp[i][1] = dp[i-1][1] + 1
			}
			continue
		}
		if dp[i-1][0] == dp[i-1][1] {
			if s[i-1] == '(' {
				dp[i][0] = dp[i-1][0] + 1
				dp[i][1] = dp[i-1][1]
			} else {
				result = max(result, dp[i-1][0])
			}
		}
	}
	result = max(result*2, dp[len(s)][1]*2)
	return result
}
