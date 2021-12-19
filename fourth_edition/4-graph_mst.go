package fourth_edition

import (
	"fmt"
	"math"
)

// 最小生成树
// 给定一幅加权无向图，找到它的一颗最小生成树
type IMST interface {
	Edges() []Edge   // 最小生成树的所有边
	Weight() float64 // 最小生成树的权重
}

// 带权重的边的数据类型
type Edge struct {
	v      int     // 顶点
	w      int     // 另一个顶点
	weight float64 // 边的权重
}

func NewEdge(v, w int, weight float64) *Edge {
	return &Edge{
		v:      v,
		w:      w,
		weight: weight,
	}
}

func (e Edge) Weight() float64 {
	return e.weight
}

func (e Edge) Either() int {
	return e.v
}

func (e Edge) Other(v int) int {
	if v == e.v {
		return e.w
	}
	if v == e.w {
		return e.v
	}
	panic("Inconsistent edge")
}

func (e Edge) CompareTo(that *Edge) int {
	if e.Weight() < that.Weight() {
		return -1
	}
	if e.Weight() > that.Weight() {
		return 1
	}
	return 0
}

// 加权无向图
type EdgeWeightedGraph struct {
	v   int       // 顶点总数
	e   int       // 边的总数
	adj [][]*Edge // 邻接表
}

func NewEdgeWeightedGraph(v int) *EdgeWeightedGraph {
	adj := make([][]*Edge, v)
	for i := 0; i < v; i++ {
		adj[i] = []*Edge{}
	}
	return &EdgeWeightedGraph{
		v:   v,
		e:   0,
		adj: adj,
	}
}

func (e EdgeWeightedGraph) V() int {
	return e.v
}

func (e EdgeWeightedGraph) E() int {
	return e.e
}

func (e *EdgeWeightedGraph) AddEdge(edge *Edge) {
	v := edge.Either()
	w := edge.Other(v)
	e.adj[v] = append(e.adj[v], edge)
	e.adj[w] = append(e.adj[w], edge)
	e.e++
}

func (e EdgeWeightedGraph) Edges(v int) []*Edge {
	return e.adj[v]
}

type edgeQueue struct {
	list []*Edge
}

func newEdgeQueue() *edgeQueue {
	return &edgeQueue{}
}

func (q *edgeQueue) enqueue(s *Edge) {
	q.list = append(q.list, s)
}

func (q *edgeQueue) dequeue() *Edge {
	if len(q.list) == 0 {
		return nil
	}
	last := q.list[len(q.list)-1]
	q.list = q.list[:len(q.list)-1]
	return last
}

func (q edgeQueue) isEmpty() bool {
	return len(q.list) == 0
}

type MinEdgePQ struct {
	pq []*Edge
	n  int // 元素个数, 则len(pq)=n+1
}

// 构造堆时，可以将元素依次通过上浮加入队列
// 但是先将元素放到数组中，对数组进行排序，此时可以跳过叶子结点，因为对叶子节点进行下沉是无用的，
// 然后从N/2到1的节点依次进行下沉，这样更快
func NewMinEdgePQ(data []*Edge) *MinEdgePQ {
	if data == nil {
		return &MinEdgePQ{}
	}
	dst := make([]*Edge, len(data)+1)
	copy(dst[1:], data)
	//BaseSort{Quick{}}.Sort(dst[1:])

	m := MinEdgePQ{
		pq: dst,
		n:  len(data),
	}
	for i := m.n / 2; i >= 1; i-- {
		m.sink(i)
	}
	return &m
}

func (m MinEdgePQ) next(k []int) []int {
	var result []int
	for _, v := range k {
		if v*2 <= m.n {
			result = append(result, v*2)
		}
		if v*2+1 <= m.n {
			result = append(result, v*2+1)
		}
		result = append(result, math.MinInt64)
	}
	return result
}

func (m MinEdgePQ) Show() {
	if m.isEmp() {
		fmt.Println("pq is empty")
		return
	}
	fmt.Println(m.pq[1])
	nextK := []int{1}
	for {
		ks := m.next(nextK)
		if len(ks) == 0 {
			break
		}
		nextK = []int{}
		for _, k := range ks {
			if k == math.MinInt64 {
				fmt.Print(" | ")
			} else {
				fmt.Print(m.pq[k], " ")
				nextK = append(nextK, k)
			}
		}
		fmt.Println()
	}
}

func (m MinEdgePQ) isEmp() bool {
	return m.n == 0
}

func (m MinEdgePQ) size() int {
	return m.n
}

// 插入堆时，将元素放入末尾，然后对其进行“上浮”
func (m *MinEdgePQ) insert(k *Edge) {
	m.n++
	m.pq = append(m.pq, k)
	m.swim(m.n)
}

// 删除最小元素时，将堆顶的元素取出，然后将末尾的元素放到堆顶，然后对其进行“下沉”
func (m *MinEdgePQ) delMin() *Edge {
	if m.isEmp() {
		return nil
	}
	max := m.pq[1]
	m.pq[1], m.pq[m.n] = m.pq[m.n], m.pq[1]
	m.pq[m.n] = nil
	m.n--
	m.sink(1)
	return max
}

// 上浮
func (m *MinEdgePQ) swim(k int) {
	for k > 1 && m.pq[k/2].CompareTo(m.pq[k]) > 0 {
		m.pq[k/2], m.pq[k] = m.pq[k], m.pq[k/2]
		k /= 2
	}
}

// 下沉
func (m MinEdgePQ) sink(k int) {
	for k*2 <= m.n {
		j := 2 * k
		if j < m.n && m.pq[j].CompareTo(m.pq[j+1]) > 0 {
			j++
		}
		if m.pq[k].CompareTo(m.pq[j]) <= 0 {
			break
		}
		m.pq[k], m.pq[j] = m.pq[j], m.pq[k]
		k = j
	}
}
