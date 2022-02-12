package hash

// 数组中两个数的和为给定值
// 给定一个整数数组 nums和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那两个整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
// 你可以按任意顺序返回答案。只会存在一个有效答案
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/two-sum
func twoSum(nums []int, target int) []int {
	h := make(map[int]int, len(nums))
	for i, num := range nums {
		if j, exist := h[target-num]; exist {
			return []int{i, j}
		}
		h[num] = i
	}
	return nil
}

// 判断数组是否含有重复元素
// 给你一个整数数组 nums 。如果任一值在数组中出现 至少两次 ，返回 true ；如果数组中每个元素互不相同，返回 false 。
func containsDuplicate(nums []int) bool {
	h := make(map[int]bool, len(nums))
	for _, num := range nums {
		if h[num] {
			return true
		}
		h[num] = true
	}
	return false
}

// 最长和谐序列
// 和谐数组是指一个数组里元素的最大值和最小值之间的差别 正好是 1 。
// 现在，给你一个整数数组 nums ，请你在所有可能的子序列中找到最长的和谐子序列的长度。
// 数组的子序列是一个由数组派生出来的序列，它可以通过删除一些元素或不删除元素、且不改变其余元素的顺序而得到。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/longest-harmonious-subsequence
func findLHS(nums []int) int {
	h := make(map[int]int, len(nums))
	for _, num := range nums {
		h[num]++
	}
	var result int
	for num, cnt := range h {
		if h[num+1] > 0 {
			result = max(result, cnt+h[num+1])
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

// 最长连续序列
// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
// 请你设计并实现时间复杂度为O(n) 的算法解决此问题。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/longest-consecutive-sequence
func longestConsecutive(nums []int) int {
	// 第一次遍历nums，对于第i个元素num放入hash中
	// 第二次遍历nums，对于第i个元素num，在hash中找到是否存在num-1或者num+1，如果存在，将当前长度加1，并在hash中标记该元素已使用
	var result int
	hash := make(map[int]bool, len(nums))
	for _, num := range nums {
		hash[num] = false
	}
	var getCnt func(num int) int
	getCnt = func(num int) int {
		used, exist := hash[num]
		if !exist || used {
			return 0
		}
		hash[num] = true
		return 1 + getCnt(num-1) + getCnt(num+1)
	}
	for _, num := range nums {
		result = max(result, getCnt(num))
	}
	return result
}

func longestConsecutive2(nums []int) int {
	// 对于连续序列 x, x+1, x+2,,, x+y中的每个元素z ，为了避免重复计算，只考虑是否存在z+1，如果存在，hash(z+1) = hash(z)+1，
	// 如果存在z-1，那么忽略即可（因为在遍历z-1时，就会计算cnt(z) = hash(z-1)+1）
	hash := make(map[int]bool, len(nums))
	for _, num := range nums {
		hash[num] = true
	}
	var result int
	for num := range hash {
		if !hash[num-1] {
			cur := 1
			for hash[num+1] {
				cur++
				num++
			}
			result = max(result, cur)
		}
	}
	return result
}
