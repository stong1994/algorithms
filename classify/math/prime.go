package math

// 素数相关

// 找到所有素数
// 统计所有小于非负整数 n 的质数的数量。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/count-primes
func countPrimes(n int) int {
	// 埃氏筛: 对于每个素数x，其倍数2x, 3x, 4x ... 一定不是素数
	var result int
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			result++
			for j := 2; j*i < n; j++ {
				isPrime[j*i] = false
			}
		}
	}
	return result
}
