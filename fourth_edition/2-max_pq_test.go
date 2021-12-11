package fourth_edition

import (
	"fmt"
	"testing"
)

func TestMaxPQ_Sort(t *testing.T) {
	pq := NewMaxPQ(sortList)
	pq.Show()
	var list []Comparable
	for !pq.isEmp() {
		max := pq.delMax()
		list = append(list, max)
	}
	fmt.Println("list", list)
	for i := 0; i < len(list)-2; i++ {
		if list[i] < list[i+1] {
			fmt.Println("not sorted")
			return
		}
	}
	fmt.Println("sorted")
}
