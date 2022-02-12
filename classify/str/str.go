package str

// 两个字符串包含的字符是否完全相同
// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
// 注意：若s 和 t中每个字符出现的次数都相同，则称s 和 t互为字母异位词。
// s 和 t 仅包含小写字母
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/valid-anagram
func isAnagram(s string, t string) bool {
	// 将s和t中的元素存入两个hash中，并保存每个字符的个数，然后对比两个hash即可
	// 但是空间复杂度较高。题目要求s和t仅包含小写字母，因此用一个26个元素大小的数组即可。
	// 第一遍遍历s，将字母对应的数组元素的数量+1
	// 第二遍遍历t，将字母对应的数组元素的数量-1
	// 第三遍遍历数组，查看每个元素的值是否都为0
	var arr [26]int32
	for _, v := range s {
		arr[v-'a']++
	}
	for _, v := range t {
		arr[v-'a']--
		if arr[v-'a'] < 0 {
			return false
		}
	}
	for _, v := range arr {
		if v != 0 {
			return false
		}
	}
	return true
}

// 字符串同构
// 给定两个字符串s和t，判断它们是否是同构的。
//如果s中的字符可以按某种映射关系替换得到t，那么这两个字符串是同构的。
//每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。
// s 和 t 由任意有效的 ASCII 字符组成
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/isomorphic-strings
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// 维护从s到t的映射表与从t到s的映射表，遍历过程中不断比较
	// 虽然题目中提到了“数量要一致”，但是解法并不需要比较数量
	h1 := make(map[uint8]uint8)
	h2 := make(map[uint8]uint8)
	for i := 0; i < len(s); i++ {
		t1, exist1 := h1[s[i]]
		s1, exist2 := h2[t[i]]
		if exist1 != exist2 {
			return false
		}
		if exist1 {
			if t1 != t[i] || s1 != s[i] {
				return false
			}
		} else {
			h1[s[i]] = t[i]
			h2[t[i]] = s[i]
		}
	}
	return true
}
