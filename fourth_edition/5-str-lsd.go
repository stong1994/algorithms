package fourth_edition

// 低位优先的字符串排序算法能够稳定地将定长字符串排序
// 基于键索引计数法实现。从右往左依次对元素进行排序，因此结果会从左往右按顺序排列——这依赖键索引技术法的排序是“稳定”的
type LSD struct{}

// 通过前w个字符将a排序
func (l LSD) sort(a []string, w int) {
	length := len(a)
	R := 256
	aux := make([]string, length+1)
	for d := w - 1; d >= 0; d-- {
		count := make([]int, R+1)
		for i := 0; i < length; i++ { // 统计频率
			count[a[i][d]+1]++
		}
		for r := 0; r < R; r++ { // 将频率转换为索引
			count[r+1] += count[r]
		}
		for i := 0; i < length; i++ { // 将元素分类
			aux[count[a[i][d]]] = a[i]
			count[a[i][d]]++
		}
		for i := 0; i < length; i++ { // 回写
			a[i] = aux[i]
		}
	}
}
