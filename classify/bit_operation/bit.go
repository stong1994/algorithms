package bit_operation

import "sort"

/* 位操作
1. a^b^c 如果三个数中有两个数相等，那么通过异或操作可以找到三个数中不同的那个数
2. a & -a 找到a中最低位为1的那位对应的值
3. n & (n - 1) 将n中的最低位1移除
*/

// 统计两个数的二进制表示有多少位不同
// https://leetcode-cn.com/problems/hamming-distance/
func hammingDistance(x int, y int) int {
	// 如果相同的话，需要为0，否则为1 =》 异或操作
	x = x ^ y
	bit := 31
	cnt := 0
	for bit >= 0 {
		if x>>bit&1 == 1 {
			cnt++
		}
		bit--

	}
	return cnt
}

// 数组中唯一一个不重复的元素
// https://leetcode-cn.com/problems/single-number/description/
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
// 说明：
// 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/single-number
func singleNumber(nums []int) int {
	// 计数可用哈希表，但是题目要求不使用额外的空间
	// 在三个数a^b^c中，如果其中两个数相等，那么结果为另一个数
	for i := 1; i < len(nums); i++ {
		nums[i] ^= nums[i-1]
	}
	return nums[len(nums)-1]
}

// 找出数组中缺失的那个数
// https://leetcode-cn.com/problems/missing-number/
func missingNumber(nums []int) int {
	/*
		n := len(nums)
		sum := 0
		total := n*(n+1)/2
		for _, v := range nums {
		    sum += v
		}
		return total - sum
	*/
	// 将该题转换为《数组中唯一不重复的元素》，即在nums后边append 0~n+1即可
	n := len(nums)
	for i := 0; i < n+1; i++ {
		nums = append(nums, i)
	}
	for i := 1; i < len(nums); i++ {
		nums[i] ^= nums[i-1]
	}
	return nums[len(nums)-1]
}

// 数组中不重复的两个元素
// 给定一个整数数组nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。你可以按 任意顺序 返回答案。
//进阶：你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现？
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/single-number-iii
func singleNumber2(nums []int) []int {
	xorSum := 0
	for _, num := range nums {
		xorSum ^= num
	}
	lsb := xorSum & -xorSum // 取出xorSum二进制中最右边为1的那位，对于异或结果为xorSum的两个数x,y，如果第l位为1，则说明一个数的第l位为1，另一个为0
	// 将nums分别两类，一类是第l位为1，另一个为0，分别进行异或操作即得两个目标数
	rst1, rst2 := 0, 0
	for _, num := range nums {
		if num&lsb > 0 {
			rst1 ^= num
		} else {
			rst2 ^= num
		}
	}
	return []int{rst1, rst2}
}

// 翻转一个数的比特位
// num 为一个长为32 的二进制串
// https://leetcode-cn.com/problems/reverse-bits/
func reverseBits(num uint32) uint32 {
	/* 从低位到高位遍历num，然后将其颠倒存到结果中
	result := uint32(0)
	for i := 0; i < 32 && num > 0; i++ {
		result |= (num & 1) << (31-i)
		num >>= 1
	}
	return result
	*/
	// 分治
	// 翻转一个字符串可以先将两边的子字符串翻转，在翻转整个字符串，通过位运算也能实现翻转比特位。
	const (
		m1 = 0x55555555 // 01010101010101010101010101010101
		m2 = 0x33333333 // 00110011001100110011001100110011
		m4 = 0x0f0f0f0f // 00001111000011110000111100001111
		m8 = 0x00ff00ff // 00000000111111110000000011111111
	)
	num = num>>1&m1 | num&m1<<1 // num>>1&m1 将偶数位放到奇数位  num&m1<<1 将奇数位放到偶数位 =》实现奇偶互换
	num = num>>2&m2 | num&m2<<2 // 将相邻的2位两两互换
	num = num>>4&m4 | num&m4<<4 // 将相邻的4位两两互换
	num = num>>8&m8 | num&m8<<8 // 将相邻的8位两两互换
	return num>>16 | num<<16    // // 将相邻的16位两两互换
}

// 不用额外变量交换两个整数
// 程序员代码面试指南 ：P317
func swap(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b // b = a^b^b
	a = a ^ b // a = a^a^b
	return a, b
}

// 判断一个数是不是 2 的 n 次方
// https://leetcode-cn.com/problems/power-of-two/
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	// 2的幂次增加的过程就是向左移位的过程
	return n&-n == n
}

// 判断一个数是不是 4 的 n 次方
// https://leetcode-cn.com/problems/power-of-four/
func isPowerOfFour(n int) bool {
	// 先判断n是2的n次方，再判断1是否存在奇数为即可
	return n > 0 && n&(n-1) == 0 && n&0xaaaaaaaa == 0
}

// 判断一个数的位级表示是否不会出现连续的 0 和 1
// https://leetcode-cn.com/problems/binary-number-with-alternating-bits/
func hasAlternatingBits(n int) bool {
	// 二进制表示中相邻两位是否永远不同
	// 如果相邻两位永远不同，那么右移1位后与原数据进行异或操作，的得到全是1，此时加1后得到的位中只有一个1，通过 num & (num-1) == 0 来判断
	num := n ^ (n >> 1) + 1
	return num&(num-1) == 0
}

// 求一个数的补码
// 对整数的二进制表示取反（0 变 1 ，1 变 0）后，再转换为十进制表示，可以得到这个整数的补数。
//例如，整数 5 的二进制表示是 "101" ，取反后得到 "010" ，再转回十进制表示得到补数 2 。
//给你一个整数 num ，输出它的补数。
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/number-complement
func findComplement(num int) int {
	// 对num进行取反，但要排除掉前置1,因此我们可以直接补充前置1，这样取反后就是0
	n := uint32(num)
	bit := 31
	for 1<<bit&n == 0 {
		n |= 1 << bit
		bit--
	}
	return int(^n)
}

// 实现整数的加法
// 给你两个整数 a 和 b ，不使用 运算符 + 和 -，计算并返回两整数之和。
// https://leetcode-cn.com/problems/sum-of-two-integers
func getSum(a int, b int) int {
	// 对于a和b之和，每位的值为a和b不进位的加法 加上 a和b进位的加法。
	// 当a和b不进位时，和为a^b
	// 当a和b做进位时，值为 a&b << 1
	// https://leetcode-cn.com/problems/sum-of-two-integers/solution/liang-zheng-shu-zhi-he-wei-yun-suan-qing-7095/
	if a == 0 {
		return b
	}
	carry := uint(a&b) << 1
	sum := a ^ b
	return getSum(int(carry), sum)
}

// 字符串数组最大乘积
// 给你一个字符串数组words ，找出并返回 length(words[i]) * length(words[j])的最大值，并且这两个单词不含有公共字母。如果不存在这样的两个单词，返回 0 。
// words[i] 仅包含小写字母
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/maximum-product-of-word-lengths
func maxProduct(words []string) int {
	//需要注意两点
	//1. 两个单词不包含公共字母: words中只包含小写字母，因此可用26个bit来表示每个字符串，没有公共bit，因此两个bit值的与操作结果应为0( b1 & b2 == 0)
	//2. 最大长度: 暴力破解，两两比较取最大值
	n := len(words)
	toBits := func(word string) uint32 {
		rst := uint32(0)
		for _, w := range word {
			rst |= 1 << (w - 'a')
		}
		return rst
	}
	lengths := make([]int, n)
	bitsArr := make([]uint32, n)
	for i := 0; i < n; i++ {
		bitsArr[i] = toBits(words[i])
		lengths[i] = len(words[i])
	}
	result := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if bitsArr[i]&bitsArr[j] == 0 {
				result = max(result, lengths[i]*lengths[j])
			}
		}
	}
	return result
}

func maxProduct2(words []string) int {
	// 需要注意两点
	//1. 两个单词不包含公共字母: words中只包含小写字母，因此可用26个bit来表示每个字符串，没有公共bit，因此两个bit值的与操作结果应为0( b1 & b2 == 0)
	//2. 最大长度: 暴力破解，两两比较取最大值
	// 优化：对于meet和met，具有相同的字符，因此只需要记录最长长度即可。在上个解法中，使用数组存储这两个单词对应的bits，但实际上这两个bits相同，因此可采用哈希表来去重
	n := len(words)
	toBits := func(word string) uint32 {
		rst := uint32(0)
		for _, w := range word {
			rst |= 1 << (w - 'a')
		}
		return rst
	}
	bitsMap := make(map[uint32]int, n)
	for i := 0; i < n; i++ {
		bitsMap[toBits(words[i])] = max(bitsMap[toBits(words[i])], len(words[i]))
	}
	result := 0
	// 用map遍历相同的两个bits要比较两次，因此用map虽然节省了空间，但是浪费了时间
	for b1, l1 := range bitsMap {
		for b2, l2 := range bitsMap {
			if b1&b2 == 0 {
				result = max(result, l1*l2)
			}
		}
	}
	return result
}

func maxProduct3(words []string) int {
	//需要注意两点
	//1. 两个单词不包含公共字母: words中只包含小写字母，因此可用26个bit来表示每个字符串，没有公共bit，因此两个bit值的与操作结果应为0( b1 & b2 == 0)
	//2. 最大长度: 暴力破解，两两比较取最大值
	// 最后比较的时候需要对全量数据两两比较。先对words按长度进行排序，那么最后比较的时候，对每个bits，先用长度大的去比较，得到值就是当前bits能够得到的最大值
	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) < len(words[j]) {
			return false
		}
		return true
	})
	n := len(words)
	toBits := func(word string) uint32 {
		rst := uint32(0)
		for _, w := range word {
			rst |= 1 << (w - 'a')
		}
		return rst
	}
	lengths := make([]int, n)
	bitsArr := make([]uint32, n)
	for i := 0; i < n; i++ {
		bitsArr[i] = toBits(words[i])
		lengths[i] = len(words[i])
	}
	result := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if bitsArr[i]&bitsArr[j] == 0 {
				result = max(result, lengths[i]*lengths[j])
				break
			}
		}
	}
	return result
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 统计从 0 ~ n 每个数的二进制表示中 1 的个数
// 给你一个整数 n ，对于 0 <= i <= n 中的每个 i ，计算其二进制表示中 1 的个数 ，返回一个长度为 n + 1 的数组 ans 作为答案。
// https://leetcode-cn.com/problems/counting-bits/
func countBits(n int) []int {
	// 对于数字6(110)可看做是4(100) + 2(10)的结果，类似于dp[i] = dp[i-1] + 1
	// 关键点在于如何把最高位干掉（110 =》 10），利用位操作(i&i-1)
	result := make([]int, n+1)
	for i := 1; i <= n; i++ {
		result[i] = result[i&(i-1)] + 1
	}
	return result
}
