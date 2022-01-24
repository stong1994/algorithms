package math

// 相遇问题

// 最少移动次数使数组元素相等 II
// 给定一个非空整数数组，找到使所有数组元素相等所需的最小移动数，其中每次移动可将选定的一个元素加1或减1。
// 您可以假设数组的长度最多为10000。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements-ii
func minMoves2(nums []int) int {
	// 移动距离最小的方式是所有元素都移动到中位数：设 m 为中位数（总数若为奇数，选择中位数，若为偶数，则选择中间两个任意一个都可以）。
	// a 和 b 是 m 两边的两个元素，且 b > a。
	// 要使 a 和 b 相等，它们总共移动的次数为 b - a，此时在增加一个元素c，c大于a小于b，那么“目的值”为c时c不用移动，因此移动次数最少。
	// 1. 找到中位数，通过快排中寻找partition的方式来寻找中位数
	var result int
	val := findPartitionVal(nums)
	for _, v := range nums {
		diff := val - v
		if diff > 0 {
			result += diff
		} else {
			result -= diff
		}
	}
	return result
}

func findPartitionVal(nums []int) int {
	var (
		left  = 0
		right = len(nums) - 1
		mid   = len(nums) / 2
	)

	for {
		p := partition(nums, left, right)
		if p == mid {
			return nums[p]
		}
		if p > mid {
			right = p - 1
		} else {
			left = p + 1
		}
	}
}

func partition(nums []int, left, right int) int {
	var (
		p  = nums[left]
		lo = left
		hi = right + 1
	)
	for {
		for lo++; lo < right && nums[lo] <= p; lo++ {
		}
		for hi--; hi > left && nums[hi] >= p; hi-- {
		}
		if lo >= hi {
			break
		}
		nums[hi], nums[lo] = nums[lo], nums[hi]
	}
	nums[left], nums[hi] = nums[hi], nums[left]
	return hi
}
