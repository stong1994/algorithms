package math

import "math"

// 给定一个 正整数 num ，编写一个函数，如果 num 是一个完全平方数，则返回 true ，否则返回 false 。
// 进阶：不要 使用任何内置的库函数，如sqrt 。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/valid-perfect-square
func isPerfectSquare(num int) bool {
	// 平方序列 1,4,9,16,25。。。间隔为等差数列3,5,7,9
	//grade := 1
	//curNum := 0
	//for {
	//	curNum = grade +curNum
	//	if curNum == num {
	//		return true
	//	}
	//	if curNum > num {
	//		return false
	//	}
	//	grade += 2
	//}
	subNum := 1 // 1, 3, 5, 7, 9 相加即为一个平方数
	for num > 0 {
		num -= subNum
		subNum += 2
	}
	return num == 0
}

// 3的幂
// 给定一个整数，写一个函数来判断它是否是 3的幂次方。如果是，返回 true ；否则，返回 false 。
// 整数 n 是 3 的幂次方需满足：存在整数 x 使得 n == 3x
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/power-of-three
func isPowerOfThree(n int) bool {
	// 3的幂次的质因子只有3，而所给出的n如果也是3的幂次，故而题目中所给整数范围内最大的3的幂次的因子只能是3的幂次，1162261467是3的19次幂，是整数范围内最大的3的幂次
	return n > 0 && 1162261467%n == 0
}

// 除自身以外数组的乘积
// 给你一个长度为n的整数数组nums，其中n > 1，返回输出数组output，其中 output[i]等于nums中除nums[i]之外其余各元素的乘积。
//	提示：题目数据保证数组之中任意元素的全部前缀元素和后缀（甚至是整个数组）的乘积都在 32 位整数范围内。
//	说明: 请不要使用除法，且在O(n) 时间复杂度内完成此题。
//	进阶：
//		你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/product-of-array-except-self
func productExceptSelf(nums []int) []int {
	// 对于第i个结果，其值为左边的元素相乘再乘以右边的元素。
	// 因此遍历两次nums，第一次从左到右，并暂存每个结果，第二次从右到左，乘以
	result := make([]int, len(nums))

	left := nums[0]
	for i := 1; i < len(nums); i++ {
		result[i] = left
		left *= nums[i]
	}
	result[0] = 1

	right := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		result[i] = right * result[i]
		right *= nums[i]
	}
	return result
}

// 三个数的最大乘积
// 给你一个整型数组 nums ，在数组中找出由三个数组成的最大乘积，并输出这个乘积。
// 提示：
// 3 <= nums.length <=10^4
// -1000 <= nums[i] <= 1000
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/maximum-product-of-three-numbers
func maximumProduct(nums []int) int {
	// 如果都是正整数，那么直接找到3个最大值即可
	// 如果存在负数，则找到最小的两个负数和最大的正数
	// 于是，我们需要找到5个值，前三个最大值+前两个最小值：max1、max2、max3、min1、min2
	// 结果为max(max1*max2*max3, max1*min1*min2)
	var (
		max1, max2, max3 = math.MinInt32, math.MinInt32, math.MinInt32
		min1, min2       = math.MaxInt32, math.MaxInt32
	)
	for _, v := range nums {
		if v > max1 {
			max3 = max2
			max2 = max1
			max1 = v
		} else if v > max2 {
			max3 = max2
			max2 = v
		} else if v > max3 {
			max3 = v
		}
		if v < min1 {
			min2 = min1
			min1 = v
		} else if v < min2 {
			min2 = v
		}
	}
	v1, v2 := max1*max2*max3, max1*min1*min2
	if v1 > v2 {
		return v1
	}
	return v2
}
