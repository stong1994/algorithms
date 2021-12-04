package fourth_edition

type Selection struct{}

func (s Selection) Sort(list []Comparable) {
	for i := 0; i < len(list); i++ {
		minVal := list[i]
		minIdx := i
		for j := i; j < len(list); j++ {
			if list[j].CompareTo(minVal) < 0 {
				minVal = list[j]
				minIdx = j
			}
		}
		list[i], list[minIdx] = list[minIdx], list[i]
	}
}
