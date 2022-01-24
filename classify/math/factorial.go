package math

// 阶乘

// 统计阶乘尾部有多少个 0
// 给定一个整数 n ，返回 n! 结果中尾随零的数量。
// 提示n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/factorial-trailing-zeroes
func trailingZeroes(n int) int {
	// 只有偶数乘以5才能产生0，于是尾随0的数量=max(5的数量，偶数的数量)，在阶乘中，偶数的数量一定大于5的数量，因此
	// 尾随0的数量就等于5的数量(准确的说是因子为5的数量，如25就能够和一个大于2的偶数形成两个0)
	var result int
	for i := 5; i <= n; i += 5 {
		for j := i; j > 1 && j%5 == 0; j /= 5 {
			result++
		}
	}
	return result
}

func trailingZeroes2(n int) int {
	// 对于一个数 N，它所包含 5 的个数为：N/5 + N/5^2 + N/5^3 + ...，其中 N/5 表示不大于 N 的数中 5 的倍数贡献一个 5，
	// N/5^2 表示不大于 N 的数中 5^2 的倍数再贡献一个 5 ...。
	if n == 0 {
		return 0
	}
	return n/5 + trailingZeroes2(n/5)
}
