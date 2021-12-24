package fourth_edition

import "fmt"

type DirectedGraph struct {
	v   int
	e   int
	adj [][]int
}

func (g *DirectedGraph) New(v int) {
	g.v = v
	g.e = 0
	for i := 0; i < v; i++ {
		g.adj = append(g.adj, []int{})
	}
}

func (g *DirectedGraph) V() int {
	return g.v
}

func (g *DirectedGraph) E() int {
	return g.e
}

func (g *DirectedGraph) AddEdge(v, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.e++
}

func (g *DirectedGraph) Adj(v int) []int {
	return g.adj[v]
}

func (g *DirectedGraph) reverse() *DirectedGraph {
	r := new(DirectedGraph)
	r.New(g.V())
	for i := 0; i < g.V(); i++ {
		for _, w := range g.Adj(i) {
			r.AddEdge(w, i)
		}
	}
	return r
}

func (g *DirectedGraph) String() string {
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

type DirectedDFS struct {
	marked []bool
}

func (d *DirectedDFS) Search(graph DirectedGraph, s int) {
	d.marked = make([]bool, graph.V())
	d.dfs(graph, s)
}

func (d *DirectedDFS) dfs(graph DirectedGraph, v int) {
	d.marked[v] = true
	for _, w := range graph.Adj(v) {
		if !d.hasMarked(w) {
			d.dfs(graph, w)
		}
	}
}

func (d DirectedDFS) hasMarked(w int) bool {
	return d.marked[w]
}
