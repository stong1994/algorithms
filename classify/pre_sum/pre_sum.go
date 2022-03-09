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
// 假设你是一位顺风车司机，车上最初有capacity个空座位可以用来载客。
// 由于道路的限制，车只能向一个方向行驶（也就是说，不允许掉头或改变方向，你可以将其想象为一个向量）。
// 这儿有一份乘客行程计划表trips[][]，其中trips[i] = [num_passengers, start_location, end_location]包含了第 i 组乘客的行程信息：
// - 必须接送的乘客数量；
// - 乘客的上车地点；
// - 以及乘客的下车地点。
// 这些给出的地点位置是从你的初始出发位置向前行驶到这些地点所需的距离（它们一定在你的行驶方向上）。
// 请你根据给出的行程计划表和车子的座位数，来判断你的车是否可以顺利完成接送所有乘客的任务
// （当且仅当你可以在所有给定的行程中接送所有乘客时，返回true，否则请返回 false）。
// 提示：
//	你可以假设乘客会自觉遵守 “先下后上” 的良好素质
//	trips.length <= 1000
//	trips[i].length == 3
//	1 <= trips[i][0] <= 100
//	0 <= trips[i][1] < trips[i][2] <= 1000
//	1 <=capacity <= 100000
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/car-pooling
func carPooling(trips [][]int, capacity int) bool {
	// 差分数组
	// trip的最大长度为1000，因此建立一个长度为1000的差分数组
	// 遍历trips，对上车地点对应的trip增加当前人数，对下车地点对应的trip减去当前人数
	// 最后遍历差分数组，看是否有超过capacity
	list := make([]int, 1001)
	for _, trip := range trips {
		list[trip[1]] += trip[0]
		list[trip[2]] -= trip[0]
	}
	if list[0] > capacity {
		return false
	}
	for i := 1; i < 1001; i++ {
		list[i] += list[i-1]
		if list[i] > capacity {
			return false
		}
	}
	return true
}

// 得分最高的最小轮调
// 给你一个数组nums，我们可以将它按一个非负整数 k 进行轮调，这样可以使数组变为[nums[k], nums[k + 1], ... nums[nums.length - 1],
// nums[0], nums[1], ..., nums[k-1]]的形式。此后，任何值小于或等于其索引的项都可以记作一分。
//例如，数组为nums = [2,4,1,3,0]，我们按k = 2进行轮调后，它将变成[1,3,0,2,4]。这将记为 3 分，
//因为 1 > 0 [不计分]、3 > 1 [不计分]、0 <= 2 [计 1 分]、2 <= 3 [计 1 分]，4 <= 4 [计 1 分]。
//在所有可能的轮调中，返回我们所能得到的最高分数对应的轮调下标 k 。如果有多个答案，返回满足条件的最小的下标 k
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/smallest-rotation-with-highest-score
func bestRotation(nums []int) int {
	/* 暴力破解
	var result, maxScore int
	n := len(nums)
	// 0<=k<n
	score := 0
	for i := 0; i < n; i++ {
		if nums[i] <= i {
			score++
		}
	}
	maxScore = score
	for k := 1; k < n; k++ {
		first := nums[0]
		score = 0
		for i := 1; i < n; i++ {
			nums[i-1] = nums[i]
			if nums[i-1] <= i-1 {
				score++
			}
		}
		nums[n-1] = first
		if first <= n-1 {
			score++
		}
		if score > maxScore {
			maxScore = score
			result = k
		}
	}
	return result
	*/
	/*
		找规律
		n := len(nums)
		ks := make([]int, n)
		for i, num := range nums {
			// 找到符合num的k区间：向左移动k位后 num <= i-k => k >= num-i
			if i >= num{
				// 当i大于num时，此时num可以向左移动到索引为num的j处，也可以向右移动到最右边
				for j := 0; j <=i-num; j++ {
					ks[j]++
				}
				for j := i+1; j < n; j++ {
					ks[j]++
				}
			}else if i < num {
				// i < num时，需要把num移到i的右边
				//l := num-i
				//r := n-1-i
				// 此时k为n-l 到 n-r
				for k := i+1; k <= n-num+i; k++ {
					ks[k]++
				}
			}
		}
		//fmt.Println(ks)
		maxK := 0
		maxCnt := 0
		for k, cnt := range ks {
			if cnt > maxCnt {
				maxK = k
				maxCnt = cnt
			}
		}
		return maxK
	*/
	// 差分数组
	n := len(nums)
	ks := make([]int, n)
	for i, num := range nums {
		// 找到符合num的k区间：向左移动k位后 num <= i-k => k >= num-i
		if i >= num {
			// 当i大于num时，此时num可以向左移动到索引为num的j处，也可以向右移动到最右边
			ks[0]++
			if i-num+1 < n {
				ks[i-num+1]--
			}
			//for j := 0; j <=i-num; j++ {
			//	ks[j]++
			//}
			//for j := i+1; j < n; j++ {
			//	ks[j]++
			//}
			if i+1 < n {
				ks[i+1]++
			}
		} else if i < num {
			// i < num时，需要把num移到i的右边
			//l := num-i
			//r := n-1-i
			// 此时k为n-l 到 n-r
			//for k := i+1; k <= n-num+i; k++ {
			//	ks[k]++
			//}
			if i+1 < n {
				ks[i+1]++
			}
			if n-num+i+1 < n {
				ks[n-num+i+1]--
			}
		}
	}
	//fmt.Println(ks)
	maxK := 0
	score, maxScore := 0, 0
	for k, diff := range ks {
		score += diff
		if score > maxScore {
			maxK = k
			maxScore = score
		}
	}
	return maxK
}
