package fourth_edition

// 选择排序。选出最小的与“当前”数组的第一个交换
type Selection struct{}

func (s Selection) Sort(list []Comparable) {
	for i := 0; i < len(list); i++ {
		minVal := list[i]
		minIdx := i
		for j := i + 1; j < len(list); j++ {
			if list[j].CompareTo(minVal) < 0 {
				minVal = list[j]
				minIdx = j
			}
		}
		list[i], list[minIdx] = list[minIdx], list[i]
	}
}
