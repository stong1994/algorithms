package pre_sum

// 重复的DNA序列
// DNA序列由一系列核苷酸组成，缩写为'A','C','G'和'T'.。
// 例如，"ACGAATTCCG"是一个 DNA序列 。
// 在研究 DNA 时，识别 DNA 中的重复序列非常有用。
// 给定一个表示 DNA序列 的字符串 s ，返回所有在 DNA 分子中出现不止一次的长度为10的序列(子字符串)。你可以按 任意顺序 返回答案。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/repeated-dna-sequences
func findRepeatedDnaSequences(s string) []string {
	// 哈希表+前缀和
	var result []string
	N, P := 1e5+10, 131313
	h, p := make([]int, int(N)), make([]int, int(N)) // 哈希数组 次方数组
	n := len(s)
	p[0] = 1
	for i := 1; i <= n; i++ {
		h[i] = h[i-1]*P + int(s[i-1])
		p[i] = p[i-1] * P
	}
	store := make(map[int]int)
	for i := 1; i+9 <= len(s); i++ {
		j := i + 9
		hash := h[j] - h[i-1]*p[j-i+1]
		cnt := store[hash]
		if cnt == 1 {
			result = append(result, s[i-1:j])
		}
		store[hash] = cnt + 1
	}
	return result
}

//  区域和检索 - 数组不可变
// 给定一个整数数组 nums，处理以下类型的多个查询:
// 计算索引left和right（包含 left 和 right）之间的 nums 元素的 和 ，其中left <= right
// 实现 NumArray 类：
// NumArray(int[] nums) 使用数组 nums 初始化对象
// int sumRange(int i, int j) 返回数组 nums中索引left和right之间的元素的 总和 ，
// 包含left和right两点（也就是nums[left] + nums[left + 1] + ... + nums[right])
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/range-sum-query-immutable
type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	sum := make([]int, len(nums)+1)
	sum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sum[i] = sum[i-1] + nums[i]
	}
	return NumArray{sums: sum}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.sums[right]
	}
	return this.sums[right] - this.sums[left-1]
}

// 二维区域和检索 - 矩阵不可变
// 给定一个二维矩阵 matrix，以下类型的多个请求：
// 计算其子矩形范围内元素的总和，该子矩阵的 左上角 为 (row1,col1) ，右下角 为 (row2,col2) 。
// 实现 NumMatrix 类：
// NumMatrix(int[][] matrix)给定整数矩阵 matrix 进行初始化
// int sumRegion(int row1, int col1, int row2, int col2)返回 左上角 (row1,col1)、右下角(row2,col2) 所描述的子矩阵的元素 总和 。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/range-sum-query-2d-immutable
type NumMatrix struct {
	sums [][]int
}

func Constructor2(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	sums := make([][]int, m)
	for i := 0; i < m; i++ {
		sums[i] = make([]int, n)
	}
	sums[0][0] = matrix[0][0]
	for i := 1; i < m; i++ {
		sums[i][0] = sums[i-1][0] + matrix[i][0]
	}
	for i := 1; i < n; i++ {
		sums[0][i] = sums[0][i-1] + matrix[0][i]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			sums[i][j] = sums[i-1][j] + sums[i][j-1] - sums[i-1][j-1] + matrix[i][j]
		}
	}
	return NumMatrix{sums: sums}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if row1 == 0 && col1 == 0 {
		return this.sums[row2][col2]
	}
	if row1 == 0 {
		return this.sums[row2][col2] - this.sums[row2][col1-1]
	}
	if col1 == 0 {
		return this.sums[row2][col2] - this.sums[row1-1][col2]
	}
	return this.sums[row2][col2] - this.sums[row1-1][col2] - this.sums[row2][col1-1] + this.sums[row1-1][col1-1]
}

// 航班预订统计
// 这里有n个航班，它们分别从 1 到 n 进行编号。
// 有一份航班预订表bookings ，表中第i条预订记录bookings[i] = [firsti, lasti, seatsi]
// 意味着在从 firsti到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi个座位。
// 请你返回一个长度为 n 的数组answer，里面的元素是每个航班预定的座位总数。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/corporate-flight-bookings
func corpFlightBookings(bookings [][]int, n int) []int {
	result := make([]int, n)
	//for _, booking := range bookings {
	//	for i := booking[0]; i <= booking[1]; i++ {
	//		result[i-1]+= booking[2]
	//	}
	//}
	// 优化——差分数组（前缀和数组）：如[1,2,2,4]的差分数组为[1,1,0,2]，即第i个值为原数组第i个元素与第i-1个元素的差值
	// 当对第2到3个元素加1时，只需要对差分数组中第二个元素加1，对第4个元素减1即可，即中间的元素不用修改
	for _, booking := range bookings {
		result[booking[0]-1] += booking[2]
		if booking[1] < n {
			result[booking[1]] -= booking[2]
		}
	}
	for i := 1; i < n; i++ {
		result[i] += result[i-1]
	}
	return result
}

// 拼车
