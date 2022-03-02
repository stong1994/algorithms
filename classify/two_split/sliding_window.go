package two_split

// 无重复字符的最长子串
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
// 提示：
// 0 <= s.length <= 5 * 104
// s 由英文字母、数字、符号和空格组成
// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	// 指定两个指针l,r。创建一个哈希表store来存储l和r之间的值，每次向右移动r，直到找到一个重复值，此时的r-l+1为当前无重复字符串的长度
	// 然后向右移动l，开始新的一轮比较
	n := len(s)
	l, r := 0, 0
	maxLen := 0
	store := make(map[byte]struct{})
	for r < n {
		if _, exist := store[s[r]]; !exist {
			store[s[r]] = struct{}{}
			r++
			continue
		}
		maxLen = max(maxLen, r-l)
		delete(store, s[l])
		l++
	}
	maxLen = max(maxLen, r-l)
	return maxLen
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func lengthOfLongestSubstring2(s string) int {
	res := 0
	l := 0
	mapChar := [128]int{}
	for i := 0; i < len(s); i++ {
		for mapChar[s[i]] > 0 {
			mapChar[s[l]]--
			l++
		}
		mapChar[s[i]]++
		res = max(res, i-l+1)
	}
	return res
}

// 重复的DNA序列
// DNA序列由一系列核苷酸组成，缩写为'A','C','G'和'T'.。
// 例如，"ACGAATTCCG"是一个 DNA序列 。
// 在研究 DNA 时，识别 DNA 中的重复序列非常有用。
// 给定一个表示 DNA序列 的字符串 s ，返回所有在 DNA 分子中出现不止一次的长度为10的序列(子字符串)。你可以按 任意顺序 返回答案。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/repeated-dna-sequences
func findRepeatedDnaSequences(s string) []string {
	// 将每10个字符串放入哈希表中，如果个数超过1，则将其放入结果列表中
	var result []string
	store := make(map[string]int)
	for i := 0; i+9 < len(s); i++ {
		sub := s[i : i+10]
		if n, exist := store[sub]; exist {
			if n == 1 {
				result = append(result, sub)
			}
		}
		store[sub]++
	}
	return result
}

func findRepeatedDnaSequences2(s string) []string {
	// 哈希表+前缀和
	var result []string
	N, P := 1e5+10, 131313
	h, p := make([]int, int(N)), make([]int, int(N)) // 哈希数组 次方数组
	n := len(s)
	p[0] = 1
	for i := 1; i <= n; i++ {
		h[i] = h[i-1]*P + int(s[i-1])
		p[i] = p[i-1] * P
	}
	store := make(map[int]int)
	for i := 1; i+9 <= len(s); i++ {
		j := i + 9
		hash := h[j] - h[i-1]*p[j-i+1]
		cnt := store[hash]
		if cnt == 1 {
			result = append(result, s[i-1:j])
		}
		store[hash] = cnt + 1
	}
	return result
}

// 字符串的排列
// 给你两个字符串s1和s2 ，写一个函数来判断 s2 是否包含 s1的排列。如果是，返回 true ；否则，返回 false 。
// 换句话说，s1 的排列之一是 s2 的 子串 。
// 提示：
//	1 <= s1.length, s2.length <= 10^4
//	s1 和 s2 仅包含小写字母
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/permutation-in-string
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	// s1和s2仅包含小写字母，即可用长度为26的数组表示，每个元素的值为对应字母在s1中的个数
	arr1 := [26]int{}
	for _, v := range s1 {
		arr1[v-'a']++
	}
	// 用一个长度为len(s1)的双指针，不断向右移动，比较字母个数是否和arr1中一致
	//store := make(map[int]int)
	arr2 := [26]int{}
	l, r := 0, len(s1)-1
	for i := 0; i < len(s1); i++ {
		arr2[s2[i]-'a']++
	}

	if arr1 == arr2 {
		return true
	}
	for r < len(s2)-1 {
		arr2[s2[l]-'a']--
		l++
		r++
		arr2[s2[r]-'a']++
		if arr1 == arr2 {
			return true
		}
	}
	return false
}