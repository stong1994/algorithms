package fourth_edition

// 插入排序，思想为“洗扑克牌”，排序过程中保证前边是排好序的
type Insertion struct{}

func (s Insertion) Sort(list []Comparable) {
	for i := 1; i < len(list); i++ {
		for j := i; j > 0 && list[j].CompareTo(list[j-1]) < 0; j-- {
			list[j-1], list[j] = list[j], list[j-1]
		}
	}
}
