package classify

import (
	"sort"
	"strconv"
	"strings"
)

// 电话号码的字母组合
// 给定一个仅包含数字2-9的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
// 提示：
// 0 <= digits.length <= 4
// digits[i] 是范围 ['2', '9'] 的一个数字。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number
func letterCombinations(digits string) []string {
	data := map[int32][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	result := []string{""}
	for _, dv := range digits {
		var list []string
		for _, rv := range result {
			for _, lv := range data[dv] {
				list = append(list, rv+lv)
			}
		}
		result = list
	}
	if len(result) == 1 {
		return []string{}
	}
	return result
}

// IP 地址划分
// 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
// 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
// 给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入'.' 来形成。
// 你不能重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/restore-ip-addresses
func restoreIpAddresses(s string) []string {
	ips := ipDfs(s, 0, nil)
	result := make([]string, 0, len(ips))
	for _, v := range ips {
		result = append(result, strings.Join(v, "."))
	}
	return result
}

func ipDfs(s string, idx int, curIp []string) [][]string {
	if idx == len(s) {
		if len(curIp) == 4 {
			return [][]string{curIp}
		}
		return nil
	}

	// 前导为0
	if s[idx] == '0' {
		return ipDfs(s, idx+1, append(curIp, "0"))
	}
	var result [][]string
	for i := idx; i < len(s); i++ {
		ip := s[idx : i+1]
		// 不能大于255
		ipNum, _ := strconv.Atoi(ip)
		if ipNum > 255 {
			break
		}
		for _, v := range ipDfs(s, i+1, append(curIp, ip)) {
			result = append(result, v)
		}
	}
	return result
}

// 在矩阵中寻找字符串
// 给定一个m x n 二维字符网格board 和一个字符串单词word 。如果word 存在于网格中，返回 true ；否则，返回 false 。
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
// 1 <= m, n <= 6
// 1 <= word.length <= 15
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/word-search
func exist(board [][]byte, word string) bool {
	ySize, xSize := len(board), len(board[0])

	occupied := make([][]bool, ySize)
	for i := range occupied {
		occupied[i] = make([]bool, xSize)
	}

	var dfsFn func(words []byte, y, x int) bool
	dfsFn = func(words []byte, y, x int) bool {
		if y >= ySize || y < 0 || x >= xSize || x < 0 || occupied[y][x] {
			return false
		}
		if len(words) == 0 {
			return true
		}
		occupied[y][x] = true
		target := words[0]
		if y < ySize-1 && board[y+1][x] == target && dfsFn(words[1:], y+1, x) {
			return true
		}
		if y > 0 && board[y-1][x] == target && dfsFn(words[1:], y-1, x) {
			return true
		}
		if x > 0 && board[y][x-1] == target && dfsFn(words[1:], y, x-1) {
			return true
		}
		if x < xSize-1 && board[y][x+1] == target && dfsFn(words[1:], y, x+1) {
			return true
		}
		occupied[y][x] = false
		return false
	}

	for y, v := range board {
		for x, m := range v {
			if m == word[0] {
				if dfsFn([]byte(word)[1:], y, x) {
					return true
				}
			}
		}
	}
	return false
}

// 二叉树的所有路径
// 给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。
// 叶子节点 是指没有子节点的节点。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/binary-tree-paths
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	var dfsFn func(root *TreeNode, curList []string) [][]string
	dfsFn = func(root *TreeNode, curList []string) [][]string {
		curList = append(curList, strconv.Itoa(root.Val)) // 注意：append如果不需要扩容，则复用底层数组,因此在dfs时需要进行copy
		if root.Left == nil && root.Right == nil {
			return [][]string{curList}
		}

		var result [][]string
		if root.Left != nil {
			newList := make([]string, len(curList))
			copy(newList, curList)
			for _, v := range dfsFn(root.Left, newList) {
				result = append(result, v)
			}
		}
		if root.Right != nil {
			newList := make([]string, len(curList))
			copy(newList, curList)
			for _, v := range dfsFn(root.Right, newList) {
				result = append(result, v)
			}
		}

		return result
	}
	data := dfsFn(root, nil)
	result := make([]string, 0, len(data))
	for _, v := range data {
		result = append(result, strings.Join(v, "->"))
	}
	return result
}

// 全排列
// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/permutations
func permute(nums []int) [][]int {
	var dfsFn func(occupied []bool, data []int) [][]int
	dfsFn = func(occupied []bool, data []int) [][]int {
		if len(data) == len(nums) {
			return [][]int{data}
		}
		var result [][]int
		for i := range nums {
			if occupied[i] {
				continue
			}
			occupied[i] = true

			for _, v := range dfsFn(occupied, append(data, nums[i])) {
				result = append(result, v)
			}
			occupied[i] = false
		}
		return result
	}

	occupied := make([]bool, len(nums))
	return dfsFn(occupied, nil)
}

// 可将一些变量全局化，降低内存占用
func permuteOpt(nums []int) [][]int {
	var (
		backTrack func(idx int)
		n         = len(nums)
		list      []int
		result    [][]int
		occupied  = make([]bool, n)
	)
	backTrack = func(idx int) {
		if idx == n {
			tmp := make([]int, len(list))
			copy(tmp, list)
			result = append(result, tmp)
			return
		}
		for i := range nums {
			if occupied[i] {
				continue
			}
			occupied[i] = true
			list = append(list, nums[i])
			backTrack(idx + 1)
			occupied[i] = false
			list = list[:len(list)-1]
		}
	}
	backTrack(0)
	return result
}

// 全排列 II
// 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/permutations-ii
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var perm []int
	var result [][]int
	visited := make([]bool, n)
	var backTrack func(idx int)
	backTrack = func(idx int) {
		if idx == n {
			result = append(result, append([]int(nil), perm...))
			return
		}
		for i, v := range nums {
			if visited[i] || i > 0 && !visited[i-1] && v == nums[i-1] {
				continue
			}
			perm = append(perm, v) // 每加一个元素，就在后边减去一个元素，保证每次循环时，perm都是相同的
			visited[i] = true
			backTrack(idx + 1)
			visited[i] = false
			perm = perm[:len(perm)-1]
		}
	}
	backTrack(0)
	return result
}

// 组合
// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
// 你可以按 任何顺序 返回答案。
// 示例 1：
// 输入：n = 4, k = 2
// 输出：
// [
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
// ]
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/combinations
func combine(n int, k int) [][]int {
	var backTrack func(int, []int) [][]int
	backTrack = func(idx int, list []int) [][]int {
		if len(list) == k {
			return [][]int{list}
		}
		var rst [][]int
		for i := idx; i <= n; i++ {
			tmp := make([]int, len(list))
			copy(tmp, list)
			tmp = append(tmp, i)
			for _, v := range backTrack(i+1, tmp) {
				rst = append(rst, v)
			}
		}
		return rst
	}
	return backTrack(1, nil)
}

func combineOpt(n int, k int) [][]int {
	var (
		result [][]int
		path   []int
	)
	var backTrack func(int)
	backTrack = func(startIdx int) {
		if len(path) == k {
			result = append(result, append([]int{}, path...))
			return
		}

		for i := startIdx; i <= n; i++ {
			if n-i+1 < k-len(path) { // 剩余元素不足
				break
			}
			path = append(path, i)
			backTrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backTrack(1)
	return result
}

// 组合总和
// 给你一个 无重复元素 的整数数组candidates 和一个目标整数target，找出candidates中可以使数字和为目标数target 的 所有不同组合 ，
// 并以列表形式返回。你可以按 任意顺序 返回这些组合。
// candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
// 对于给定的输入，保证和为target 的不同组合数少于 150 个。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/combination-sum
func combinationSum(candidates []int, target int) [][]int {
	var (
		result [][]int
		path   []int
		sum    int
	)
	var backTrack func(int)
	backTrack = func(idx int) {
		if sum > target {
			return
		}
		if sum == target {
			result = append(result, append([]int{}, path...))
			return
		}
		for i := idx; i < len(candidates); i++ {
			sum += candidates[i]
			path = append(path, candidates[i])
			backTrack(i)
			path = path[:len(path)-1]
			sum -= candidates[i]
		}
	}
	backTrack(0)
	return result
}

// 含有相同元素的组合求和
// 给你一个由候选元素组成的集合 candidates和一个目标数target，找出candidates中所有可以使数字和为target的组合。
// candidates 中的每个元素在每个组合中只能使用 一次 。
// 注意：解集不能包含重复的组合。
// 1 <= candidates.length <= 100
// 1 <= candidates[i] <= 50
// 1 <= target <= 30
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/combination-sum-ii
func combinationSum2(candidates []int, target int) [][]int {
	// 难点在于解集中不能包含重复的组合，如对target为2的candidates [2,2]，结果为{{2}}而不是{{2},{2}}
	// 普通的回溯法做不到这一点。
	// 可以对每个元素进行计数统计，对该数进行统计时，按照次数进行递归。这样，可以根据元素值进行occupied统计，只遍历没有被occupied的数据
	var (
		result    [][]int
		path      []int
		sum       int
		backTrack func(int)
		occupied  = make([]bool, 51)
		times     = make(map[int]int)
	)
	sort.Ints(candidates)
	for _, v := range candidates {
		times[v]++
	}

	backTrack = func(idx int) {
		if sum > target {
			return
		}
		if sum == target {
			result = append(result, append([]int{}, path...))
			return
		}
		for i := idx; i < len(candidates); i++ {
			v := candidates[i]
			if occupied[v] {
				continue
			}
			if i > 0 && candidates[i-1] == v {
				continue
			}
			occupied[v] = true

			for n := 1; n <= times[v]; n++ {
				sum += v * n
				for t := 0; t < n; t++ {
					path = append(path, v)
				}
				backTrack(i + 1)
				path = path[:len(path)-n]
				sum -= v * n
			}

			occupied[v] = false
		}

	}
	backTrack(0)
	return result
}

func combinationSum2Opt(candidates []int, target int) [][]int {
	// 难点在于组合是不可重复的。
	// 观察可知，如果candidates是有序的，
	// 那么 当对第startIdx个元素进行回溯时，如果第startIdx+n与第startIdx+n-1相等时，取消掉第startIdx+n的回溯 即可
	sort.Ints(candidates)
	var (
		result    [][]int
		sum       int
		path      []int
		backTrack func(int)
	)
	backTrack = func(startIdx int) {
		if sum > target {
			return
		}
		if sum == target {
			result = append(result, append([]int{}, path...))
			return
		}
		for i := startIdx; i < len(candidates); i++ {
			if i > startIdx && candidates[i-1] == candidates[i] {
				continue
			}
			sum += candidates[i]
			path = append(path, candidates[i])
			backTrack(i + 1)
			path = path[:len(path)-1]
			sum -= candidates[i]
		}
	}
	backTrack(0)
	return result
}

// 1-9 数字的组合求和
// 找出所有相加之和为n 的k个数的组合。组合中只允许含有 1 -9 的正整数，并且每种组合中不存在重复的数字。
// 说明：
// 所有数字都是正整数。
// 解集不能包含重复的组合。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/combination-sum-iii
func combinationSum3(k int, n int) [][]int {
	var (
		result    [][]int
		sum       int
		path      []int
		backTrack func(int)
	)
	backTrack = func(startIdx int) {
		if sum == n {
			if len(path) != k {
				return
			}
			result = append(result, append([]int{}, path...))
			return
		}
		if sum > n {
			return
		}
		if len(path) >= k {
			return
		}
		for i := startIdx; i <= 9; i++ {
			sum += i
			path = append(path, i)
			backTrack(i + 1)
			path = path[:len(path)-1]
			sum -= i
		}
	}
	backTrack(1)
	return result
}

// 子集
// 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
func subsets(nums []int) [][]int {
	var (
		result   = [][]int{}
		occupied = make([]bool, len(nums))
		path     []int
	)
	var backTrack func(int)
	backTrack = func(startIdx int) {
		result = append(result, append([]int{}, path...))
		for i := startIdx; i < len(nums); i++ {
			if occupied[i] {
				continue
			}
			occupied[i] = true
			path = append(path, nums[i])
			backTrack(i + 1)
			path = path[:len(path)-1]
			occupied[i] = false

		}
	}
	backTrack(0)
	return result
}

// 含有相同元素求子集
// 给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
// 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/subsets-ii
func subsetsWithDup(nums []int) [][]int {
	var (
		result    [][]int
		path      []int
		backTrack func(int)
	)
	sort.Ints(nums)
	backTrack = func(startIdx int) {
		result = append(result, append([]int{}, path...))
		for i := startIdx; i < len(nums); i++ {
			if i > startIdx && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			backTrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backTrack(0)
	return result
}

// 分割字符串使得每个部分都是回文数*****
// 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
// 回文串 是正着读和反着读都一样的字符串。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/palindrome-partitioning
func partition(s string) [][]string {
	var (
		result    [][]string
		path      []string
		backTrack func(int)
	)

	isPalindrome := func(str string) bool {
		i, j := 0, len(str)-1
		for {
			if i > j {
				return true
			}
			if str[i] != str[j] {
				return false
			}
			i++
			j--
		}
	}

	backTrack = func(startIdx int) {
		if startIdx == len(s) {
			result = append(result, append([]string{}, path...))
			return
		}

		for i := startIdx + 1; i <= len(s); i++ {
			if !isPalindrome(s[startIdx:i]) {
				continue
			}
			path = append(path, s[startIdx:i])
			backTrack(i)
			path = path[:len(path)-1]
		}
	}
	backTrack(0)
	return result
}

// 解数独*****
// 编写一个程序，通过填充空格来解决数独问题。
// 数独的解法需 遵循如下规则：
// 数字1-9在每一行只能出现一次。
// 数字1-9在每一列只能出现一次。
// 数字1-9在每一个以粗实线分隔的3x3宫内只能出现一次。（请参考示例图）
// 数独部分空格内已填入了数字，空白格用'.'表示。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/sudoku-solver
func solveSudoku(board [][]byte) {
	// 隐藏条件：棋盘是一个9x9的网格(可以简化初始化步骤)
	var (
		column = [9][9]bool{}    // 每列已占有的数字
		row    = [9][9]bool{}    // 每行已占有的数字
		grid   = [3][3][9]bool{} // 每个9宫格已占有的数字
		blank  [][2]int          // 空格，即'.'所在宫格
	)

	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			n := board[y][x]
			if n == '.' {
				blank = append(blank, [2]int{y, x})
				continue
			}
			d := n - '1'
			column[x][d] = true
			row[y][d] = true
			grid[y/3][x/3][d] = true
		}
	}

	var backTrack func(idx int) bool
	backTrack = func(idx int) bool {
		if idx == len(blank) {
			return true
		}
		y, x := blank[idx][0], blank[idx][1]
		for j := 0; j <= 8; j++ {
			if !column[x][j] && !row[y][j] && !grid[y/3][x/3][j] {
				column[x][j] = true
				row[y][j] = true
				grid[y/3][x/3][j] = true
				board[y][x] = byte(j + '1')

				if backTrack(idx + 1) {
					return true
				}

				column[x][j] = false
				row[y][j] = false
				grid[y/3][x/3][j] = false

			}
		}
		return false
	}
	backTrack(0)
}

// N 皇后
// n皇后问题 研究的是如何将 n个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
// 给你一个整数 n ，返回所有不同的n皇后问题 的解决方案。
// 每一种解法包含一个不同的n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/n-queens
func solveNQueens(n int) [][]string {
	// 把皇后看成象棋中的“车”（比车厉害，因为可以斜着吃子），横竖斜可以直接吃子，
	// 因此皇后的横竖斜不能有其他皇后，观察可知每一行和每一列都必须有皇后
	var (
		column = make([]bool, n)
		//row    = make([]bool, n)
		tmp    [][]string
		result [][]string
	)
	// 初始化
	for i := 0; i < n; i++ {
		tmp = append(tmp, make([]string, n))
	}

	biasIsValid := func(y, x int) bool {
		for j, i := y-1, x-1; j >= 0 && i >= 0; j, i = j-1, i-1 {
			if tmp[j][i] == "Q" {
				return false
			}
		}
		for j, i := y+1, x+1; j < n && i < n; j, i = j+1, i+1 {
			if tmp[j][i] == "Q" {
				return false
			}
		}
		for j, i := y+1, x-1; j < n && i >= 0; j, i = j+1, i-1 {
			if tmp[j][i] == "Q" {
				return false
			}
		}
		for j, i := y-1, x+1; j >= 0 && i < n; j, i = j-1, i+1 {
			if tmp[j][i] == "Q" {
				return false
			}
		}
		return true
	}

	add2Rst := func() {
		var list []string
		for _, v := range tmp {
			for i, vv := range v {
				if vv != "Q" {
					v[i] = "."
				}
			}
			list = append(list, strings.Join(v, ""))
		}
		result = append(result, list)
	}

	var backTrack func(int)
	backTrack = func(idx int) {
		if idx == n {
			add2Rst()
			return
		}
		for x := 0; x < n; x++ {
			if column[x] {
				continue
			}
			if !biasIsValid(idx, x) {
				continue
			}
			column[x] = true
			tmp[idx][x] = "Q"
			backTrack(idx + 1)
			tmp[idx][x] = "."
			column[x] = false
		}
	}

	backTrack(0)
	return result
}
