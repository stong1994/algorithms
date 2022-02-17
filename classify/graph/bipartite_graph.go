package graph

/* 二分图
定义：若无向图 G=(V,E)G=(V,E) 的顶点集 VV 可以分割为两个互不相交的子集，且图中每条边的两个顶点分别属于不同的子集，则称图 GG 为一个二分图。
判断是否是二分图：
	1. 深度优先搜索/广度优先搜索：从任意顶点开始遍历整个连通域，遍历过程中使用两种不同的颜色对顶点染色，相邻顶点染成相反的颜色。过程中如发现相邻顶点染成了相同的颜色，说明不是二分图。
	2. 并查集：每个顶点的相邻顶点都应在同一集合，且与顶点不在同一集合。将当前顶点的所有连接点合并，并判断这些连接点的邻接点是否和当前顶点处于同一集合中，若是，则不是二分图。
https://leetcode-cn.com/problems/is-graph-bipartite/solution/bfs-dfs-bing-cha-ji-san-chong-fang-fa-pan-duan-er-/
*/

// 判断是否为二分图
// 存在一个 无向图 ，图中有 n 个节点。其中每个节点都有一个介于 0 到 n - 1 之间的唯一编号。
// 给你一个二维数组 graph ，其中 graph[u] 是一个节点数组，由节点 u 的邻接节点组成。
// 形式上，对于graph[u] 中的每个 v ，都存在一条位于节点 u 和节点 v 之间的无向边。该无向图同时具有以下属性：
//	不存在自环（graph[u] 不包含 u）。
//	不存在平行边（graph[u] 不包含重复值）。
//	如果 v 在 graph[u] 内，那么 u 也应该在 graph[v] 内（该图是无向图）
//	这个图可能不是连通图，也就是说两个节点 u 和 v 之间可能不存在一条连通彼此的路径。
//	二分图 定义：如果能将一个图的节点集合分割成两个独立的子集 A 和 B ，并使图中的每一条边的两个节点一个来自 A 集合，一个来自 B 集合，就将这个图称为 二分图 。
// 如果图是二分图，返回 true ；否则，返回 false 。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/is-graph-bipartite
func isBipartite(graph [][]int) bool {
	n := len(graph)
	visit := make([]bool, n)
	color := make([]bool, n)
	ok := true

	var dfs func(i int)
	dfs = func(i int) {
		if !ok {
			return
		}
		visit[i] = true
		for _, j := range graph[i] {
			if !visit[j] {
				color[j] = !color[i]
				dfs(j)
			} else {
				if color[i] == color[j] {
					ok = false
				}
			}
		}
	}

	for i := 0; i < n; i++ { // 不是连通图，所以对每个节点进行遍历
		if !visit[i] {
			dfs(i)
		}
	}
	return ok
}

// 可能的二分法
// 给定一组n人（编号为1, 2, ..., n），我们想把每个人分进任意大小的两组。每个人都可能不喜欢其他人，那么他们不应该属于同一组。
// 给定整数 n和数组 dislikes，其中dislikes[i] = [ai, bi]，表示不允许将编号为 ai和bi的人归入同一组。当可以用这种方法将所有人分进两组时，返回 true；否则返回 false。
// 提示：
//	1 <= n <= 2000
//	0 <= dislikes.length <= 104
//	dislikes[i].length == 2
//	1 <= dislikes[i][j] <= n
//	ai< bi
//	dislikes中每一组都 不同
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/possible-bipartition
func possibleBipartition_dfs(n int, dislikes [][]int) bool {
	// 可以把 dislikes 中每个元组看成结点之间的边，根据结点关系画出图，相邻的结点之间的颜色不能相同。如果我们能够利用两种颜色把所有结点着色就说明可以把这些结点分为两类。
	if len(dislikes) == 0 {
		return true
	}
	colors := make([]bool, n+1)
	visit := make([]bool, n+1)

	dislikeMap := make(map[int][]int)
	for _, d := range dislikes {
		dislikeMap[d[0]] = append(dislikeMap[d[0]], d[1])
		dislikeMap[d[1]] = append(dislikeMap[d[1]], d[0])
	}

	ok := true
	var dfs func(k int)
	dfs = func(k int) {
		if !ok {
			return
		}
		visit[k] = true
		color := colors[k]
		for _, v := range dislikeMap[k] {
			if !visit[v] {
				colors[v] = !color
				dfs(v)
			} else {
				if colors[v] == color {
					ok = false
				}
			}
		}
	}
	for i := 1; i <= n; i++ { // 不是连通图，所以对每个节点进行遍历
		if !visit[i] {
			dfs(i)
		}
	}
	return ok
}
