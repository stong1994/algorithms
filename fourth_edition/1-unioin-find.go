package fourth_edition

import (
	"fmt"
)

/** 动态连通性问题
编写一个程序来过滤调序列中所有无意义的整数对。
当程序从输入中读取了整数对p q时，如果已知所有整数对都不能说明p和q是联通的，那么将这一对整数写入输出。
否则忽略p和q，并继续处理下一对整数。
应用：通过已有的连接建立两个点的通信或者找到人际关系网
*/

// 问题抽象
type IUF interface {
	Union(uf *UF, p, q int)
	Find(uf *UF, p int) int
	Init(n int)
}

type UF struct {
	id    []int // 分量id（以触点作为索引）
	count int   // 分量数量
	iuf   IUF
}

// 初始化并保证没有连通分量
func (uf *UF) Init(n int) {
	uf.count = n
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}
	uf.id = id
	uf.iuf.Init(n)
}

func (uf UF) Count() int {
	return uf.count
}

func (uf UF) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

func (uf *UF) Union(p, q int) {
	uf.iuf.Union(uf, p, q)
}

func (uf *UF) Find(p int) int {
	return uf.iuf.Find(uf, p)
}

func ExecUF(uf *UF, n int, stdIn [][2]int) {
	uf.Init(n)
	for _, v := range stdIn {
		p := v[0]
		q := v[1]
		fmt.Println("p", p, "q", q)
		if uf.Connected(p, q) {
			fmt.Println("p connected q", p, q)
			continue
		}
		uf.Union(p, q)
	}
	fmt.Println("count", uf.Count())
}

// 算法实现
// 1. Quick-find
// 保证当id[p]==id[q]时p和q是连通的，即同一个连通分量上的触点在id中的值是相同的
// 当union p和q时，要将两个分量在id上的值设为相同
type quickFind struct{}

func (qf *quickFind) Init(n int) {}

// O(n)
func (qf *quickFind) Union(uf *UF, p, q int) {
	fmt.Println("p union q", p, q)
	pID := uf.Find(p)
	qID := uf.Find(q)
	if pID == qID {
		return
	}
	for i, v := range uf.id {
		if v == pID {
			uf.id[i] = qID
		}
	}
	uf.count--
	return
}

func (qf *quickFind) Find(uf *UF, p int) int {
	return uf.id[p]
}

// 2. Quick union
// 保证通过find能找到p的连通分量的根触点
type quickUnion struct{}

func (qu *quickUnion) Init(n int) {}

func (qu *quickUnion) Union(uf *UF, p, q int) {
	fmt.Println("p union q", p, q)
	pRoot := uf.Find(p)
	qRoot := uf.Find(q)
	if pRoot == qRoot {
		return
	}
	uf.id[p] = qRoot
	uf.count--
}

func (qu *quickUnion) Find(uf *UF, p int) int {
	for p != uf.id[p] {
		p = uf.id[p]
	}
	return p
}

// 3. weighted Quick union
// Quick union随机选择一颗树连接到另一颗树上，可以通过加权来优化,使得较小的树去连接较大的树
type weightedQuickUnion struct {
	weights map[int]int
}

func (qu *weightedQuickUnion) Init(n int) {
	w := make(map[int]int, n)
	for i := 0; i < n; i++ {
		w[i] = 1
	}
	qu.weights = w
}

func (qu *weightedQuickUnion) Union(uf *UF, p, q int) {
	fmt.Println("p union q", p, q)
	pRoot := uf.Find(p)
	qRoot := uf.Find(q)
	if pRoot == qRoot {
		return
	}
	pWeight := qu.weights[pRoot]
	qWeight := qu.weights[qRoot]
	if pWeight > qWeight {
		uf.id[q] = pRoot
		qu.weights[pRoot] += qWeight
	} else {
		uf.id[p] = qRoot
		qu.weights[qRoot] += pWeight
	}
	uf.count--
}

func (qu *weightedQuickUnion) Find(uf *UF, p int) int {
	for p != uf.id[p] {
		p = uf.id[p]
	}
	return p
}
