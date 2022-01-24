package math

import (
	"strconv"
)

// 进制数

// 七进制数
// 给定一个整数 num，将其转化为 7 进制，并以字符串形式输出。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/base-7
func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	// 辗转相除
	var (
		result string
		neg    = num < 0
	)
	if neg {
		num = -num
	}
	for num > 0 {
		result = strconv.Itoa(num%7) + result
		num /= 7
	}
	if neg {
		result = "-" + result
	}
	return result
}

//  数字转换为十六进制数
// 给定一个整数，编写一个算法将这个数转换为十六进制数。对于负整数，我们通常使用补码运算方法。
// 注意:
// 十六进制中所有字母(a-f)都必须是小写。
// 十六进制字符串中不能包含多余的前导零。如果要转化的数为0，那么以单个字符'0'来表示；对于其他情况，十六进制字符串中的第一个字符将不会是0字符。
// 给定的数确保在32位有符号整数范围内。
// 不能使用任何由库提供的将数字直接转换或格式化为十六进制的方法。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/convert-a-number-to-hexadecimal
func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	if num < 0 { // 处理负数
		num = 1<<32 + num
	}
	hexMap := map[int]string{
		0:  "0",
		1:  "1",
		2:  "2",
		3:  "3",
		4:  "4",
		5:  "5",
		6:  "6",
		7:  "7",
		8:  "8",
		9:  "9",
		10: "a",
		11: "b",
		12: "c",
		13: "d",
		14: "e",
		15: "f",
	}
	var result string
	for num > 0 {
		result = hexMap[num%16] + result
		num /= 16
	}
	return result
}

// Excel表列名称: 26进制
// 给你一个整数columnNumber ，返回它在 Excel 表中相对应的列名称。
//例如：
//A -> 1
//B -> 2
//C -> 3
//...
//Z -> 26
//AA -> 27
//AB -> 28
//...
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/excel-sheet-column-title
func convertToTitle(columnNumber int) string {
	// 可以看做是26进制的计算
	numMap := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
		"U", "V", "W", "X", "Y", "Z"}
	var (
		result string
	)
	for columnNumber > 0 {
		// A对应1，Z对应26，如果余数为0，说明此时应对应Z
		remain := columnNumber % 26
		if remain == 0 {
			result = "Z" + result
			columnNumber = columnNumber/26 - 1
		} else {
			result = numMap[remain-1] + result
			columnNumber /= 26
		}
	}
	return result
}
