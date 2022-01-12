package classify

// DFS

// 岛屿的最大面积
// 给你一个大小为 m x n 的二进制矩阵 grid 。
// 岛屿是由一些相邻的1(代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在 水平或者竖直的四个方向上 相邻。
// 你可以假设grid 的四个边缘都被 0（代表水）包围着。
// 岛屿的面积是岛上值为 1 的单元格的数目。
// 计算并返回 grid 中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/max-area-of-island
func maxAreaOfIsland(grid [][]int) int {
	ySize := len(grid)
	xSize := len(grid[0])
	occupied := make([][]bool, ySize)
	for i := range occupied {
		occupied[i] = make([]bool, xSize)
	}

	var max = 0
	for y := 0; y < ySize; y++ {
		for x := 0; x < xSize; x++ {
			if occupied[y][x] {
				continue
			}

			if grid[y][x] == 1 {
				area := occupyIsland(grid, occupied, y, x, ySize, xSize)
				if area > max {
					max = area
				}
			}
		}
	}
	return max
}

func occupyIsland(grid [][]int, occupied [][]bool, y, x, ySize, xSize int) int {
	if y >= ySize || x >= xSize || y < 0 || x < 0 || grid[y][x] == 0 {
		return 0
	}
	if occupied[y][x] {
		return 0
	}
	occupied[y][x] = true
	return occupyIsland(grid, occupied, y, x+1, ySize, xSize) + occupyIsland(grid, occupied, y, x-1, ySize, xSize) +
		occupyIsland(grid, occupied, y+1, x, ySize, xSize) + occupyIsland(grid, occupied, y-1, x, ySize, xSize) + 1
}

// 矩阵中的连通分量数目
// 给你一个由'1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
// 此外，你可以假设该网格的四条边均被水包围。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/number-of-islands
func numIslands(grid [][]byte) int {
	ySize := len(grid)
	xSize := len(grid[0])
	occupied := make([][]bool, ySize)
	for i := range occupied {
		occupied[i] = make([]bool, xSize)
	}

	var result int
	for y := 0; y < ySize; y++ {
		for x := 0; x < xSize; x++ {
			if occupied[y][x] {
				continue
			}

			if grid[y][x] == '1' {
				occupyIsland2(grid, occupied, y, x, ySize, xSize)
				result++
			}
		}
	}
	return result
}

func occupyIsland2(grid [][]byte, occupied [][]bool, y, x, ySize, xSize int) {
	if y >= ySize || x >= xSize || y < 0 || x < 0 || grid[y][x] == '0' {
		return
	}
	if occupied[y][x] {
		return
	}
	occupied[y][x] = true
	occupyIsland2(grid, occupied, y, x+1, ySize, xSize)
	occupyIsland2(grid, occupied, y, x-1, ySize, xSize)
	occupyIsland2(grid, occupied, y+1, x, ySize, xSize)
	occupyIsland2(grid, occupied, y-1, x, ySize, xSize)
}

// 好友关系的连通分量数目
// 有 n 个城市，其中一些彼此相连，另一些没有相连。如果城市 a 与城市 b 直接相连，且城市 b 与城市 c 直接相连，那么城市 a 与城市 c 间接相连。
// 省份 是一组直接或间接相连的城市，组内不含其他没有相连的城市。
// 给你一个 n x n 的矩阵 isConnected ，其中 isConnected[i][j] = 1 表示第 i 个城市和第 j 个城市直接相连，
// 而 isConnected[i][j] = 0 表示二者不直接相连。
// 返回矩阵中 省份 的数量。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/number-of-provinces
func findCircleNum(isConnected [][]int) int {
	occupied := make([]bool, len(isConnected))
	result := 0

	var occupyFn func(y int)
	occupyFn = func(y int) {
		if occupied[y] {
			return
		}
		occupied[y] = true
		for i := range isConnected[y] {
			if isConnected[y][i] == 1 {
				occupyFn(i)
			}
		}
	}

	for y := range isConnected {
		if occupied[y] {
			continue
		}
		result++
		occupyFn(y)
	}

	return result
}

// 填充封闭区域
// 给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
// 来源：力扣（LeetCode）
// https://leetcode-cn.com/problems/surrounded-regions
func solve(board [][]byte) {
	// 只要找到连接中的O，剩下的都被X占领
	ySize, xSize := len(board), len(board[0])
	oOccupied := make([][]bool, ySize)
	for i := range oOccupied {
		oOccupied[i] = make([]bool, xSize)
	}

	var occupyFn func(y, x int)
	occupyFn = func(y, x int) {
		if y >= ySize || y < 0 || x >= xSize || x < 0 || oOccupied[y][x] || board[y][x] == 'X' {
			return
		}
		oOccupied[y][x] = true
		occupyFn(y, x+1)
		occupyFn(y, x-1)
		occupyFn(y+1, x)
		occupyFn(y-1, x)
	}

	for x := 0; x < xSize; x++ {
		occupyFn(0, x)
		occupyFn(ySize-1, x)
	}
	for y := 0; y < ySize; y++ {
		occupyFn(y, 0)
		occupyFn(y, xSize-1)
	}
	for y := 0; y < ySize; y++ {
		for x := 0; x < xSize; x++ {
			if !oOccupied[y][x] {
				board[y][x] = 'X'
			}
		}
	}
}

// 能到达的太平洋和大西洋的区域
// 给定一个 m x n 的非负整数矩阵来表示一片大陆上各个单元格的高度。“太平洋”处于大陆的左边界和上边界，而“大西洋”处于大陆的右边界和下边界。
// 规定水流只能按照上、下、左、右四个方向流动，且只能从高到低或者在同等高度上流动。
// 请找出那些水流既可以流动到“太平洋”，又能流动到“大西洋”的陆地单元的坐标。
// 给定下面的 5x5 矩阵:
//
//  太平洋 ~   ~   ~   ~   ~
//       ~  1   2   2   3  (5) *
//       ~  3   2   3  (4) (4) *
//       ~  2   4  (5)  3   1  *
//       ~ (6) (7)  1   4   5  *
//       ~ (5)  1   1   2   4  *
//          *   *   *   *   * 大西洋
//
//返回:
//[[0, 4], [1, 3], [1, 4], [2, 2], [3, 0], [3, 1], [4, 0]] (上图中带括号的单元).
// 输出坐标的顺序不重要
// m 和 n 都小于150
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/pacific-atlantic-water-flow
// TODO
func pacificAtlantic(heights [][]int) [][]int {
	// 反向思考：分别找到大平洋和大西洋能够到达的区域，再获取两者重复的坐标
	// 能够达到的区域的坐标放在map中，因为两个元素都小于150，因此将第1个元素左移6位（2^8=256）再于第二个元素相加即可
	ySize, xSize := len(heights), len(heights[0])

	getTotalValFn := func(y, x int) int {
		return y<<6 + x
	}
	getRealValFn := func(v int) (int, int) {
		return v >> 6, 150 & v
	}

	var visitFn func(y, x int, visited map[int]bool)
	visitFn = func(y, x int, visited map[int]bool) {
		if y >= ySize || y < 0 || x >= xSize || x < 0 {
			return
		}
		v := getTotalValFn(y, x)
		if visited[v] {
			return
		}
		visited[v] = true

		if y < ySize-1 && heights[y][x] <= heights[y+1][x] {
			visitFn(y+1, x, visited)
		}
		if y > 0 && heights[y][x] <= heights[y-1][x] {
			visitFn(y-1, x, visited)
		}
		if x > 0 && heights[y][x] <= heights[y][x-1] {
			visitFn(y, x-1, visited)
		}
		if x < xSize-1 && heights[y][x] <= heights[y][x+1] {
			visitFn(y, x+1, visited)
		}
	}
	m1 := make(map[int]bool)
	m2 := make(map[int]bool)
	var result [][]int
	visitFn(0, 0, m1)
	visitFn(ySize-1, xSize-1, m2)
	for k := range m1 {
		if m2[k] {
			y, x := getRealValFn(k)
			result = append(result, []int{y, x})
		}
	}
	return result
}
