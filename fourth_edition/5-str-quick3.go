package fourth_edition

// 三向字符串快速排序
// 高位优先的字符串可能会创建大量（空）子数组，而三向字符串快速排序的切分总是只有三个。
// 因此三向字符串快速排序能够很好处理等值键、有较长公共前缀的键、取值范围较小的键和小数组。
// *此算法特别适合含有较长公共前缀的字符串
type Quick3 struct{}

func (q Quick3) Sort(a []string) {
	q.sort(a, 0, len(a)-1, 0)
}

func (q Quick3) sort(a []string, lo, hi, d int) {
	if hi <= lo {
		return
	}
	lt, gt := lo, hi
	v := charAt(a[lo], d)
	for i := lo + 1; i <= gt; {
		t := charAt(a[i], d)
		if t < v {
			a[lt], a[i] = a[i], a[lt]
			lt++
			i++
		} else if t > v {
			a[i], a[gt] = a[gt], a[i]
			gt--
		} else {
			i++
		}
	}
	q.sort(a, lo, lt-1, d)
	if v >= 0 {
		q.sort(a, lt, gt, d+1)
	}
	q.sort(a, gt+1, hi, d)
}
