package fourth_edition

// 子字符串查找算法，在txt中找到pat第一次出现的未知

// 暴力查找：设置两个指针i，j，i为pat首字母在txt的位置，j为pat时的位置。
// i从0开始遍历txt，一旦找到字母为pat首字母的元素，就将j重置为0进行匹配，如果能够完全匹配pat，则返回i，否则i++
func violentSearchSubStr(txt, pat string) int {
	M := len(txt)
	N := len(pat)
	for i := 0; i <= M-N; i++ { // 对于特殊情况txt和pat为空，此时i为0，M-N也为0，因此会进入循环，且返回为0
		var j int
		for j = 0; j < N; j++ {
			if txt[i+j] != pat[j] {
				break
			}
		}
		if j == N {
			return i
		}
	}
	return -1
}

// 另一种形式的暴力破解：显式回退
func violentSearchSubStr2(txt, pat string) int {
	var (
		M    = len(txt)
		N    = len(pat)
		i, j = 0, 0
	)

	for ; i < M && j < N; i++ {
		if txt[i] == pat[j] {
			j++
			continue
		}
		i -= j
		j = 0
	}
	if j == N {
		return i - N
	}
	return -1
}

// KMP算法
// 当i,j匹配失败时，无需将i回退i-j个。在匹配过程中记录所匹配的字符索引，然后回退时则可以从该索引处进行匹配
// TODO 有限状态机？
func searchSubStrKMP(txt, pat string) int {
	var (
		M    = len(txt)
		N    = len(pat)
		i, j = 0, 0
	)

	dfa := searchSubStrKMPMakeDFA(pat)
	for ; i < M && j < N; i++ {
		j = dfa[txt[i]][j]
	}
	if j == N {
		return i - N
	}
	return -1
}

// 构建有限状态机 dfa[i][j]
// 暴力破解中，如果匹配失败，需要重新扫描的字符是pat[1]到pat[j-1]（首字母被忽略是因为要向右移动1位，忽略最后一位是因为匹配失败）
// 这些字符都是已知的，所以对于每个匹配失败的位置都能预先找到“重启”DFA的正确状态。
// 重启位置设为X，X<j，所以由已构造的DFA中，X的下一个值为dfa[pat[j]][X]
func searchSubStrKMPMakeDFA(pat string) [][]int {
	M := len(pat)
	R := 256
	dfa := make([][]int, R) // 竖列为256个基础元素、横列为目标字符串的元素
	for i := range dfa {
		dfa[i] = make([]int, M)
	}
	dfa[pat[0]][0] = 1
	//		匹配失败时：将dfa[][X]复制到dfa[][j]
	// 		匹配成功时：将dfa[j][j]设置为j+1
	//      更新X
	var X = 0
	for j := 1; j < M; j++ {
		for c := 0; c < R; c++ {
			dfa[c][j] = dfa[c][X] // 对于每个字符元素的子数组（dfa[c]），
			// 其对应的模式串字母对应的值为”重启位置“，即上一个该字符对应的位置
		}
		dfa[pat[j]][j] = j + 1
		X = dfa[pat[j]][X]
	}
	return dfa
}

// Boyer-Moore字符串查找算法
// 当回退时，如果可以从右向左扫描模式串并将它和文本匹配，就能得到很快的查找算法。
func searchSubStrBM(txt, pat string) int {
	M := len(pat)
	N := len(txt)
	R := 256
	// 初始化right数组，构建一个记录了pat中每个元素在pat中最右的位置
	right := make([]int, R)
	for c := 0; c < R; c++ {
		right[c] = -1
	}
	for j := 0; j < M; j++ {
		right[pat[j]] = j
	}

	var skip int
	for i := 0; i <= N-M; i += skip {
		skip = 0
		// 从右往左扫描，一旦匹配失败，则跳跃到文本中的字符和它在模式串中出现的最右位置对齐
		for j := M - 1; j >= 0; j-- {
			if pat[j] != txt[i+j] {
				skip = j - right[txt[i+j]]
				if skip < 1 {
					skip = 1
				}
				break
			}
		}
		if skip == 0 {
			return i
		}
	}
	return -1
}

// Rabin-Karp指纹字符串查找算法
// 基于hash的思想来比较子字符串，传统上的算法很慢，通过一些数学性质能够加快hash计算的速度，以此来加快查找
type RabinKarp struct {
	R       int // 字母表的大小 256
	Q       int // 一个很大的素数 暂定为3571
	pat     string
	M       int // 模式字符串的长度
	RM      int // R^(M-1)%Q
	patHash int
}

func NewRabinKarp(pat string) *RabinKarp {
	rk := &RabinKarp{
		R:   256,
		Q:   3571,
		pat: pat,
		M:   len(pat),
	}
	rk.initPatHash()
	return rk
}

func (rk RabinKarp) search(txt string) int {
	N := len(txt)
	if N < rk.M {
		return -1
	}
	txtHash := rk.hash(txt, rk.M)
	if txtHash == rk.patHash && rk.check(0) {
		return 0
	}
	for i := rk.M; i < N; i++ {
		txtHash = (txtHash + rk.Q - rk.RM*int(txt[i-rk.M])%rk.Q) % rk.Q
		txtHash = (txtHash*rk.R + int(txt[i])) % rk.Q
		if rk.patHash == txtHash {
			if rk.check(i - rk.M + 1) {
				return i - rk.M + 1
			}
		}
	}
	return -1
}

func (rk *RabinKarp) initPatHash() {
	RM := 1
	for i := 1; i <= rk.M-1; i++ {
		RM = (rk.R * RM) % rk.Q
	}
	rk.patHash = rk.hash(rk.pat, rk.M)
	rk.RM = RM
}

func (rk RabinKarp) hash(key string, M int) int {
	h := 0
	for j := 0; j < M; j++ {
		h = (rk.R*h + int(key[j])) % rk.Q
	}
	return h
}

func (rk RabinKarp) check(i int) bool {
	return true
}

func next(str string) []int {
	result := make([]int, len(str))
	result[0] = -1
	i, j := 0, -1
	for i < len(str)-1 {
		if j == -1 || str[i] == str[j] {
			i++
			j++
			result[i] = j
		} else {
			j = result[j] // 为什么是result[j]而不是-1，对于str中的第i个元素，
			// 其前缀长度最大是result[i-1]+1,其次是result[result[i-1]]+1（已知str[i-1]与str[0]相等，
			// 设最大长度为n，则str[i-1-n]与str[n]相等，此时result[i-1]=n）,以此类推
			//
		}
	}
	return result
}

func kmp(target, pat string) int {
	nexts := next(pat)
	i, j := 0, 0
	for i < len(target) && j < len(pat) {
		if j == -1 || target[i] == pat[j] {
			i++
			j++
			continue
		}
		j = nexts[j]
	}
	if j == len(pat) {
		return i - j
	}
	return -1
}
