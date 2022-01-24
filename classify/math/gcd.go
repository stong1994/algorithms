package math

import "math"

// 最大公约数

// 找到数组中的最大公约数
// 给你一个整数数组 nums ，返回数组中最大数和最小数的 最大公约数 。
// 两个数的最大公约数 是能够被两个数整除的最大正整数。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/find-greatest-common-divisor-of-array
func findGCD(nums []int) int {
	var (
		max = math.MinInt32
		min = math.MaxInt32
	)
	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	return gcd2(min, max)
}

// 最大公约数
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// 使用位操作和减法求解最大公约数
// 对于 a 和 b 的最大公约数 f(a, b)，有：
//	如果 a 和 b 均为偶数，f(a, b) = 2*f(a/2, b/2);
//	如果 a 是偶数 b 是奇数，f(a, b) = f(a/2, b);
//	如果 b 是偶数 a 是奇数，f(a, b) = f(a, b/2);
//	如果 a 和 b 均为奇数，f(a, b) = f(b, a-b);
// 乘 2 和除 2 都可以转换为移位操作。
func gcd2(a, b int) int {
	if a < b {
		return gcd2(b, a)
	}
	if b == 0 {
		return a
	}
	isAEven, isBEven := a&1 == 0, b&1 == 0
	if isAEven && isBEven {
		return 2 * gcd2(a>>1, b>>1)
	} else if isAEven && !isBEven {
		return gcd2(a>>1, b)
	} else if !isAEven && isBEven {
		return gcd2(a, b>>1)
	} else {
		return gcd2(b, a-b)
	}
}

//最小公倍数为两数的乘积除以最大公约数。
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
