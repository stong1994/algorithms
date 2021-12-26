package fourth_edition

// 对于长度不同的字符串进行排序，“低位优先”不再适用，需要升级为“高位优先”
// 对首字母通过键索引计数法进行排序，然后将各个“分组”内的子字符串进行“索引计数“排序，递归直到结束
// 缺点是对相同的字符串排序时，仍要”全量“比较
type MSD struct {
	R   int      // 基数
	M   int      // 小数组的切换阈值
	aux []string // 辅助数组
}

func charAt(s string, d int) int {
	if d >= len(s) {
		return -1
	}
	return int(s[d])
}

func NewMSD() *MSD {
	return &MSD{
		R: 256,
		M: 2, // 工业上用15左右的数字
	}
}

func (m *MSD) Sort(a []string) {
	length := len(a)
	m.aux = make([]string, length+1)
	m.sort(a, 0, length-1, 0)
}

// 以第d个字符为键将a[lo]至a[hi]排序
func (m *MSD) sort(a []string, lo, hi, d int) {
	if hi <= lo+m.M { // 对剩余部分用插入排序
		m.insertionSort(a, lo, hi, d)
		return
	}
	count := make([]int, m.R+2)
	for i := lo; i <= hi; i++ { // 计算频率
		count[charAt(a[i], d)+2]++ // TODO 理解+2
	}
	for r := 0; r < m.R; r++ { // 将频率转换为索引
		count[r+1] += count[r]
	}
	for i := lo; i <= hi; i++ { // 数据分类
		m.aux[count[charAt(a[i], d)+1]] = a[i]
		count[charAt(a[i], d)+1]++
	}
	for i := lo; i <= hi; i++ { // 回写
		a[i] = m.aux[i-lo]
	}
	for r := 0; r < m.R; r++ {
		m.sort(a, lo+count[r], lo+count[r+1]-1, d+1)
	}
}

// 从第d个字符开始对a[lo]至a[hi]排序
func (m MSD) insertionSort(a []string, lo, hi, d int) {
	for i := lo; i <= hi; i++ {
		for j := i; j > lo && m.less(a[j], a[j-1], d); j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}

func (m MSD) less(v, w string, d int) bool {
	vl := d
	if len(v) <= d {
		vl = len(v)
	}
	wl := d
	if len(w) <= d {
		wl = len(w)
	}
	return v[vl:] < w[wl:]
}
