package bit_operation

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
