package math

// 摩尔投票

// 多数元素
// 给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于n/2的元素。
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/majority-element
func majorityElement(nums []int) int {
	return moore2(nums)
}

func moore1(nums []int) int {
	// 简化问题，只能存在一种多数元素，并且中位数一定为多数元素，因此可转换为找找中位数
	return findPartitionVal(nums)
}

func moore2(nums []int) int {
	// 任意其他元素都能够被众数一一抵消。
	// 遍历nums，选择一个候选人，如果碰到一个同人，就cont++,如果碰到一个不同人，就count--，
	// 如果count等于0，重新选择候选人
	var (
		candidate int
		count     int
	)

	for _, v := range nums {
		if count == 0 {
			candidate = v
			count++
			continue
		}
		if candidate == v {
			count++
			continue
		}
		count--
	}
	return candidate
}
