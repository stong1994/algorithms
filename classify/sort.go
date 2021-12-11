package classify

// 排序应用（源于《算法》第2.5章节-应用）
// 1. 找出重复元素——对数组排序，然后遍历一遍即可找到重复元素。（在不考虑内存的情况下，使用hash更快）
// 2. 排名
// 3. 优先队列
// 4. 中位数与顺序统计（如TopN）
// 找到一组数中的第K小的元素，可以通过对数组排序，或者使用优先队列。但是这种只需“部分”数据而对整体进行了排序
// 可以参考快排中的partition方法，当partition为k时，这个位置的元素就是我们想要的元素
func topN(list []int, k int) int {
	k-- // 第k个数的索引为k-1
	left, right := 0, len(list)-1
	for left < right {
		j := quickPartition(list, left, right)
		if j == k {
			return list[j]
		}
		if j > k {
			right = j - 1
		} else {
			left = j + 1
		}
	}
	return list[k]
}

func quickPartition(list []int, left, right int) int {
	i, j, p := left, right+1, left
	for {
		for i++; list[p] <= list[i] && i < right; i++ {
		}
		for j--; list[p] >= list[j] && j > left; j-- {
		}
		if i >= j {
			break
		}
		list[i], list[j] = list[j], list[i]
	}
	list[p], list[j] = list[j], list[p]
	return j
}
