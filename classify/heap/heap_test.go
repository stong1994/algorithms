package heap

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MaxHeap(t *testing.T) {
	heap := NewMaxHeap()
	assert.True(t, heap.isEmpty())
	heap.insert(7)
	heap.insert(6)
	heap.insert(5)
	heap.insert(8)
	heap.insert(6)
	heap.insert(4)
	heap.insert(3)
	heap.insert(2)
	heap.insert(1)
	fmt.Println(heap.pq)
	assert.Equal(t, 8, heap.delMax())
	fmt.Println(heap.pq)
	assert.Equal(t, 7, heap.delMax())
	fmt.Println(heap.pq)
	assert.Equal(t, 6, heap.delMax())
	fmt.Println(heap.pq)
	assert.Equal(t, 6, heap.delMax())
	fmt.Println(heap.pq)
	assert.Equal(t, 5, heap.delMax())
	fmt.Println(heap.pq)
}

func Test_MinHeap(t *testing.T) {
	heap := NewMinHeap()
	assert.True(t, heap.isEmpty())
	heap.insert(7)
	heap.insert(4)
	heap.insert(6)
	heap.insert(3)
	heap.insert(5)
	heap.insert(1)
	heap.insert(8)
	heap.insert(6)
	heap.insert(2)
	fmt.Println(heap.pq)
	assert.Equal(t, 1, heap.delMin())
	fmt.Println(heap.pq)
	assert.Equal(t, 2, heap.delMin())
	fmt.Println(heap.pq)
	assert.Equal(t, 3, heap.delMin())
	fmt.Println(heap.pq)
	assert.Equal(t, 4, heap.delMin())
	fmt.Println(heap.pq)
	assert.Equal(t, 5, heap.delMin())
	fmt.Println(heap.pq)
}

func Test_kthSmallest(t *testing.T) {
	type args struct {
		matrix [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				matrix: [][]int{
					{1, 4, 7, 11, 15},
					{2, 5, 8, 12, 19},
					{3, 6, 9, 16, 22},
					{10, 13, 14, 17, 24},
					{18, 21, 23, 26, 30},
				},
				k: 5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.args.matrix, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
