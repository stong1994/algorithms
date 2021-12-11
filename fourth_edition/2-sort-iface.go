package fourth_edition

import "fmt"

type Comparable int

func (c Comparable) CompareTo(v Comparable) int {
	return int(c) - int(v)
}

var comparableNull Comparable = 0

type Comparables []Comparable

func Exch(data Comparables, m, n int) {
	data[m], data[n] = data[n], data[m]
}

type ISortable interface {
	Sort(list []Comparable)
}

type BaseSort struct {
	SortImpl ISortable
}

func (bs BaseSort) Sort(list []Comparable) {
	bs.SortImpl.Sort(list)
}

func (bs BaseSort) Less(j, k Comparable) bool {
	return j.CompareTo(k) < 0
}

func (bs BaseSort) Show(list []Comparable) {
	for _, v := range list {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func (bs BaseSort) IsSorted(list []Comparable) bool {
	for i := 1; i < len(list); i++ {
		if bs.Less(list[i], list[i-1]) {
			return false
		}
	}
	return true
}
