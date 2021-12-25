package fourth_edition

// prim算法是一条边一条边地来构建最小生成树，每一步都为一棵树添加一条边。
// kruskal算法也是一条边一条边地构造，但是它寻找的边会连接成一颗森林地两棵树，然后不断将树合并。
type KruskalMST struct {
	mst *edgeQueue
}

func NewKruskalMST(g EdgeWeightedGraph) *KruskalMST {
	mst := newEdgeQueue()
	pq := NewMinEdgePQ(nil)
	for _, v := range g.adj {
		if len(v) > 0 {
			pq.insert(v[0])
		}
	}

	uf := new(UF)
	uf.Init(g.V())
	for !pq.isEmp() && len(mst.list) <= g.V() {
		e := pq.delMin() // 得到权重最小地边和它地顶点
		v := e.Either()
		w := e.Other(v)
		if uf.Connected(v, w) {
			continue
		}
		uf.Union(v, w) // 合并分量
		mst.enqueue(e) // 将边添加到最小生成树中
	}
	return &KruskalMST{mst: mst}
}
