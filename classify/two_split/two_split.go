package two_split

// 找出数组中重复的数，数组值在 [1, n] 之间
// 给定一个包含n + 1 个整数的数组nums ，其数字都在[1, n]范围内（包括 1 和 n），可知至少存在一个重复的整数。
//假设 nums 只有 一个重复的整数 ，返回这个重复的数 。
//你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/find-the-duplicate-number
func findDuplicate(nums []int) int {
	// 要求只是用常数级的额外空间，不修改数组nums
	// 定义cnt[i]数组表示nums数组中小于等于i的数量，
	// 对于无重复数的nums，如[1,5,3,2,4],cnt为[1,2,3,4,5],
	// 对于有重复数的nums，如[1,5,2,3,2,4], cnt为[1,3,4,5,6]
	// 可以看到，对于重复数i存在：当j<i时，cnt[j]<= j; 当j>=i时，cnt[j]>i
	// 题目要求不允许保存cnt数组，因此可使用二分法找到该值
	// 通过二分法，找到mid，并找到数组中元素小于等于mid的数量cnt，如果cnt大于mid，则将mid赋值给right，否则将mid+1赋值给left。
	n := len(nums)
	left, right := 1, n-1 // left和right在这里不是指针，还是值
	for left < right {
		mid := left + (right-left)/2
		cnt := 0
		for _, v := range nums {
			if v <= mid {
				cnt++
			}
		}
		if cnt > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}
