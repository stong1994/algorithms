package classify

import (
	"reflect"
	"testing"
)

func Test_eraseOverlapIntervals(t *testing.T) {
	tests := []struct {
		name string
		intervals [][]int
		want int
	}{
		{
			name: "empty",
			intervals: nil,
			want: 0,
		},
		{
			name: "one",
			intervals: [][]int{{1,2}},
			want: 0,
		},
		{
			name: "least",
			intervals: [][]int{{1,2}, {1,3}},
			want: 1,
		},
		{
			name: "three",
			intervals: [][]int{{1,2}, {1,3}, {3,4}},
			want: 1,
		},
		{
			name: "four",
			intervals: [][]int{{1,2}, {2,3}, {3,4}, {1,3}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eraseOverlapIntervals(tt.intervals); got != tt.want {
				t.Errorf("eraseOverlapIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reconstructQueue(t *testing.T) {
	tests := []struct {
		name string
		people [][]int
		want [][]int
	}{
		{
			name: "example",
			people: [][]int{{7,0}, {4,4}, {7,1}, {5,0}, {6,1}, {5,2}},
			want: [][]int{{5,0}, {7,0}, {5,2}, {6,1}, {4,4}, {7,1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reconstructQueue(tt.people); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reconstructQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}