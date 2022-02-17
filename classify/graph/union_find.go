package graph

// 并查集
// 并查集可以动态地连通两个点，并且可以非常快速地判断两个点是否连通。

// 冗余连接
// 树可以看成是一个连通且 无环的无向图。
// 给定往一棵n 个节点 (节点值1～n) 的树中添加一条边后的图。添加的边的两个顶点包含在 1 到 n中间，且这条附加的边不属于树中已存在的边。
// 图的信息记录于长度为 n 的二维数组 edges，edges[i] = [ai, bi]表示图中在 ai 和 bi 之间存在一条边。
// 请找出一条可以删去的边，删除后可使得剩余部分是一个有着 n 个节点的树。如果有多个答案，则返回数组edges中最后出现的边。
// 提示:
//	n == edges.length
//	3 <= n <= 1000
//	edges[i].length == 2
//	1 <= ai< bi<= edges.length
//	ai != bi
//	edges 中无重复元素
//	给定的图是连通的
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/redundant-connection
func findRedundantConnection(edges [][]int) []int {
	// 将父节点和子节点join，如果在join前就已经连通，则说明需要移除
	// 需要的操作是判断是否连通、合并集合，可以使用并查集
	n := len(edges)
	uf := NewUnionFind(n + 1)
	for _, edge := range edges {
		if uf.isConnected(edge[0], edge[1]) {
			return edge
		}
		uf.join(edge[0], edge[1])
	}
	return nil
}

type unionFind struct {
	nodes []int
}

func NewUnionFind(n int) *unionFind {
	nodes := make([]int, n)
	for i := 0; i < n; i++ {
		nodes[i] = i
	}
	return &unionFind{nodes}
}

func (uf *unionFind) find(p int) int {
	if p == uf.nodes[p] {
		return p
	}
	return uf.find(uf.nodes[p])
}

func (uf *unionFind) join(p, q int) {
	uf.nodes[uf.find(q)] = uf.nodes[uf.find(p)]
}

func (uf *unionFind) isConnected(p, q int) bool {
	return uf.find(p) == uf.find(q)
}

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
func isBipartite_uf(graph [][]int) bool {
	type unionFind struct {
		roots []int // 第i个节点的连通根节点 即i与roots[i]是连通的
	}
	newUF := func(n int) *unionFind {
		r := make([]int, n)
		for i := 0; i < n; i++ {
			r[i] = i
		}
		return &unionFind{roots: r}
	}
	// 寻找k的连通根节点
	var find func(uf *unionFind, k int) int
	find = func(uf *unionFind, k int) int {
		if uf.roots[k] == k {
			return k
		}
		return find(uf, uf.roots[k])
	}
	// 判断pq是否在同一集合
	isConnected := func(uf *unionFind, p, q int) bool {
		return find(uf, p) == find(uf, q)
	}
	// 合并q所在的集合到p所在的集合
	union := func(uf *unionFind, p, q int) {
		uf.roots[find(uf, q)] = find(uf, p)
	}

	// 初始化并查集
	uf := newUF(len(graph))
	// 遍历每个顶点，将当前顶点的所有邻接点进行合并
	for i, adjs := range graph {
		for _, w := range adjs {
			// 若某个邻接点与当前顶点已经在一个集合中了，说明不是二分图，返回 false。
			if isConnected(uf, i, w) {
				return false
			}
			union(uf, adjs[0], w)
		}
	}
	return true
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
func possibleBipartition(n int, dislikes [][]int) bool {
	// 将dislikes中的元素构成图，判断该图是否是二分图。
	// 将顶点的相邻节点进行连接，如果是二分图，则每个根节点都不会与相邻节点连接。
	adjs := make([][]int, n+1)
	for _, d := range dislikes {
		adjs[d[0]] = append(adjs[d[0]], d[1])
		adjs[d[1]] = append(adjs[d[1]], d[0])
	}
	uf := NewUnionFind(n + 1)
	for parent, childs := range adjs {
		for _, child := range childs {
			if uf.isConnected(parent, child) {
				return false
			}
			uf.join(childs[0], child)
		}
	}
	return true
}
