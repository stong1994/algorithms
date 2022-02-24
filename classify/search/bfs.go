package search

// BFS
// 广度优先常用于寻找最短路径，最初是一个节点，然后找到下一步能够到达的所有节点，再通过这些节点找到下一个有效节点的列表。
// 以此类推，最先找到最终的节点所走的步长就是最短路径的长度

// 二进制矩阵中的最短路径
// 给你一个 n x n 的二进制矩阵 grid 中，返回矩阵中最短 畅通路径 的长度。如果不存在这样的路径，返回 -1 。
// 二进制矩阵中的 畅通路径 是一条从 左上角 单元格（即，(0, 0)）到 右下角 单元格（即，(n - 1, n - 1)）的路径，
// 该路径同时满足下述要求：
// 路径途经的所有单元格都的值都是 0 。
// 路径中所有相邻的单元格应当在 8 个方向之一 上连通（即，相邻两单元之间彼此不同且共享一条边或者一个角）。
// 畅通路径的长度 是该路径途经的单元格总数。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/shortest-path-in-binary-matrix
func shortestPathBinaryMatrix(grid [][]int) int {
	// 找到最短路径用广度优先，找到可用路径可用深度优先
	sd := NewShortestBFS(grid)
	var (
		n   int
		cur = [][]int{{0, 0}}
	)

	for {
		var valid [][]int
		for _, node := range cur {
			y, x := node[0], node[1]
			if sd.isUsed(y, x) {
				continue
			}
			if !sd.isValid(y, x) {
				continue
			}
			if sd.isEnd(y, x) {
				return n + 1
			}
			sd.markUsed(y, x)
			if sd.bottomRightIsValid(y, x) {
				valid = append(valid, []int{y + 1, x + 1})
			}
			if sd.bottomIsValid(y, x) {
				valid = append(valid, []int{y + 1, x})
			}
			if sd.topIsValid(y, x) {
				valid = append(valid, []int{y - 1, x})
			}
			if sd.rightIsValid(y, x) {
				valid = append(valid, []int{y, x + 1})
			}
			if sd.leftIsValid(y, x) {
				valid = append(valid, []int{y, x - 1})
			}
			if sd.bottomLeftIsValid(y, x) {
				valid = append(valid, []int{y + 1, x - 1})
			}
			if sd.topRightIsValid(y, x) {
				valid = append(valid, []int{y - 1, x + 1})
			}
			if sd.topLeftIsValid(y, x) {
				valid = append(valid, []int{y - 1, x - 1})
			}
		}
		if len(valid) == 0 {
			return -1
		}
		n++
		cur = valid
	}
}

type shortestBFS struct {
	grids [][]int
	used  [][]bool
	xSize int
	ySize int
}

func NewShortestBFS(grids [][]int) *shortestBFS {
	used := make([][]bool, len(grids))
	for i := range used {
		used[i] = make([]bool, len(grids[0]))
	}
	return &shortestBFS{
		grids: grids,
		used:  used,
		xSize: len(grids[0]),
		ySize: len(grids),
	}
}

func (sd *shortestBFS) leftIsValid(y, x int) bool {
	if y >= sd.ySize || x <= 0 {
		return false
	}
	return sd.isValid(y, x-1)
}

func (sd *shortestBFS) rightIsValid(y, x int) bool {
	if y >= sd.ySize || x >= sd.xSize-1 {
		return false
	}
	return sd.isValid(y, x+1)
}

func (sd *shortestBFS) bottomIsValid(y, x int) bool {
	if y >= sd.ySize-1 || x >= sd.xSize {
		return false
	}
	return sd.isValid(y+1, x)
}

func (sd *shortestBFS) topIsValid(y, x int) bool {
	if y <= 0 || x >= sd.xSize {
		return false
	}
	return sd.isValid(y-1, x)
}

func (sd *shortestBFS) bottomRightIsValid(y, x int) bool {
	if y >= sd.ySize-1 || x >= sd.xSize-1 {
		return false
	}
	return sd.isValid(y+1, x+1)
}

func (sd *shortestBFS) bottomLeftIsValid(y, x int) bool {
	if y >= sd.ySize-1 || x <= 0 {
		return false
	}
	return sd.isValid(y+1, x-1)
}

func (sd *shortestBFS) topLeftIsValid(y, x int) bool {
	if y <= 0 || x <= 0 {
		return false
	}
	return sd.isValid(y-1, x-1)
}

func (sd *shortestBFS) topRightIsValid(y, x int) bool {
	if y <= 0 || x >= sd.xSize-1 {
		return false
	}
	return sd.isValid(y-1, x+1)
}

func (sd *shortestBFS) isEnd(y, x int) bool {
	return y == sd.ySize-1 && x == sd.xSize-1
}

func (sd *shortestBFS) isValid(y, x int) bool {
	return sd.grids[y][x] == 0
}

func (sd *shortestBFS) isUsed(y, x int) bool {
	return sd.used[y][x]
}

func (sd *shortestBFS) markUsed(y, x int) {
	sd.used[y][x] = true
}

// 组成整数的最小平方数数量
// 给定正整数n，找到若干个完全平方数（比如1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
// 给你一个整数 n ，返回和为 n 的完全平方数的 最少数量 。
// 完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/perfect-squares
func numSquares(n int) int {
	// 假设平方数a和b，那么b-a>a一定成立：1,4,9,16
	// 因此先找到小于n的最大平方数a，再在[0,a]中找到最大的平方数b，以此类推
	// 但这样找不到最优解，如12按照上述找到的平方数为3,1,1,1 而最优解为2,2,2
	// 按照BFS的思想，对[1, maxSquare]进行层层递进，直到找到和为n的平方数为止. maxSquare为小于一个数的最大的平方数
	// 关键点在于找到剩余的可能节点
	var (
		result  int
		curN    = map[int]struct{}{n: {}}
		handled = make(map[int]struct{})
	)

	for {
		ns := make(map[int]struct{})
		for v := range curN {
			if _, ok := handled[v]; ok {
				continue
			}
			handled[v] = struct{}{}
			s := findMaxSquare(v)
			left := v - s*s
			if left == 0 {
				return result + 1
			}
			for i := 1; i <= s; i++ {
				ns[v-i*i] = struct{}{}
			}
		}
		result++
		curN = ns
	}
}

func findMaxSquare(n int) int {
	// 二分法
	var (
		result int
		lo, hi = 0, n
	)
	for lo <= hi {
		mid := lo + (hi-lo)/2
		rst := mid * mid
		if rst <= n {
			result = mid
			lo = mid + 1
			continue
		}
		hi = mid - 1
	}
	return result
}

// 最短单词路径
// 字典wordList 中从单词 beginWord和 endWord 的 转换序列 是一个按下述规格形成的序列：
// 序列中第一个单词是 beginWord 。
// 序列中最后一个单词是 endWord 。
// 每次转换只能改变一个字母。
// 转换过程中的中间单词必须是字典wordList 中的单词。
// 给你两个单词 beginWord和 endWord 和一个字典 wordList ，找到从beginWord 到endWord 的 最短转换序列 中的 单词数目 。
// 如果不存在这样的转换序列，返回 0。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/word-ladder
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// 所有的字符串的长度都相等
	// BFS: 从beginWord开始，找到其改变一个字符能达到的word

	var (
		curStrs = map[string]struct{}{beginWord: {}}
		handled = make(map[string]struct{})
		result  int
	)

	for len(curStrs) != 0 {
		validStr := make(map[string]struct{})
		for str := range curStrs {
			if equalStr(str, endWord) {
				return result + 1
			}
			for _, v := range wordList {
				if _, ok := handled[v]; ok {
					continue
				}
				if matchStr(str, v) {
					handled[v] = struct{}{}
					validStr[v] = struct{}{}
				}
			}
		}
		result++
		curStrs = validStr
	}
	return 0
}

func matchStr(a, b string) bool {
	var tag bool
	for i := range a {
		if a[i] != b[i] {
			if tag {
				return false
			}
			tag = true
		}
	}
	return tag
}

func equalStr(a, b string) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
