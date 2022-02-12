package str

// 马拉车算法: 最长的回文字符串
// 相关资料：
//   博客：https://www.acwing.com/file_system/file/content/whole/index/content/446985/#fn:1
//   动图：http://manacher-viz.s3-website-us-east-1.amazonaws.com/#/
// 预处理字符串，将其变为奇数长度
func preProcess(s string) string {
	if len(s) == 0 {
		return "^$"
	}
	result := []byte("^")
	for _, v := range s {
		result = append(result, '#', byte(v))
	}
	result = append(result, '#', '$')
	return string(result)
}

func manacher(s string) string {
	newStr := preProcess(s)
	N := len(newStr)
	p := make([]int, N) // newStr对应的数组，元素值为在newStr中以当前元素为中心的回文字符串的最大半径。那么p[i]-1就是原字符串s以i为中心的最大回文长度
	var (
		rt  = 0 // 已经计算过的回文字符串能达到的最远右边界的下一个元素索引,即rt=max(j+p[j]), 1<=j<=i-1
		mid = 0 // mid表示rt所对应的最左侧的回文中心，有mid + p[mid] = rt
	)
	for i := 1; i < N-1; i++ { // i为0时，newStr对应的元素为“哨兵”^，其肯定不是回文中心，因此可以从第一个元素开始。同理，最后一个元素也可以忽略
		if rt > i { // i在已知的回文字符串内，可利用已知条件求出p[i]的已知值，稍后使用中心扩展寻找其最大值
			iMirror := 2*mid - i         // 以mid为中心的回文字符串中的与第i个元素对称的元素索引
			p[i] = min(rt-i, p[iMirror]) // 防止p[i]超出已知的最大回文右边界,若p[iMrror]超出已知的右边界，那么p[i]等于i到右边界的距离
		} else {
			p[i] = 1 // i大于等于已知的最大回文字符串的右边界，因此将当前值设为1，并随后进行“中心扩展”
		}
		// 中心扩展（有两端的哨兵存在，不用担心越界问题）
		for newStr[i+p[i]] == newStr[i-p[i]] {
			p[i]++
		}
		if i+p[i] > rt { // 更新当前最大的rt
			mid = i
			rt = i + p[i]
		}
	}
	// 找出p的最大值
	var (
		maxLen    = 0
		centerIdx = 0
	)
	for i := 0; i < N-1; i++ {
		if p[i] > maxLen {
			maxLen = p[i] - 1
			centerIdx = i
		}
	}
	startIdx := (centerIdx - maxLen) / 2
	return s[startIdx : startIdx+maxLen]

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
