package classify

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

// 不同路径
// 一个机器人位于一个 m x n网格的左上角 （起始点在下图中标记为 “Start” ）。
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
// 问总共有多少条不同的路径？
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/unique-paths
func uniquePaths(m int, n int) int {
	// 找到“经验”：dp[i][j] = dp[i-1][j] +dp[i][j-1]
	var dp [][]int
	for i := 0; i <= m; i++ {
		dp = append(dp, make([]int, n+1))
	}

	var uniquePathsFib func(m int, n int) int
	uniquePathsFib = func(m int, n int) int {
		if m == 1 && n == 1 {
			return 1
		}
		if dp[m][n] > 0 {
			return dp[m][n]
		}
		if m == 1 {
			dp[m][n] = uniquePathsFib(m, n-1)
			return dp[m][n]
		}
		if n == 1 {
			dp[m][n] = uniquePathsFib(m-1, n)
			return dp[m][n]
		}
		dp[m][n] = uniquePathsFib(m-1, n) + uniquePathsFib(m, n-1)
		return dp[m][n]
	}
	return uniquePathsFib(m, n)
}

func uniquePathsNormal(m int, n int) int {
	// 找到“经验”：dp[i][j] = dp[i-1][j] +dp[i][j-1]
	var dp [][]int
	for i := 0; i < m; i++ {
		dp = append(dp, make([]int, n))
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// 空间复杂度优化 将二维转变为一维
// 将第一行初始化，计算第二行时，直接覆盖第一行的数据，即d[i] = d[i-1]+d[i]
func uniquePathsMemOpt(m int, n int) int {
	dp := make([]int, n)
	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			if j == 0 || i == 0 {
				dp[i] = 1
				continue
			}
			dp[i] += dp[i-1]
		}
	}
	return dp[n-1]
}

// 最小路径和
// 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
// 说明：每次只能向下或者向右移动一步。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/minimum-path-sum
func minPathSum(grid [][]int) int {
	// sum[i][j] = min(sum[i-1][j], sum[i][j-1]) + grid[i][j]
	if len(grid) == 0 {
		return 0
	}
	var sum [][]int
	for i := 0; i < len(grid); i++ {
		sum = append(sum, make([]int, len(grid[0])))
	}
	sum[0][0] = grid[0][0]
	for i := 1; i < len(grid); i++ {
		sum[i][0] = sum[i-1][0] + grid[i][0]
	}
	for i := 1; i < len(grid[0]); i++ {
		sum[0][i] = sum[0][i-1] + grid[0][i]
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[i]); j++ {
			sum[i][j] = min(sum[i-1][j], sum[i][j-1]) + grid[i][j]
		}
	}
	return sum[len(grid)-1][len(grid[0])-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
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
	min := func(a, b, c int) int {
		m := a
		if m > b {
			m = b
		}
		if m > c {
			m = c
		}
		return m
	}

	var dp [][]int // dp[i][j]表示word1中长度为i的字串变成word2中长度为j的字串需要的操作数
	for i := 0; i <= len(word1); i++ {
		dp = append(dp, make([]int, len(word2)+1))
	}
	for i := 1; i <= len(word1); i++ {
		// 长度为0的word2转换为长度为i的word1，只能插入或删除i次
		dp[i][0] = i
	}
	for j := 1; j <= len(word2); j++ {
		// 长度为0的word1转换为长度为j的word2，只能插入或者删除j次
		dp[0][j] = j
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			// 如果word1[i]与word2[j]相等，那么不需要任何操作，即dp[i][j] = dp[i-1][j-1]
			// 如果不等，则要进行三种操作：
			// - word1替换一个字符，即 dp[i][j] = dp[i-1][j-1] + 1
			// - word1插入一个字符，即 dp[i][j] = dp[i][j-1] + 1
			// - word1删除一个字符，即 dp[i][j] = dp[i-1][j] + 1
			if word1[i-1] == word2[j-1] { // 长度为i的word1所在的字符的索引为i-1
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
