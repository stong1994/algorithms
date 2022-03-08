package stack_queue

// 子数组范围和
// 给你一个整数数组 nums 。nums 中，子数组的 范围 是子数组中最大元素和最小元素的差值。
// 返回 nums 中 所有 子数组范围的 和 。
// 子数组是数组中一个连续 非空 的元素序列。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/sum-of-subarray-ranges
func subArrayRanges(nums []int) int64 {
	/** 暴力破解
	n := len(nums)
	result := int64(0)
	for i := 0; i < n; i++ {
		maxV, minV := nums[i], nums[i]
		for j := i+1; j < n; j++ {
			maxV = max(maxV, nums[j])
			minV = min(minV, nums[j])
			result += int64(maxV - minV)
		}
	}
	return result
	*/
	// 结果值为“最大值”累加值减去“最小值”的累加值，重点在于找到“最大值”和“最小值”
	// 对于“最大值”（索引为i），需要找到其为最大值的区间。在区间l,j中这个值为最大值，那么l,j之间存在的序列的个数为(i-l+1)*(j-i+1)
	// “最小值”同理
	// 可以通过单调栈可以找到找到一个最值的区间的一个边界：如对于“最大值”，从左向右遍历nums，每次将num和栈顶元素比较，
	// 如果栈顶元素大于num，则栈顶元素对应的索引就是num作为“最大值”的左边界。
	// 于是通过四次遍历即可获得两个最值的四个边界
	n := len(nums)
	minStk := make([]int, 0)        // 栈顶元素不小于栈中其他元素，在遍历过程中，如果当前元素不大于栈顶，说明该元素就是当前最小值，否则，将栈顶元素移除，直到找到小于该元素的栈顶
	minBorderLeft := make([]int, n) // 第i个元素作为最小值的左边界（左边界是前i-1个元素中比第i个元素小的元素处的索引）
	maxStk := make([]int, 0)
	maxBorderLeft := make([]int, n)
	for i, num := range nums {
		for len(minStk) > 0 && nums[minStk[len(minStk)-1]] > num {
			minStk = minStk[:len(minStk)-1]
		}
		if len(minStk) > 0 {
			minBorderLeft[i] = minStk[len(minStk)-1]
		} else {
			minBorderLeft[i] = -1
		}
		minStk = append(minStk, i)

		for len(maxStk) > 0 && nums[maxStk[len(maxStk)-1]] <= num {
			maxStk = maxStk[:len(maxStk)-1]
		}
		if len(maxStk) > 0 {
			maxBorderLeft[i] = maxStk[len(maxStk)-1]
		} else {
			maxBorderLeft[i] = -1
		}
		maxStk = append(maxStk, i)
	}

	minStk = minStk[:0]
	minBorderRight := make([]int, n)
	maxStk = maxStk[:0]
	maxBorderRight := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for len(minStk) > 0 && nums[minStk[len(minStk)-1]] >= nums[i] {
			minStk = minStk[:len(minStk)-1]
		}
		if len(minStk) > 0 {
			minBorderRight[i] = minStk[len(minStk)-1]
		} else {
			minBorderRight[i] = n
		}
		minStk = append(minStk, i)

		for len(maxStk) > 0 && nums[maxStk[len(maxStk)-1]] < nums[i] {
			maxStk = maxStk[:len(maxStk)-1]
		}
		if len(maxStk) > 0 {
			maxBorderRight[i] = maxStk[len(maxStk)-1]
		} else {
			maxBorderRight[i] = n
		}
		maxStk = append(maxStk, i)
	}

	result := int64(0)
	for i := 0; i < n; i++ {
		result += int64((maxBorderRight[i]-i)*(i-maxBorderLeft[i])-(minBorderRight[i]-i)*(i-minBorderLeft[i])) * int64(nums[i])
	}
	return result
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
