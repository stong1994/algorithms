package fourth_edition

import "math"

// 归并排序——将两个更大的有序数组归并为一个更大的有序数组——分为自顶向下归并、自底向上归并
// 当数组长度为2的幂时，两者比较次数和数组访问次数相同，只是顺序不同。
type Merge struct {
	sorter ISortable
}

func (s Merge) Sort(list []Comparable) {
	s.bottom2topSort(list)
}

// 自顶向下的归并排序
// 将数组由整体切割成多层的子数组，在最底层的子数组通过另一个排序算法进行排序，然后对这些子数组进行归并
func (s Merge) top2bottomSort(list []Comparable, left, right int) {
	if left >= right {
		return
	}
	// 子数组排序，长度自定义
	if right-left < 3 {
		s.sorter.Sort(list[left : right+1])
		return
	}
	mid := (left + right) / 2
	s.top2bottomSort(list, left, mid)
	s.top2bottomSort(list, mid+1, right)
	s.merge(list, left, mid, right)
}

// 自底向上的归并排序
// 先归并那些微型数组，然后再成对归并得到的子数组
// 代码量要比自顶向下的排序少很多
func (s Merge) bottom2topSort(list []Comparable) {
	l := len(list)
	for sz := 1; sz < l; sz += sz {
		for left := 0; left < l-sz; left += sz + sz {
			s.merge(list, left, left+sz-1, int(math.Min(float64(left+sz+sz-1), float64(l-1))))
		}
	}
}

// 将list[left,mid]和list[mid+1,right]归并
// 通过逆序双指针来对list[left:right+1]排序
func (s Merge) merge(list []Comparable, left, mid, right int) {
	rightCopy := make([]Comparable, right-mid)
	copy(rightCopy, list[mid+1:right+1])
	for i, j, tail := mid, len(rightCopy)-1, right; tail >= left; tail-- {
		var cur Comparable
		if i == left-1 {
			cur = rightCopy[j]
			j--
		} else if j == -1 {
			cur = list[i]
			i--
		} else if list[i].CompareTo(rightCopy[j]) <= 0 {
			cur = rightCopy[j]
			j--
		} else {
			cur = list[i]
			i--
		}
		list[tail] = cur
	}

}
