package fourth_edition

import "fmt"

// 键索引计数法：适用于小整数键的简单排序算法
// 步骤：1. 频率统计 2. 将频率转换为索引 3. 数据分类 4. 回写
type KeyIdxCount struct{}

// a为待排序的数组，元素范围大小为(0，r]，r为最大整数值
func (k KeyIdxCount) sort(a []int, r int) {
	length := len(a)
	aux := make([]int, length+1) // 辅助数组
	count := make([]int, r+2)    // 频率统计数组，索引为对应的元素，值为出现的频率
	// 计算频率 // 假设a=[2,1,1,5,3,3]
	for _, v := range a {
		count[v+1]++ // TODO： 必须为v+1？
	} // count = [0 0 2 1 2 0 1]
	fmt.Println("1", count)
	// 将频率转换为索引
	for i := 0; i < r; i++ {
		count[i+1] += count[i] // 1有2个，2有1个，那么3的起始索引应该是3
	} // count = [0 0 2 3 5 5 1] 最后一个元素应舍弃。这意味着可以得到结果[1,1,2,3,3,5]
	fmt.Println("2", count)
	// 数据分类
	for i := 0; i < length; i++ {
		aux[count[a[i]]] = a[i]
		count[a[i]]++
	}
	// aux = [1 1 2 3 3 5 0]
	fmt.Println("3", aux)
	// 回写
	for i := range a {
		a[i] = aux[i]
	}
}
