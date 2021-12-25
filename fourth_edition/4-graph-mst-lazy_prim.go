package fourth_edition

// 最小生成树-Prim算法的延时实现
// 延时：将这些边先留在优先队列中，等到要删除的时候再检查边的有效性
// Prim算法: 每一步都会为一棵生长中的树添加一条边，
// 	每次将下一条连接树中的顶点与不在树中的顶点且权重最小的边加入树中
type LazyPrimMST struct {
	marked []bool     // 最小生成树的顶点
	mst    *edgeQueue // 最小生成树的边
	pq     *MinEdgePQ // 横切边（包括失效的边）
}

func NewLazyPrimMST(g EdgeWeightedGraph) *LazyPrimMST {
	pq := NewMinEdgePQ(nil)
	marked := make([]bool, g.V())
	mst := newEdgeQueue()
	l := &LazyPrimMST{
		mst:    mst,
		marked: marked,
		pq:     pq,
	}
	l.visit(g, 0) // 假设g是连通的
	for !pq.isEmp() {
		e := pq.delMin()
		v := e.Either()
		w := e.Other(v)
		if l.marked[v] && l.marked[w] { // 跳过失效的边
			continue
		}
		l.mst.enqueue(e)
		if !l.marked[v] {
			l.visit(g, v)
		}
		if !l.marked[w] {
			l.visit(g, w)
		}
	}
	return l
}

// 标记顶点v并将所有连接v和未被标记顶点的边加入pq
func (l *LazyPrimMST) visit(g EdgeWeightedGraph, v int) {
	l.marked[v] = true
	for _, e := range g.Edges(v) {
		if !l.marked[e.Other(v)] {
			l.pq.insert(e)
		}
	}
}
