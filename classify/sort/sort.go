package sort

// 煎饼排序
// 给你一个整数数组 arr ，请使用 煎饼翻转 完成对数组的排序。
// 一次煎饼翻转的执行过程如下：
// 选择一个整数 k ，1 <= k <= arr.length
// 反转子数组 arr[0...k-1]（下标从 0 开始）
// 例如，arr = [3,2,1,4] ，选择 k = 3 进行一次煎饼翻转，反转子数组 [3,2,1] ，得到 arr = [1,2,3,4] 。
// 以数组形式返回能使 arr 有序的煎饼翻转操作所对应的 k 值序列。任何将数组排序且翻转次数在10 * arr.length 范围内的有效答案都将被判断为正确。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/pancake-sorting
func pancakeSort(arr []int) []int {
	// 规律：先将最大值翻转到最后，然后再翻转次大值，直到有序(类似于选择排序)
	n := len(arr)
	getCurrIdx := func(k int) int {
		for i := n - 1; i >= 0; i-- {
			if arr[i] == k {
				return i
			}
		}
		return 0
	}
	var result []int
	currVal := n
	for currVal >= 1 {
		idx := getCurrIdx(currVal)
		if idx != currVal-1 {
			// 将currVal翻转到第1个，然后再对前currVal-1个翻转
			if idx != 0 {
				reverse(arr, idx+1)
				result = append(result, idx+1)
			}
			reverse(arr, currVal)
			result = append(result, currVal)
		}
		currVal--
	}
	return result
}

func reverse(arr []int, k int) {
	for l, r := 0, k-1; l < r; l, r = l+1, r-1 {
		arr[l], arr[r] = arr[r], arr[l]
	}
}
