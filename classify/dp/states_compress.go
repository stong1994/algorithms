package dp

// 好子集的数目
// 给你一个整数数组nums。如果nums的一个子集中，所有元素的乘积可以表示为一个或多个 互不相同的质数 的乘积，那么我们称它为好子集。
//比方说，如果nums = [1, 2, 3, 4]：
//[2, 3]，[1, 2, 3]和[1, 3]是 好子集，乘积分别为6 = 2*3，6 = 2*3和3 = 3。
//[1, 4] 和[4]不是 好子集，因为乘积分别为4 = 2*2 和4 = 2*2。
//请你返回 nums中不同的好子集的数目对10^9 + 7取余的结果。
//nums中的 子集是通过删除 nums中一些（可能一个都不删除，也可能全部都删除）元素后剩余元素组成的数组。
//如果两个子集删除的下标不同，那么它们被视为不同的子集。
// 提示：
// 1 <= nums.length <= 105
// 1 <= nums[i] <= 30
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/the-number-of-good-subsets
func numberOfGoodSubsets(nums []int) int {
	// 暴力破解：首先要找到所有子集，再判断每个子集的乘积是否是多个互不相同的质数的乘积
	// 状态压缩动态规划:
	// 观察到num不小于1，不大于30，因此可以将数据分为三类
	// - 1: 1与任何子集都能够组合，因此如果nums中包含1，则符合条件的子集的数量要加倍
	// - 2,3,5,6,7,10,11,13,14,15,17,19,21,22,23,26,29,30:不包含平方因子，每个数在子集中能出现一次
	// - 4,8,9,12,16,18,20,24,25,27,28: 包含平方因子，每个数字不能在子集中出现
	// 1到30中一共有10个质数，可以用10个bit来表示，设为mask
	// f[i][mask]表示当我们选择[2,i]的范围内的数，并且质数的使用情况为mask时的方案数
	// 如果i是上述第三类数据，我们无法选择，因此 f[i][mask] = f[i-1][mask]
	// 如果i是上述第二类数据，设其包含的质因子的二进制表示为subset，i在nums中出现的次数为freq[i].得
	// 	f[i][mask] = f[i-1][mask\subset]*freq[i]
	// mask\subset 表示从二进制表示 mask 中去除所有在subset 中出现的1，可以使用按位异或运算实现。
	// 这里需要保证 subset 是 mask 的子集，可以使用按位与运算来判断。
	var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	var ans int
	const mod int = 1e9 + 7
	freq := [31]int{}
	for _, num := range nums {
		freq[num]++
	}

	f := make([]int, 1<<len(primes))
	f[0] = 1
	for i := 0; i < freq[1]; i++ {
		f[0] = f[0] * 2 % mod
	}
next:
	for i := 2; i < 31; i++ {
		if freq[i] == 0 {
			continue
		}

		// 检查 i 的每个质因数是否均不超过 1 个
		subset := 0
		for j, prime := range primes {
			if i%(prime*prime) == 0 {
				continue next
			}
			if i%prime == 0 {
				subset |= 1 << j
			}
		}

		// 动态规划
		for mask := 1 << len(primes); mask > 0; mask-- {
			if mask&subset == subset {
				f[mask] = (f[mask] + f[mask^subset]*freq[i]) % mod
			}
		}
	}

	for _, v := range f[1:] {
		ans = (ans + v) % mod
	}
	return ans
}
