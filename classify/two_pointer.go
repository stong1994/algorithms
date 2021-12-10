package classify

import (
	"math"
)

// 双指针系列

// 有序数组的两数之和
// https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/description/
// 思路：将两个指针分别指向数组首尾，通过其和的大小与target比较，然后再移动指针。时间复杂度为O(N)
func twoSum_167(numbers []int, target int) []int {
	for i, j := 0, len(numbers)-1; i < j; {
		sum := numbers[i] + numbers[j]
		if sum > target {
			j--
		} else if sum < target {
			i++
		} else {
			return []int{i + 1, j + 1}
		}
	}
	return nil
}

// 两数平方和
// https://leetcode-cn.com/problems/sum-of-square-numbers/description/
// 假设两个数字low和high，如果low^2+high^2小于c，则low++，如果大于c，则high--
// 这里有一篇分析为什么是low++，为什么是high--的文章：
//	https://leetcode-cn.com/problems/sum-of-square-numbers/solution/shuang-zhi-zhen-de-ben-zhi-er-wei-ju-zhe-ebn3/
func judgeSquareSum(c int) bool {
	var (
		low  = 0
		high = int(math.Sqrt(float64(c)))
	)
	for low <= high {
		sum := low*low + high*high
		if sum > c {
			high--
		} else if sum < c {
			low++
		} else {
			return true
		}
	}
	return false
}

// 反转字符串中的元音字母
// https://leetcode-cn.com/problems/reverse-vowels-of-a-string/description/
func reverseVowels(s string) string {
	yuan := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	result := []byte(s)
	for left, right := 0, len(result)-1; left < right; {
		if ok := yuan[result[left]]; !ok {
			left++
		} else if ok := yuan[result[right]]; !ok {
			right--
		} else {
			result[left], result[right] = result[right], result[left]
			left++
			right--
		}
	}
	return string(result)
}

// 验证回文字符串 Ⅱ
// https://leetcode-cn.com/problems/valid-palindrome-ii/description/
func validPalindrome(s string) bool {
	bts := []byte(s)
	for left, right := 0, len(bts)-1; left < right; {
		if bts[left] == bts[right] {
			left++
			right--
		} else {
			f1, f2 := true, true
			for i, j := left+1, right; i < j; i, j = i+1, j-1 {
				if bts[i] != bts[j] {
					f1 = false
					break
				}
			}
			for i, j := left, right-1; i < j; i, j = i+1, j-1 {
				if bts[i] != bts[j] {
					f2 = false
					break
				}
			}
			return f1 || f2
		}
	}
	return true
}

// 归并两个有序数组
// https://leetcode-cn.com/problems/merge-sorted-array/description/
// 构建一个临时数组，对两个数组分别设置两个指针，然后判断大小，将较小的元素放到临时数组中
// 如果不需要临时数组，则可以通过逆双指针，从后往前比较，然后赋值到nums1
func merge(nums1 []int, m int, nums2 []int, n int) {
	var result []int
	for i, j := 0, 0; ; {
		if i == m {
			result = append(result, nums2[j:]...)
			break
		}
		if j == n {
			result = append(result, nums1[i:len(nums1)-n]...)
			break
		}
		if nums1[i] <= nums2[j] {
			result = append(result, nums1[i])
			i++
		} else {
			result = append(result, nums2[j])
			j++
		}
	}
	copy(nums1, result)
}

// 逆向双指针
func mergeReverse(nums1 []int, m int, nums2 []int, n int) {
	for i, j, tail := m-1, n-1, m+n-1; tail >= 0; tail-- {
		var cur int
		if i == -1 {
			cur = nums2[j]
			j--
		} else if j == -1 {
			break
		} else if nums1[i] <= nums2[j] {
			cur = nums2[j]
			j--
		} else {
			cur = nums1[i]
			i--
		}
		nums1[tail] = cur
	}
}

// 判断链表中是否有环
// https://leetcode-cn.com/problems/linked-list-cycle/description/
// 通过hash可以求解。如果用双指针，则需使用快慢双指针——如果有环的话，移动快的指针一定会和移动慢的指针重合
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

// 通过删除字母匹配到字典里最长单词
// https://leetcode-cn.com/problems/longest-word-in-dictionary-through-deleting/description/
func findLongestWord(s string, dictionary []string) string {
	var match = func(str, match string) bool {
		// match: 12345678
		// str: 148
		// return: true
		if len(str) > len(match) {
			return false
		}
		for i, j := 0, 0; j < len(match); {
			if str[i] == match[j] {
				if i == len(str)-1 {
					return true
				}
				i++
				j++
			} else {
				j++
			}
		}
		return false
	}
	var result string
	for _, v := range dictionary {
		if match(v, s) {
			if result == "" || len(v) > len(result) || (len(v) == len(result) && v < result) {
				result = v
			}
		}
	}
	return result
}
