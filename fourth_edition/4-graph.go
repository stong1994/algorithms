package fourth_edition

import "fmt"

type IGraph interface {
	New(v int)        // 创建一个含有V个顶点但不含有边的图
	V() int           // 顶点数
	E() int           // 边数
	AddEdge(v, w int) // 向图中添加一条边 v-w
	Adj(v int) []int  // 和v相邻的所有顶点
	String() string
}

// 计算v的度数
func Degree(graph IGraph, v int) int {
	return len(graph.Adj(v))
}

func MaxDegree(graph IGraph) int {
	max := 0
	for v := 0; v < graph.V(); v++ {
		if Degree(graph, v) > max {
			max = Degree(graph, v)
		}
	}
	return max
}

// 所有顶点的平均度数
func AvgDegree(graph IGraph) float64 {
	return 2.0 * float64(graph.E()) / float64(graph.V())
}

// 计算自环的个数
func NumOfSelfLoops(graph IGraph) int {
	count := 0
	for v := 0; v < graph.V(); v++ {
		for _, w := range graph.Adj(v) {
			if v == w {
				count++
			}
		}
	}
	return count / 2 // todo ?
}

type Graph struct {
	v   int
	e   int
	adj [][]int
}

func (g *Graph) New(v int) {
	g.v = v
	g.e = 0
	for i := 0; i < v; i++ {
		g.adj = append(g.adj, []int{})
	}
}

func (g *Graph) V() int {
	return g.v
}

func (g *Graph) E() int {
	return g.e
}

func (g *Graph) AddEdge(v, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.adj[w] = append(g.adj[w], v)
	g.e++
}

func (g *Graph) Adj(v int) []int {
	return g.adj[v]
}

func (g *Graph) String() string {
	s := fmt.Sprintf("%d veritices, %d edges\n", g.V(), g.E())
	for v := 0; v < g.V(); v++ {
		s += fmt.Sprintf("%d: ", v)
		for _, w := range g.Adj(v) {
			s += fmt.Sprintf("%d: ", w)
		}
		s += "\n"
	}
	return s
}

// DFS: 深度优先搜索
type DepthFirstSearch struct {
	marked []bool // 这个顶点调用过dfs()了吗
	edgeTo []int  // 从起点到一个顶点的已知路径的最后一个顶点
	s      int    // 起点
}

func (d DepthFirstSearch) Search(graph Graph, s int) {
	d.marked = make([]bool, graph.V())
	d.edgeTo = make([]int, graph.V())
	d.s = s
	d.dfs(graph, s)
}

func (d *DepthFirstSearch) dfs(graph Graph, v int) {
	d.marked[v] = true
	for i, w := range graph.Adj(v) {
		if !d.hasMarked(i) {
			d.edgeTo[w] = v
			d.dfs(graph, i)
		}
	}
}

func (d DepthFirstSearch) hasMarked(w int) bool {
	return d.marked[w]
}

// 由根节点s到v的搜索路径
func (d DepthFirstSearch) pathTo(v int) []int {
	if !d.hasMarked(v) {
		return nil
	}
	var result []int
	for i := v; i != d.s; i = d.edgeTo[i] {
		result = append(result, i)
	}
	result = append(result, d.s)

	for i := 0; i < len(result)/2; i++ {
		r := len(result) - 1 - i
		result[i], result[r] = result[r], result[i]
	}
	return result
}

// BFS 广度优先
type BreadthFirstPaths struct {
	marked []bool
	edgeTo []int
	s      int
}

func (b *BreadthFirstPaths) Search(graph Graph, s int) {
	b.marked = make([]bool, graph.V())
	b.edgeTo = make([]int, graph.V())
	b.s = s

}

func (b BreadthFirstPaths) hasMarked(w int) bool {
	return b.marked[w]
}

// 由根节点s到v的搜索路径
func (b BreadthFirstPaths) pathTo(v int) []int {
	if !b.hasMarked(v) {
		return nil
	}
	var result []int
	for i := v; i != b.s; i = b.edgeTo[i] {
		result = append(result, i)
	}
	result = append(result, b.s)

	for i := 0; i < len(result)/2; i++ {
		r := len(result) - 1 - i
		result[i], result[r] = result[r], result[i]
	}
	return result
}

func (b *BreadthFirstPaths) bfs(graph Graph, s int) {
	queue := newQueue()
	b.marked[s] = true
	queue.enqueue(s)
	for !queue.isEmpty() {
		// 取下一个顶点并标记他
		// 将v的所有相邻而又未被标记的顶点加入数据结构
		v := queue.dequeue()
		for _, w := range graph.Adj(v) {
			if b.hasMarked(w) {
				continue
			}
			b.edgeTo[w] = v
			b.marked[w] = true
			queue.enqueue(w)
		}
	}
}

type queue struct {
	list []int
}

func newQueue() *queue {
	return &queue{}
}

func (q *queue) enqueue(s int) {
	q.list = append(q.list, s)
}

func (q *queue) dequeue() int {
	if len(q.list) == 0 {
		return -1
	}
	last := q.list[len(q.list)-1]
	q.list = q.list[:len(q.list)-1]
	return last
}

func (q queue) isEmpty() bool {
	return len(q.list) == 0
}

type queueString struct {
	list []string
}

func newQueueString() *queueString {
	return &queueString{}
}

func (q *queueString) enqueue(s string) {
	q.list = append(q.list, s)
}

func (q *queueString) dequeue() string {
	if len(q.list) == 0 {
		return ""
	}
	last := q.list[len(q.list)-1]
	q.list = q.list[:len(q.list)-1]
	return last
}

func (q queueString) isEmpty() bool {
	return len(q.list) == 0
}
