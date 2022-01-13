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
