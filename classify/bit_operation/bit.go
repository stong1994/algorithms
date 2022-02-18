package bit_operation

/* 位操作
1. a^b^^c 如果三个数中有两个数相等，那么通过异或操作可以找到三个数中不同的那个数
2. b := a & -a 找到a中最低位为1的那位对应的值
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
