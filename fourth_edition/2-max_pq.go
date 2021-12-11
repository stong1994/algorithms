package fourth_edition

import (
	"fmt"
	"math"
)

// 优先队列
// 书中分别用数组、链表、堆来实现，这里只用堆来实现
// 二叉堆的实现基于数组，首元素不使用。位置k的节点的父节点为k/2，子节点的位置分别为2k和2k+1
// 堆只保证父节点大于两个子节点，而不保证兄弟节点的大小顺序

type MaxPQ struct {
	pq Comparables
	n  int // 元素个数, 则len(pq)=n+1
}

// 构造堆时，可以将元素依次通过上浮加入队列
// 但是先将元素放到数组中，对数组进行排序，此时可以跳过叶子结点，因为对叶子节点进行下沉是无用的，
// 然后从N/2到1的节点依次进行下沉，这样更快
func NewMaxPQ(data []Comparable) *MaxPQ {
	dst := make([]Comparable, len(data)+1)
	copy(dst[1:], data)
	BaseSort{Quick{}}.Sort(dst[1:])

	m := MaxPQ{
		pq: dst,
		n:  len(data),
	}
	for i := m.n / 2; i >= 1; i-- {
		m.sink(i)
	}
	return &m
}

func (m MaxPQ) next(k []int) []int {
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

func (m MaxPQ) Show() {
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

func (m MaxPQ) isEmp() bool {
	return m.n == 0
}

func (m MaxPQ) size() int {
	return m.n
}

// 插入堆时，将元素放入末尾，然后对其进行“上浮”
func (m *MaxPQ) insert(k Comparable) {
	m.n++
	m.pq = append(m.pq, k)
	m.swim(m.n)
}

// 删除最大元素时，将堆顶的元素取出，然后将末尾的元素放到堆顶，然后对其进行“下沉”
func (m *MaxPQ) delMax() Comparable {
	if m.isEmp() {
		return comparableNull
	}
	max := m.pq[1]
	Exch(m.pq, 1, m.n)
	m.pq[m.n] = comparableNull
	m.n--
	m.sink(1)
	return max
}

// 上浮
func (m *MaxPQ) swim(k int) {
	for k > 1 && m.pq[k/2].CompareTo(m.pq[k]) < 0 {
		Exch(m.pq, k/2, k)
		k /= 2
	}
}

// 下沉
func (m MaxPQ) sink(k int) {
	for k*2 <= m.n {
		j := 2 * k
		if j < m.n && m.pq[j].CompareTo(m.pq[j+1]) < 0 {
			j++
		}
		if m.pq[k].CompareTo(m.pq[j]) >= 0 {
			break
		}
		Exch(m.pq, k, j)
		k = j
	}
}
