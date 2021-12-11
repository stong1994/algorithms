package fourth_edition

// 快速排序。
// 快排与归并排序的区别在于归并排序是直接将数组等比例拆分成小数组，然后进行归并，而快排是随机找一个数，然后将数组拆分成两个子数组。
// 关于快排为什么这么快，可以阅读这篇文章：http://mindhacks.cn/2008/06/13/why-is-quicksort-so-quick/
// 快排的速度依赖原始数据的顺序，为了保证其性能，我们往往会在排序前打乱原始数组
// 优化快排：
// 	1. 对于小数组，快速排序慢于插入排序，因此在小数组中可使用插入排序
//  2. 拆分数组时，取消随机找值，而是计算出中位数，并且拆分成三个数组，分别是[小于中位数]、[等于中位数]、[大于中位数]

type Quick struct{}

func (s Quick) Sort(list []Comparable) {
	s.quickSort(list, 0, len(list)-1)
}

func (s Quick) quickSort(list []Comparable, left, right int) {
	if left >= right {
		return
	}
	j := s.partition(list, left, right)
	s.quickSort(list, left, j-1)
	s.quickSort(list, j+1, right)
}

// 在list中找到partition，使得list[left:partition]<=list[partition]<=list[partition+1:right]
func (s Quick) partition(list []Comparable, left int, right int) int {
	var (
		p    = left
		i, j = left, right + 1
	)

	for {
		for i++; list[i].CompareTo(list[p]) <= 0 && i < right; i++ {
		}
		for j--; list[j].CompareTo(list[p]) >= 0 && j > left; j-- {
		}
		if i >= j {
			break
		}
		list[i], list[j] = list[j], list[i]
	}
	list[p], list[j] = list[j], list[p] // 此时j<i 取j与p交换
	return j
}
