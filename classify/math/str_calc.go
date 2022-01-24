package math

import "strings"

// 字符串计算

// 二进制求和
// 给你两个二进制字符串，返回它们的和（用二进制表示）。
// 输入为 非空 字符串且只包含数字1和0。
// 提示：
//	每个字符串仅由字符 '0' 或 '1' 组成。
//	1 <= a.length, b.length <= 10^4
//	字符串如果不是 "0" ，就都不含前导零。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/add-binary
func addBinary(a string, b string) string {
	// 简化问题：如果两个字符串不相等，则为较短的字符串前置补零
	l1, l2 := len(a), len(b)
	if l1 > l2 {
		b = strings.Repeat("0", l1-l2) + b
	} else {
		a = strings.Repeat("0", l2-l1) + a
	}
	flag := 0
	result := make([]byte, len(a))
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == '0' && b[i] == '0' {
			result[i] = uint8('0' + flag)
			flag = 0
			continue
		}
		if a[i] == '1' && b[i] == '1' {
			result[i] = uint8('0' + flag)
			flag = 1
			continue
		}
		result[i] = uint8('1' - flag)
		flag = flag & 1 // 旧flag为1，新flag才为1
	}
	if flag == 1 {
		return "1" + string(result)
	}
	return string(result)
}

// 字符串相加
// 给定两个字符串形式的非负整数num1 和num2，计算它们的和并同样以字符串形式返回。
// 你不能使用任何內建的用于处理大整数的库（比如 BigInteger），也不能直接将输入的字符串转换为整数形式。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/add-strings
func addStrings(num1 string, num2 string) string {
	// 简化问题：对较短的字符串前置补0
	l1, l2 := len(num1), len(num2)
	if l1 > l2 {
		num2 = strings.Repeat("0", l1-l2) + num2
	} else {
		num1 = strings.Repeat("0", l2-l1) + num1
	}
	result := make([]byte, len(num1))
	preAdd := uint8(0)
	for i := len(num1) - 1; i >= 0; i-- {
		a := num1[i] - '0'
		b := num2[i] - '0'
		val := a + b + preAdd
		if val >= 10 {
			preAdd = 1
			result[i] = val - 10 + '0'
		} else {
			preAdd = 0
			result[i] = val + '0'
		}
	}
	if preAdd == 0 {
		return string(result)
	}
	return "1" + string(result)
}
