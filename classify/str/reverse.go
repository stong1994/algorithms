package str

import "strings"

// 字符串移位包含的问题
// 给定两个字符串s1和s2，要求判定s2是否能够被s1做循环移位(rotate)得到的字符串包含。
// 例如，给定s1 = AABCD 和s2 = CDAA,返回true；给定s1 = ABCD和s2=ACBD，返回false
// 来源：《编程之美》3.1
func strRotate(s1, s2 string) bool {
	// 方法1：暴力破解，对s1分别从第1个元素到最后一个元素进行移位
	// AABCD可得 ABCDA  BCDAA CDAAB DAABC 分别与s2进行前缀匹配，如果存在，返回true
	//for i := 0; i < len(s1); i++ {
	//	if strings.HasPrefix(s1[i:]+ s1[:i], s2) {
	//		return true
	//	}
	//}
	//return false

	// 方法2：循环移位后的s1一定是s1s1的字串，而s2又是循环移位后的s1的字串，因此判断s2是否是s1s1的子串即可
	return strings.Contains(s1+s1, s2)
	// 能否更进一步，不需要申请过多新的空间？
}

// 数组循环移位
// 设计一个算法，把一个含有N个元素的数组循环右移K位，要求时间复杂度为O(n),且只允许使用两个附加变量。
// 来源：《编程之美》2.17
func reverseShift(arr []int, N, K int) {
	// 将123abcd右移4位的abcd123，为了满足时间和空间要求，不能用暴力破解。
	// 我们可以把123和abcd分别反转，得到321dcba，再整体进行反转得abcd123，符合题目要求。
	reverse := func(arr []int, left, right int) {
		for ; left < right; left, right = left+1, right-1 {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	K %= N
	reverse(arr, 0, N-K-1)
	reverse(arr, N-K, N-1)
	reverse(arr, 0, N-1)
}

// 字符串中单词的翻转，如字符串为“I am a student”，翻转为“student a am I”
// 来源：程序员代码面试指南
func reverseWords(s string) string {
	// 先翻转每个单词，最后翻转整个字符串
	list := []byte(s)
	reverse := func(data []byte, left, right int) {
		for ; left < right; left, right = left+1, right-1 {
			data[left], data[right] = data[right], data[left]
		}
	}
	l, r := -1, -1 // 当l和r都不为-1时，list[l:r+1]为一个单词
	for i, v := range list {
		if v != ' ' {
			if l == -1 {
				l = i
			} else {
				r = i
			}
		} else { // 当前元素为空格，如果是第一个空格，那么翻转前一个单词，并初始化l和r
			if l != -1 && r != -1 {
				reverse(list, l, r)
			}
			l = -1
			r = -1
		}
	}
	if l != -1 && r != -1 {
		reverse(list, l, r)
	}
	reverse(list, 0, len(list)-1)
	return string(list)
}
