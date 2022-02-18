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
