package array

import "math"

// 把数组中的 0 移到末尾
// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
// https://leetcode-cn.com/problems/move-zeroes/description/
func moveZeroes(nums []int) {
	/*
		// 方法1：遍历nums，找到为0的值，一直将其和相邻元素交换，直到最后一个元素
		for i := 0; i < len(nums); i ++ {
			f := false
			for j := i; j < len(nums); j ++ {
				if nums[j] != 0 {
					f = true
					break
				}
			}
			if !f {
				return
			}
			for nums[i] == 0 { // 避免两个0相邻，交换后i的值还是0，用for循环
				for j := i; j < len(nums)-1; j ++ {
					nums[j], nums[j+1] = nums[j+1], nums[j]
				}
			}
		}
	*/

	/*
		// 方法二，先遍历一遍，找到0的个数，及其索引，第二次遍历，按照当前元素之前0的个数cnt一次性移动cnt位
		list := []int{0} // 对于第i个元素v，i为0的个数，v为对应nums中的位置
		for i, v := range nums {
			if v == 0 {
				list = append(list, i)
			}
		}
		L := len(nums)
		list = append(list, L) // 方便下边计算
		for i := 1; i < len(list)-1; i++ {
			l, r := list[i], list[i+1]
			for j := l + 1; j < r; j++ {
				nums[j-i] = nums[j]
			}
		}
		var totalZero int
		totalZero = len(list) - 2
		for i := 0; i < totalZero; i++ {
			nums[L-1-i] = 0
		}
		// 方法三，双指针：通过设置快慢双指针，如果一直不遇到0，那么快指针和慢指针位置相同，如遇到0，则只移动快指针，并对随后的元素将快指针的值赋值到慢指针上
		// 这能保证慢指针之前都是非0，快慢双指针之间是0
		var slow, fast int
		for ; fast < len(nums); fast++ {
			if nums[fast] != 0 {
				nums[slow] = nums[fast]
				slow++
			}
		}
		for ; slow < len(nums); slow++ {
			nums[slow] = 0
		}
	*/
	// 方法四：对非0字符进行覆盖，用idx记录当前覆盖的进度
	idx := 0
	for _, v := range nums {
		if v != 0 {
			nums[idx] = v
			idx++
		}
	}
	for i := idx; i < len(nums); i++ {
		nums[i] = 0
	}
}

// 改变矩阵维度
// 在 MATLAB 中，有一个非常有用的函数 reshape ，它可以将一个m x n 矩阵重塑为另一个大小不同（r x c）的新矩阵，但保留其原始数据。
//给你一个由二维数组 mat 表示的m x n 矩阵，以及两个正整数 r 和 c ，分别表示想要的重构的矩阵的行数和列数。
//重构后的矩阵需要将原始矩阵的所有元素以相同的 行遍历顺序 填充。
//如果具有给定参数的 reshape 操作是可行且合理的，则输出新的重塑矩阵；否则，输出原始矩阵。
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/reshape-the-matrix
func matrixReshape(mat [][]int, r int, c int) [][]int {
	oldR, oldC := len(mat), len(mat[0])
	if oldR*oldC != r*c {
		return mat
	}
	/*方法一：通过闭包获取每个原数组元素，并依次赋给新数组
	getNext := func() func()int {
		var cr, cc = 0, 0
		return func() int {
			if cc == oldC {
				cr++
				cc = 0
			}
			rst := mat[cr][cc]
			cc++
			return rst
		}
	}
	next := getNext()

	var result [][]int
	for i := 0; i < r; i++ {
		var row []int
		for j := 0; j < c; j++ {
			row = append(row, next())
		}
		result = append(result, row)
	}
	return result
	*/
	// 方法二：找到每个元素位置与二维数组的r与c的关系
	// 对于第i个元素，其对应的行索引为i/c，对应的列索引为i%c
	result := make([][]int, r)
	for i := 0; i < r; i++ {
		result[i] = make([]int, c)
	}
	for i := 0; i < r*c; i++ {
		result[i/c][i%c] = mat[i/oldC][i%oldC]
	}
	return result
}

// 找出数组中最长的连续 1
// 给定一个二进制数组 nums ， 计算其中最大连续 1 的个数。
func findMaxConsecutiveOnes(nums []int) int {
	/*
		curCnt := 0
		maxCnt := 0
		for _, num := range nums {
			if num != 1 {
				maxCnt = max(maxCnt, curCnt)
				curCnt = 0
			}else {
				curCnt++
			}
		}
		maxCnt = max(maxCnt, curCnt)
		return maxCnt
	*/
	// 根据题意：nums[i] 不是 0 就是 1.
	// 对每个元素乘以当前结果并加上当前结果，每次计算后都进行比较获取当前最大值
	curVal := 0
	maxVal := 0
	for _, num := range nums {
		curVal = curVal*num + num
		maxVal = max(maxVal, curVal)
	}
	return maxVal
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 有序矩阵查找
// 编写一个高效的算法来搜索mxn矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：
//每行的元素从左到右升序排列。
//每列的元素从上到下升序排列。
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/search-a-2d-matrix-ii
func searchMatrix(matrix [][]int, target int) bool {
	/*
		// 对每一行进行二分查找
		bs := func(data []int) bool {
			for lo, hi := 0, len(data)-1; lo <= hi; {
				mid := lo + (hi-lo)/2
				if data[mid] == target {
					return true
				}
				if data[mid] > target {
					hi = mid-1
				}else {
					lo = mid+1
				}
			}
			return false
		}
		for _, v := range matrix {
			if bs(v) {
				return true
			}
		}
		return false
	*/
	// 方法二：从左到右、从上到下都是升序，因此可以从右上角或者左下角开始查找，相当于另类的二分查找
	// 假设从左下角开始查找，如果matrix[r][c]小于target，那么其需要变大，因此往右移；否则徐奥变小，往上移
	rMax, cMax := len(matrix), len(matrix[0])
	r, c := rMax-1, 0
	for r >= 0 && c < cMax {
		if matrix[r][c] == target {
			return true
		} else if matrix[r][c] < target {
			c++
		} else {
			r--
		}
	}
	return false
}

// 有序矩阵的 Kth Element
// 给你一个n x n矩阵matrix ，其中每行和每列元素均按升序排序，找到矩阵中第 k 小的元素。
// 请注意，它是 排序后 的第 k 小元素，而不是第 k 个 不同 的元素。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/kth-smallest-element-in-a-sorted-matrix
func kthSmallest(matrix [][]int, k int) int {
	// 对于二维有序矩阵，左上角是最小值，右下角是最大值，从左下角画一条线到右上角可保证右下区域的值都大于左上区域的值，也就是说存在中间元素mid，
	// 左上部分都小于等于mid，右下部份都大于mid
	// 令left表示当前最小值，right表示当前最大值 mid=left+(right-left)/2，cnt表示左上区域的元素数量
	// 若cnt < k，则表示第k小元素位于mid与right之间，即left = mid+1,right = right保证了第k小元素位于left和right之间
	// 若cnt >= k,则表示第k小元素位于left与mid之间，即left = left, right = mid（这里不是mid-1，因为mid可能就是第k小元素）保证第k小元素位于left和right之间
	// 在循环中始终保证第k小元素位于left和right之间，那么当left==right时，left或者right即为第k小元素
	n := len(matrix)
	getCnt := func(mid int) int {
		i, j := n-1, 0 // 从左下角开始
		cnt := 0
		for i >= 0 && j < n {
			if matrix[i][j] <= mid {
				cnt += i + 1
				j++
			} else {
				i--
			}
		}
		return cnt
	}

	left, right := matrix[0][0], matrix[n-1][n-1]
	for left < right {
		mid := left + (right-left)/2
		if getCnt(mid) < k {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 一个数组元素在 [1, n] 之间，其中一个数被替换为另一个数，找出重复的数和丢失的数
// 集合 s 包含从 1 到n的整数。不幸的是，因为数据错误，导致集合里面某一个数字复制了成了集合里面的另外一个数字的值，导致集合 丢失了一个数字 并且 有一个数字重复 。
// 给定一个数组 nums 代表了集合 S 发生错误后的结果。
// 请你找出重复出现的整数，再找到丢失的整数，将它们以数组的形式返回。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/set-mismatch
func findErrorNums(nums []int) []int {
	// 元素在1~n之间，构建一个长度为n+1的数组list，遍历nums，将对应的值作为索引存入到list，如果该索引已有值，说明这就是重复的整数。最后遍历一遍list，找到为空的值
	list := make([]int, len(nums)+1)
	for _, num := range nums {
		list[num]++
	}
	var dup, lost int
	for num := 1; num < len(list); num++ {
		if list[num] == 0 {
			lost = num
		}
		if list[num] == 2 {
			dup = num
		}
	}
	return []int{dup, lost}
}

// 找出数组中重复的数，数组值在 [1, n] 之间
// 给定一个包含n + 1 个整数的数组nums ，其数字都在[1, n]范围内（包括 1 和 n），可知至少存在一个重复的整数。
//假设 nums 只有 一个重复的整数 ，返回这个重复的数 。
//你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/find-the-duplicate-number
func findDuplicate(nums []int) int {
	// 将nums的每个元素放到二进制数组中，再将无重复数的[1,n]放到新的二进制数组中，此时比较两个二进制数组，前者比后者多出来的元素组成的值就是重复值
	// 如[1,2,3,2]对应的二进制数组为[2,3](第0个1有两个，第1个1有3个)，[1,2,3]对应的二进制数组为[2,2]，多出来的二进制元素为10,也就是2
	// 由于只能使用常量级的额外空间，因此不能直接使用数组，只能用for循环来找到每个元素
	n := len(nums)
	maxBit := 31
	for (n-1)>>maxBit == 0 {
		maxBit--
	}
	maxBit++
	var result int
	for bit := 0; bit <= maxBit; bit++ {
		x, y := 0, 0 // x:nums中该位对应的数量，y:[1,n-1]中该位对应的数量
		for i := 0; i < n; i++ {
			// i作为nums的索引
			if nums[i]&(1<<bit) > 0 {
				x++
			}
			// i作为[1,n-1]中的值
			if i >= 1 && (i&(1<<bit)) > 0 {
				y++
			}
		}
		if x > y {
			result |= 1 << bit
		}
	}
	return result
}

// 数组相邻差值的个数
// 给你两个整数 n 和 k ，请你构造一个答案列表 answer ，该列表应当包含从 1 到 n 的 n 个不同正整数，并同时满足下述条件：
// 假设该列表是 answer =[a1, a2, a3, ... , an] ，那么列表 [|a1 - a2|, |a2 - a3|, |a3 - a4|, ... , |an-1 - an|] 中应该有且仅有 k 个不同整数。
// 返回列表 answer 。如果存在多种答案，只需返回其中 任意一种 。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/beautiful-arrangement-ii
func constructArray(n int, k int) []int {
	// 由n可以构成连续的等差数列[1,2,3,..,n]
	// 为了保证相邻元素的差互不相等，可考虑将最大值和最小值放一起，即[1,2,3,..,n]可以变为[1,n,2,n-1,...]
	result := make([]int, n)
	// 要求有k个不同的差值，那么后k+1个元素的差值不同即可，所以前n-(k+1)个元素差值要相同，即为1
	for i := 0; i < n-k-1; i++ {
		result[i] = i + 1
	}
	var (
		j     = 0
		left  = n - k
		right = n
	)
	// 要求有k个不同的差值，那么后k+1个元素的差值不同即可
	// 交替取值，第m个取right，第m+1取left
	for i := n - k - 1; i < n; i++ {
		if j%2 == 0 {
			result[i] = left
			left++
		} else {
			result[i] = right
			right--
		}
		j++
	}
	return result
}

// 数组的度
// 给定一个非空且只包含非负数的整数数组nums，数组的 度 的定义是指数组里任一元素出现频数的最大值。
// 你的任务是在 nums 中找到与nums拥有相同大小的度的最短连续子数组，返回其长度。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/degree-of-an-array
func findShortestSubArray(nums []int) int {
	// 找到度最大的num，然后找到该num在nums中的左右两个索引，差值就是长度
	hash := make(map[int][3]int)
	for i, num := range nums {
		if v, exist := hash[num]; exist {
			hash[num] = [3]int{v[0] + 1, v[1], i}
		} else {
			hash[num] = [3]int{1, i, i}
		}
	}
	var (
		maxLen = math.MaxInt32
		maxCnt = 0
	)
	for _, v := range hash {
		if v[0] > maxCnt {
			maxCnt = v[0]
			maxLen = v[2] - v[1] + 1
		} else if v[0] == maxCnt {
			maxLen = min(maxLen, v[2]-v[1]+1)
		}
	}
	return maxLen
}

// 对角元素相等的矩阵
// 给你一个 m x n 的矩阵 matrix 。如果这个矩阵是托普利茨矩阵，返回 true ；否则，返回 false 。
// 如果矩阵上每一条由左上到右下的对角线上的元素都相同，那么这个矩阵是 托普利茨矩阵 。
// 进阶：
//	如果矩阵存储在磁盘上，并且内存有限，以至于一次最多只能将矩阵的一行加载到内存中，该怎么办？
//	如果矩阵太大，以至于一次只能将不完整的一行加载到内存中，该怎么办？
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/toeplitz-matrix
func isToeplitzMatrix(matrix [][]int) bool {
	// 每个元素都要和其右下相邻的元素相等
	m := len(matrix)
	n := len(matrix[0])
	/*var match func(i, j int) bool
	match = func(i, j int) bool {
		if i >= m-1 || j >= n-1 {
			return true
		}
		return matrix[i][j] == matrix[i+1][j+1] && match(i, j+1) && match(i+1, j) && match(i+1, j+1)
	}
	return match(0, 0)*/
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] != matrix[i-1][j-1] {
				return false
			}
		}
	}
	return true
}

// 嵌套数组
// 索引从0开始长度为N的数组A，包含0到N - 1的所有整数。找到最大的集合S并返回其大小，其中 S[i] = {A[i], A[A[i]], A[A[A[i]]], ... }且遵守以下的规则。
//假设选择索引为i的元素A[i]为S的第一个元素，S的下一个元素应该是A[A[i]]，之后是A[A[A[i]]]... 以此类推，不断添加直到S出现重复的元素。
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/array-nesting
func arrayNesting(nums []int) int {
	/*maxCnt := 0
	used := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		cnt := 0
		for !used[i] {
			used[i] = true
			i = nums[i]
			cnt++
		}
		maxCnt = max(maxCnt, cnt)
	}
	return maxCnt*/
	// 空间优化
	maxCnt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == -1 {
			continue
		}
		cnt := 0
		for j := i; nums[j] != -1; {
			nums[j], j = -1, nums[j]
			cnt++
		}
		maxCnt = max(maxCnt, cnt)
	}
	return maxCnt
}
