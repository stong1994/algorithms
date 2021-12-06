package fourth_edition

// 希尔排序。对插入排序的优化：在插入排序中，将当前比较的元素
// 移动到最前面，需要交换途中的所有元素。希尔排序优化此问题，通过间隔h个元素进行比较-移动，
// 每遍历完一遍后，数组整体的混乱程度更低，然后不断降低间隔重复这个过程，直到间隔为1。
type Shell struct {
	N int
}

func (s Shell) Sort(list []Comparable) {
	N := len(list)
	h := 1
	for h*3+1 <= N {
		h = h*3 + 1 // 是否加1似乎影响不大？
	}

	for ; h >= 1; h = h / 3 {
		for i := h; i < N; i++ {
			// 把i插入到以h为间隔的前方数组
			for j := i; j >= h; j -= h {
				if list[j].CompareTo(list[j-1]) < 0 {
					list[j-1], list[j] = list[j], list[j-1]
				}
			}
		}
	}

}
