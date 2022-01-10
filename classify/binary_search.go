package classify

import (
	"math"
)

// 二分查找

// 给你一个非负整数 x ，计算并返回x的 算术平方根 。
// 由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
// 注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/sqrtx
func mySqrt(x int) int {
	// 基本思路：找到一个数，其平方值小于等于x，而大于该数的平方值大于x
	var result int
	lo, hi := 0, x
	for lo <= hi {
		mid := lo + (hi-lo)/2 // 防止整数溢出
		if mid*mid <= x {
			result = mid
			lo = mid + 1
			continue
		}
		hi = mid - 1
	}
	return result
}

// 大于给定元素的最小元素
// 给你一个排序后的字符列表 letters ，列表中只包含小写英文字母。另给出一个目标字母target，请你寻找在这一有序列表里比目标字母大的最小字母。
// 在比较时，字母是依序循环出现的。举个例子：
// 如果目标字母 target = 'z' 并且字符列表为letters = ['a', 'b']，则答案返回'a'
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/find-smallest-letter-greater-than-target
func nextGreatestLetter(letters []byte, target byte) byte {
	// 找到letters中比target大的最小的字符，如果不存在，就返回letters中最小的字符
	globalMin := byte(math.MaxUint8)
	partMin := byte(math.MaxUint8)
	for _, v := range letters {
		if v < globalMin {
			globalMin = v
		}
		if v > target && v < partMin {
			partMin = v
		}
	}
	if partMin == byte(math.MaxUint8) {
		return globalMin
	}
	return partMin
}

// 注意到letters是排好序的，所以能否有优化算法。
// 通过二分查找来进行比较
func nextGreatestLetterOpt(letters []byte, target byte) byte {
	if len(letters) == 0 {
		return 0
	}
	lo, hi := 0, len(letters)-1
	result := letters[0]
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if letters[mid] > target {
			result = letters[mid]
			hi = mid - 1
			continue
		}
		lo = mid + 1
	}
	return result
}

// 有序数组的 Single Element
// 给你一个仅由整数组成的有序数组，其中每个元素都会出现两次，唯有一个数只会出现一次。
// 请你找出并返回只出现一次的那个数。
// 你设计的解决方案必须满足 O(log n) 时间复杂度和 O(1) 空间复杂度。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/single-element-in-a-sorted-array
func singleNonDuplicate(nums []int) int {
	// 如果没有复杂度方面的要求，那么直接遍历即可。难点在于时间复杂度为O(log n)。
	// 通过时间复杂度的要求，能够猜测出应该是“二分”思想。即通过中间的两个元素来判断出所求元素的位置。
	// 通过观察数据发现，nums的个数一定是单数
	// 获取中间的元素，如果与两边的元素都不同，则返回中间的元素，否则向两边进行”二分“
	// 如果中间元素与左边的元素相同，那么判断左边的元素个数，如果是奇数，则单元素在右边，如果是偶数，则单元素在左边
	// 如果中间元素与右边的元素相同，那么判断右边的元素个数，如果是奇数，则单元素在左边，如果是偶数，则单元素在右边
	lo, hi := 0, len(nums)-1
	for {
		if hi == lo {
			return nums[lo]
		}
		if hi-lo == 1 {
			if lo > 0 {
				if nums[lo-1] == nums[lo] {
					return nums[hi]
				}
				return nums[lo]
			}
			if nums[hi] == nums[hi+1] {
				return nums[lo]
			}
			return nums[hi]
		}
		if hi-lo == 2 {
			if nums[hi] == nums[hi-1] {
				return nums[lo]
			}
			return nums[hi]
		}

		mid := lo + (hi-lo)/2
		if nums[mid] == nums[mid-1] {
			if (mid-lo)&1 == 1 { // 左边为奇数，则单元素在右边
				lo = mid + 1
			} else {
				hi = mid - 2 // 注意：这里需要移动两位
			}
			continue
		}
		if nums[mid] == nums[mid+1] {
			if (hi-mid)&1 == 1 { // 右边为奇数，则单元素在左边
				hi = mid - 1
			} else {
				lo = mid + 2 // 注意：这里需要移动两位
			}
			continue
		}
		return nums[mid]
	}
}

// nums的长度一定是奇数，且正常情况下偶数位与偶数位+1的元素应该是相同的。如果不同，说明偶数位以下含有单个元素
func singleNonDuplicateOpt(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if mid&1 == 1 {
			mid-- // mid如果是奇数，则--
		}
		if nums[mid] == nums[mid+1] {
			lo = mid + 2 // 注意：这里移动两位
			continue
		}
		hi = mid
	}
	return nums[lo]
}

// 第一个错误的版本
// 你是产品经理，目前正在带领一个团队开发新的产品。不幸的是，你的产品的最新版本没有通过质量检测。
// 由于每个版本都是基于之前的版本开发的，所以错误的版本之后的所有版本都是错的。
// 假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。
// 你可以通过调用bool isBadVersion(version)接口来判断版本号 version 是否在单元测试中出错。
// 实现一个函数来查找第一个错误的版本。你应该尽量减少对调用 API 的次数。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/first-bad-version
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
//func firstBadVersion(n int) int {
//	return sort.Search(n, func(version int) bool {
//		return isBadVersion(version)
//	})
//}

// 旋转数组的最小数字
// 已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,2,4,5,6,7] 在变化后可能得到：
// 若旋转 4 次，则可以得到 [4,5,6,7,0,1,2]
// 若旋转 7 次，则可以得到 [0,1,2,4,5,6,7]
// 注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。
// 给你一个元素值 互不相同 的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array
func findMin(nums []int) int {
	// 观察数据可知，最小元素的左边的元素 都大于 最小元素的右边的元素
	// 选定两个指针lo,hi,并获得mid元素，如果mid元素 < hi元素，说明最小元素在lo与mid中间，如果mid > hi,说明最小元素在mid和hi之间
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] < nums[hi] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return nums[hi]
}

// 查找区间
// 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
// 如果数组中不存在目标值 target，返回[-1, -1]。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array
func searchRange(nums []int, target int) []int {
	l, r := -1, -1
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] < target {
			lo = mid + 1
		} else if nums[mid] > target {
			hi = mid - 1
		} else {
			l = findLeft(nums, target, lo, mid)
			r = findRight(nums, target, mid, hi)
			break
		}

	}
	return []int{l, r}
}

func findLeft(nums []int, target int, lo, hi int) int {
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			hi = mid
			continue
		}
		if nums[mid] < target {
			lo = mid + 1
		}
	}
	return hi
}

func findRight(nums []int, target int, lo, hi int) int {
	result := lo
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			result = mid
			lo = mid + 1
			continue
		}
		if nums[mid] > target {
			hi = mid - 1
		}
	}
	return result
}
