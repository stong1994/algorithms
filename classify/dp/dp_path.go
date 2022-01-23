package dp

// 动态规划-矩阵路径问题
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
