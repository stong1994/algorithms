package dp

import (
	"sort"
)

// 我能赢吗
// 在 "100 game" 这个游戏中，两名玩家轮流选择从 1 到 10 的任意整数，累计整数和，先使得累计整数和 达到或超过 100 的玩家，即为胜者。
// 如果我们将游戏规则改为 “玩家 不能 重复使用整数” 呢？
// 例如，两个玩家可以轮流从公共整数池中抽取从 1 到 15 的整数（不放回），直到累计整数和 >= 100。
// 给定两个整数maxChoosableInteger（整数池中可选择的最大数）和desiredTotal（累计和），
// 若先出手的玩家是否能稳赢则返回 true，否则返回 false 。假设两位玩家游戏时都表现 最佳 。
// 提示:
// 1 <= maxChoosableInteger <= 20
// 0 <= desiredTotal <= 300
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/can-i-win
func canIWin2(maxChoosableInteger int, desiredTotal int) bool {
	// 如果desiredTotal不大于maxChoosableInteger，那么第一个玩家直接选择desiredTotal即可胜利
	if maxChoosableInteger >= desiredTotal {
		return true
	}
	// 如果全部总和不能达到desiredTotal，则返回false
	if (maxChoosableInteger*(maxChoosableInteger+1))/2 < desiredTotal {
		return false
	}
	// 递归+记录中间结果
	// 找到一个“路线”，不管另一名玩家如何选择，第一个玩家都能直接获胜
	// 需要知道哪些值已经选择，设已选择的值为selected，用20个bit即可
	// 需要遍历所有“路线”来找到目标“路线”,即第一个玩家先从1开pick，获取第二个玩家pick剩下所有可能的结果，如果达到desiredTotal，则此路线不通
	// 否则，第二个玩家再次pick。用pick的次数cnt即可知道是哪个玩家
	storeSelect := make(map[uint32]bool)
	var isValid func(cnt int, selected uint32, sum int) bool
	isValid = func(cnt int, selected uint32, sum int) bool {
		if cnt%2 == 0 { // 第一个玩家选择
			for pick := 1; pick <= maxChoosableInteger; pick++ {
				if selected&(1<<(pick-1)) == 0 {
					pickVal := sum + pick
					newSelect := selected | 1<<(pick-1)
					if pickVal >= desiredTotal {
						return true
					}
					if valid, exist := storeSelect[newSelect]; exist {
						if valid {
							return true
						}
					} else {
						valid = isValid(cnt+1, newSelect, pickVal)
						storeSelect[newSelect] = valid
						if valid {
							return true
						}
					}
				}
			}
			return false
		}
		// 第二个玩家选择
		for pick := 1; pick <= maxChoosableInteger; pick++ {
			if selected&(1<<(pick-1)) == 0 {
				pickVal := sum + pick
				newSelect := selected | 1<<(pick-1)
				if pickVal >= desiredTotal {
					return false
				}
				if valid, exist := storeSelect[newSelect]; exist {
					if !valid {
						return false
					}
				} else {
					valid = isValid(cnt+1, newSelect, pickVal)
					storeSelect[newSelect] = valid
					if !valid {
						return false
					}
				}
			}
		}
		return true
	}

	for pick := 1; pick <= maxChoosableInteger; pick++ {
		if isValid(1, 1<<(pick-1), pick) {
			return true
		}
	}
	return false
}

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if maxChoosableInteger*(maxChoosableInteger+1)/2 < desiredTotal {
		return false
	}
	// dp数组: key——选择的方式 value——0-未被使用 1-被第一个人选择后大于目标值 2-被选中，但和未大于目标值
	var find func(dp []byte, v, target int) bool
	find = func(dp []byte, v, target int) bool {
		if dp[v] != 0 {
			return dp[v] == 1
		}
		dp[v] = 2
		for i := 1; i <= maxChoosableInteger; i++ {
			selectNum := 1 << (i - 1)
			if v&selectNum > 0 { //i 已被选择
				continue
			}
			if i >= target {
				dp[v] = 1
				break
			}
			next := find(dp, v|selectNum, target-i)
			if !next {
				dp[v] = 1
				break
			}
		}
		return dp[v] == 1
	}
	dp := make([]byte, 1<<maxChoosableInteger-1) // 对于maxChoosableInteger个值，有1<<maxChoosableInteger-1种选择方式
	return find(dp, 0, desiredTotal)
}

// 火柴拼正方形
// 你将得到一个整数数组 matchsticks ，其中 matchsticks[i] 是第 i个火柴棒的长度。
// 你要用 所有的火柴棍拼成一个正方形。你 不能折断 任何一根火柴棒，但你可以把它们连在一起，而且每根火柴棒必须 使用一次 。
// 如果你能使这个正方形，则返回 true ，否则返回 false 。
// 提示:
//	1 <= matchsticks.length <= 15
//	1 <= matchsticks[i] <= 108
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/matchsticks-to-square
func makesquare(matchsticks []int) bool {
	// 将matchsticks分成四份，每份的和相等
	// 暴力破解
	sum := 0
	for _, v := range matchsticks {
		sum += v
	}
	if sum%4 != 0 {
		return false
	}
	sideNum := sum / 4
	sides := make([]int, 4)
	var find func(selected uint16) bool
	find = func(selected uint16) bool {
		if sides[0] == sides[1] && sides[0] == sides[2] && sides[0] == sides[3] && sides[0] == sideNum {
			return true
		}
		for i, stick := range matchsticks {
			if selected&(1<<i) != 0 {
				continue
			}
			for j, side := range sides {
				if side+stick > sideNum {
					continue
				}
				sides[j] = side + stick
				if find(selected | 1<<i) {
					return true
				}
				sides[j] = side
			}
		}
		return false
	}
	return find(0)
}

// 每次都要循环遍历matchsticks，可以去掉。每次都给一根火柴找位置
func makesquare2(matchsticks []int) bool {
	// 将matchsticks分成四份，每份的和相等
	// 暴力破解
	sum := 0
	for _, v := range matchsticks {
		sum += v
	}
	if sum%4 != 0 {
		return false
	}
	sideNum := sum / 4
	sides := make([]int, 4)
	var find func(idx int) bool
	find = func(idx int) bool {
		if idx == len(matchsticks) {
			if sides[0] == sides[1] && sides[0] == sides[2] && sides[0] == sides[3] && sides[0] == sideNum {
				return true
			}
			return false
		}

		for j, side := range sides {
			if side+matchsticks[idx] > sideNum {
				continue
			}
			sides[j] = side + matchsticks[idx]
			if find(idx + 1) {
				return true
			}
			sides[j] = side
		}
		return false
	}
	return find(0)
}

// 如果数组前边的火柴长度较短，那么就需要多次递归，对其排序
func makesquare3(matchsticks []int) bool {
	sort.Slice(matchsticks, func(i, j int) bool {
		return matchsticks[i] > matchsticks[j]
	})
	sum := 0
	for _, v := range matchsticks {
		sum += v
	}
	if sum%4 != 0 {
		return false
	}
	sideNum := sum / 4
	sides := make([]int, 4)
	var find func(idx int) bool
	find = func(idx int) bool {
		if idx == len(matchsticks) {
			if sides[0] == sides[1] && sides[0] == sides[2] && sides[0] == sides[3] && sides[0] == sideNum {
				return true
			}
			return false
		}

		for j, side := range sides {
			if side+matchsticks[idx] > sideNum {
				continue
			}
			sides[j] = side + matchsticks[idx]
			if find(idx + 1) {
				return true
			}
			sides[j] = side
		}
		return false
	}
	return find(0)
}

// 优化：如果后边的边和前边的边相同，前边的没成功，后边的也不会成功
func makesquare4(matchsticks []int) bool {
	sort.Slice(matchsticks, func(i, j int) bool {
		return matchsticks[i] > matchsticks[j]
	})
	sum := 0
	for _, v := range matchsticks {
		sum += v
	}
	if sum%4 != 0 {
		return false
	}
	sideNum := sum / 4
	sides := make([]int, 4)
	var find func(idx int) bool
	find = func(idx int) bool {
		if idx == len(matchsticks) {
			if sides[0] == sides[1] && sides[0] == sides[2] && sides[0] == sideNum {
				return true
			}
			return false
		}

		for j, side := range sides {
			if side+matchsticks[idx] > sideNum || (j > 0 && sides[j-1] == sides[j]) || (idx == len(matchsticks)-1 && j != 0) {
				continue
			}
			sides[j] = side + matchsticks[idx]
			if find(idx + 1) {
				return true
			}
			sides[j] = side
		}
		return false
	}
	return find(0)
}

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
	// 	f[i][mask] = f[i-1][mask] +f[i-1][mask\subset]*freq[i]
	// mask\subset 表示从二进制表示 mask 中去除所有在subset 中出现的1，可以使用按位异或运算实现。
	// 这里需要保证 subset 是 mask 的子集，可以使用按位与运算来判断。
	var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	var ans int
	const mod int = 1e9 + 7 // 防止整数和溢出
	freq := [31]int{}
	for _, num := range nums {
		freq[num]++
	}

	f := make([]int, 1<<len(primes))
	f[0] = 1
	for i := 0; i < freq[1]; i++ { // nums中有n个1，最终结果就要乘以n个2
		f[0] = f[0] * 2 % mod
	}
next:
	for i := 2; i < 31; i++ {
		if freq[i] == 0 {
			continue
		}

		// 检查 i 的每个质因数是否均不超过 1 个，将乘积是i的不同质因数对应的位存入subset中
		subset := 0
		for j, prime := range primes {
			if i%(prime*prime) == 0 { // 如果i由两个以上相同的质因数组成，则不满足条件
				continue next
			}
			if i%prime == 0 { // 如果i是质数，则将对应的质数位存入subset中
				subset |= 1 << j
			}
		}

		// 动态规划
		for mask := 1 << len(primes); mask > 0; mask-- {
			if mask&subset == subset { // 说明subset是mask的子集
				f[mask] = (f[mask] + f[mask^subset]*freq[i]) % mod // 通过异或找到subset中未使用的位数
			}
		}
	}

	for _, v := range f[1:] {
		ans = (ans + v) % mod
	}
	return ans
}
