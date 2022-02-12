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
	rebuild := preProcess(s)
	var (
		rMax = 0                         // 当前右边界
		p    = make([]int, len(rebuild)) // 回环长度数组
		mid  = 0                         // 当前右边界为rMax的回环子串的中心索引
	)
	for i := 1; i < len(rebuild)-1; i++ { // 首尾两个哨兵不用管
		if i > rMax { // 没有经验可供借鉴，需要从零开始进行中心扩展
			p[i] = 0
		} else {
			p[i] = min(rMax-i, p[mid*2-i]) // 借鉴下历史经验——镜像值,如果镜像值超过rMax-i，说明借鉴的不是当前的历史经验
		}
		// 中心扩展
		for rebuild[i+p[i]+1] == rebuild[i-p[i]-1] { // 因为哨兵一定与其他元素不等，因此无需考虑边界溢出
			p[i]++
		}
		// 动态更新右边界和中心索引
		if i+p[i] > rMax {
			mid = i
			rMax = mid + p[mid] // 右边界索引 = 当前中心索引+半径长度
		}
	}
	var (
		maxLen    = 0
		centerIdx = 0
	)
	for i := 1; i < len(p)-1; i++ {
		if p[i] > maxLen {
			maxLen = p[i]
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
