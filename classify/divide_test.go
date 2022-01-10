package classify

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func Test_diffWaysToCompute(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		want       []int
	}{
		{
			name:       "example1",
			expression: "2-1-1",
			want:       []int{0, 2},
		},
		{
			name:       "example2",
			expression: "2*3-4*5",
			want:       []int{-34, -10, -14, -10, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := diffWaysToCompute(tt.expression)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diffWaysToCompute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateTrees(t *testing.T) {
	tests := []struct {
		name string
		n    int
	}{
		{
			name: "example1",
			n:    3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateTrees(tt.n)
			for i, v := range got {
				printTreeNode(0, strconv.Itoa(i), v)
				fmt.Println()
			}
		})
	}
}
