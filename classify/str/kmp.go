package str

// 演化
func match_old1(txt, pat string) int {
	pmt := getPMT(pat)
	i, j := 0, 0 // i、j分别为txt和pat当前匹配的索引
	for i < len(txt) && j < len(pat) {
		if txt[i] == pat[j] {
			i++
			j++
			continue
		}
		if j == 0 { // j为0，说明子串中没有能够匹配的前后缀，直接将残尺移动一位
			i++
			continue
		}
		j = pmt[j-1]
	}
	if j == len(pat) {
		return i - j
	}
	return -1
}

func getPMT(pat string) []int {
	pmt := make([]int, len(pat))
	pmt[0] = 0 // 一个字符没有前缀和后缀（由如何使用决定第一个元素的值为0）
	i, j := 1, 0
	for i < len(pat) {
		if pat[i] == pat[j] {
			j++ // 长度等于索引+1
			pmt[i] = j
			i++
			continue
		}
		if j == 0 {
			pmt[i] = 0
			i++
			continue
		}
		j = pmt[j-1]
	}
	return pmt
}

func getNext_old(pat string) []int {
	next := make([]int, len(pat))
	next[0] = 0
	i, j := 1, 0
	for i < len(pat)-1 { // 可知next数组和pat的最后一个字符无关
		if pat[i] == pat[j] {
			i++
			j++
			next[i] = j
			continue
		}
		if j == 0 {
			i++
			next[i] = 0
			continue
		}
		j = next[j]
	}
	return next
}

func match_old2(txt, pat string) int {
	next := getNext_old(pat)
	i, j := 0, 0 // i、j分别为txt和pat当前匹配的索引
	for i < len(txt) && j < len(pat) {
		if txt[i] == pat[j] {
			i++
			j++
			continue
		}
		if j == 0 { // j为0，说明子串中没有能够匹配的前后缀，直接将残尺移动一位
			i++
			continue
		}
		j = next[j]
	}
	if j == len(pat) {
		return i - j
	}
	return -1
}

func getNext(pat string) []int {
	next := make([]int, len(pat))
	next[0] = -1
	i, j := 1, -1
	for i < len(pat)-1 { // 可知next数组和pat的最后一个字符无关
		if j == -1 || pat[i] == pat[j] {
			i++
			j++
			next[i] = j
			continue
		}
		j = next[j]
	}
	return next
}

func match(txt, pat string) int {
	next := getNext(pat)
	i, j := 0, 0 // i、j分别为txt和pat当前匹配的索引
	for i < len(txt) && j < len(pat) {
		if j == -1 || txt[i] == pat[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == len(pat) {
		return i - j
	}
	return -1
}
