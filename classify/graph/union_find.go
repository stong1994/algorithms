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
