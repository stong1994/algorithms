package classify

import (
	"math/rand"
)

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
	p := rand.Intn(right-left+1) + left // 随机选择一个元素
	pivot := list[p]
	list[p], list[left] = list[left], list[p] // 将随机选择的元素放到第一个
	i, j := left, right+1
	for {
		for i++; i < right && list[i] >= pivot; i++ {
		}
		for j--; j > left && list[j] <= pivot; j-- {
		}
		if i >= j {
			break
		}
		list[i], list[j] = list[j], list[i]
	}
	list[left], list[j] = list[j], list[left]
	return j
}

// 以下源自：https://github.com/CyC2018/CS-Notes/blob/master/notes/Leetcode%20%E9%A2%98%E8%A7%A3%20-%20%E6%8E%92%E5%BA%8F.md
// 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/description/
func findKthLargest(nums []int, k int) int {
	k-- // 第k个数的索引为k-1
	left, right := 0, len(nums)-1
	for left < right {
		j := quickPartition(nums, left, right)
		if j == k {
			return nums[j]
		}
		if j > k {
			right = j - 1
		} else {
			left = j + 1
		}
	}
	return nums[k]
}

// 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
// https://leetcode-cn.com/problems/top-k-frequent-elements/description/
// 转换为TopN问题：构造结构体{num, count}，然后将nums数组转换为新的结构体数组，然后进行快速排序找到partition
func topKFrequent_0(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}

	type node struct {
		num   int
		count int
	}

	nodeQuickPartition := func(list []node, left, right int) int {
		p := rand.Intn(right-left+1) + left // 随机选择一个元素
		pivot := list[p].count
		list[p], list[left] = list[left], list[p] // 将随机选择的元素放到第一个
		i, j := left, right+1
		for {
			for i++; i < right && list[i].count >= pivot; i++ {
			}
			for j--; j > left && list[j].count <= pivot; j-- {
			}
			if i >= j {
				break
			}
			list[i], list[j] = list[j], list[i]
		}
		list[left], list[j] = list[j], list[left]
		return j
	}

	m := make(map[int]int)
	arr := make([]node, 0, k)
	for _, v := range nums {
		if r, ok := m[v]; ok {
			m[v] = r + 1
			continue
		}
		m[v] = 1
	}

	if len(m) <= k {
		return mapKey2Arr(m)
	}

	for num, count := range m {
		arr = append(arr, node{
			num:   num,
			count: count,
		})
	}

	getMaxKNode := func() []node {
		k-- // 第k个数的索引为k-1
		left, right := 0, len(arr)-1
		for left < right {
			j := nodeQuickPartition(arr, left, right)
			if j == k {
				return arr[:j+1]
			}
			if j > k {
				right = j - 1
			} else {
				left = j + 1
			}
		}
		return arr[:k+1]
	}

	ns := getMaxKNode()
	result := make([]int, len(ns))
	for i, v := range ns {
		result[i] = v.num
	}
	return result
}

func mapKey2Arr(m map[int]int) []int {
	result := make([]int, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
}

// 解法2：桶排序
func topKFrequent(nums []int, k int) []int {
	arr := []map[int]struct{}{0: {}, 1: {}} // 数量-值
	m := make(map[int]int)
	for _, v := range nums {
		if n, ok := m[v]; ok {
			m[v] = n + 1
			delete(arr[n], v)
			if len(arr) == n+1 {
				arr = append(arr, map[int]struct{}{v: {}})
			} else if arr[n+1] == nil {
				arr[n+1] = map[int]struct{}{v: {}}
			} else {
				arr[n+1][v] = struct{}{}
			}
			continue
		}
		m[v] = 1
		arr[1][v] = struct{}{}
	}
	var result []int
	for i := len(arr) - 1; i >= 1; i-- {
		for n := range arr[i] {
			if len(result) < k {
				result = append(result, n)
			}
		}
	}
	return result
}

// 给定一个字符串，请将字符串里的字符按照出现的频率降序排列。
// https://leetcode-cn.com/problems/sort-characters-by-frequency/description/
// 基于桶排序
func frequencySort(s string) string {
	if s == "" {
		return ""
	}
	arr := []map[int32]struct{}{0: {}, 1: {}} // 数量-值
	m := make(map[int32]int)
	for _, v := range s {
		if n, ok := m[v]; ok {
			m[v] = n + 1
			delete(arr[n], v)
			if len(arr) == n+1 {
				arr = append(arr, map[int32]struct{}{v: {}})
			} else if arr[n+1] == nil {
				arr[n+1] = map[int32]struct{}{v: {}}
			} else {
				arr[n+1][v] = struct{}{}
			}
			continue
		}
		m[v] = 1
		arr[1][v] = struct{}{}
	}
	var result []byte
	for i := len(arr) - 1; i >= 1; i-- {
		for n := range arr[i] {
			for j := 0; j < i; j++ {
				result = append(result, byte(n))
			}
		}
	}
	return string(result)
}

// 荷兰国旗问题
// 给定一个包含红色、白色和蓝色，一共n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
// 此题中，我们使用整数 0、1 和 2 分别表示红色、白色和蓝色。
// https://leetcode-cn.com/problems/sort-colors/description/
// 桶排序
func sortColors(nums []int) {
	arr := []int{0, 0, 0}
	for _, v := range nums {
		arr[v]++
	}
	i := 0
	for n, v := range arr {
		for j := 0; j < v; j++ {
			nums[i] = n
			i++
		}
	}
}
