package classify

import (
	"reflect"
	"testing"
)

func Test_eraseOverlapIntervals(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      int
	}{
		{
			name:      "empty",
			intervals: nil,
			want:      0,
		},
		{
			name:      "one",
			intervals: [][]int{{1, 2}},
			want:      0,
		},
		{
			name:      "least",
			intervals: [][]int{{1, 2}, {1, 3}},
			want:      1,
		},
		{
			name:      "three",
			intervals: [][]int{{1, 2}, {1, 3}, {3, 4}},
			want:      1,
		},
		{
			name:      "four",
			intervals: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
			want:      1,
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
		name   string
		people [][]int
		want   [][]int
	}{
		{
			name:   "example",
			people: [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}},
			want:   [][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}},
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

func Test_maxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "example",
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			name:   "2-4-1",
			prices: []int{2, 4, 1},
			want:   2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canPlaceFlowers(t *testing.T) {
	type args struct {
		flowerbed []int
		n         int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example1",
			args: args{
				flowerbed: []int{1, 0, 0, 0, 1},
				n:         1,
			},
			want: true,
		},
		{
			name: "example2",
			args: args{
				flowerbed: []int{1, 0, 0, 0, 1},
				n:         2,
			},
			want: false,
		},
		{
			name: "example3",
			args: args{
				flowerbed: []int{1, 0, 0, 0, 0, 0, 1},
				n:         2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPlaceFlowers(tt.args.flowerbed, tt.args.n); got != tt.want {
				t.Errorf("canPlaceFlowers() = %v, want %v", got, tt.want)
			}
			if got := canPlaceFlowersOpt(tt.args.flowerbed, tt.args.n); got != tt.want {
				t.Errorf("canPlaceFlowersOpt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example1",
			args: args{
				s: "abc",
				t: "ahbgdc",
			},
			want: true,
		},
		{
			name: "example2",
			args: args{
				s: "axc",
				t: "ahbgdc",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkPossibility(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "example1",
			nums: []int{4, 2, 3},
			want: true,
		},
		{
			name: "example1",
			nums: []int{4, 2, 1},
			want: false,
		},
		{
			name: "example3",
			nums: []int{3, 4, 2, 3},
			want: false,
		},
		{
			name: "example4",
			nums: []int{5, 7, 1, 8},
			want: true,
		},
		{
			name: "example5",
			nums: []int{-1, 4, 2, 3},
			want: true,
		},
		{
			name: "example6",
			nums: []int{1, 4, 1, 2},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPossibility(tt.nums); got != tt.want {
				t.Errorf("checkPossibility() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example1",
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
		{
			name: "example2",
			nums: []int{1},
			want: 1,
		},
		{
			name: "example3",
			nums: []int{5, 4, -1, 7, 8},
			want: 23,
		},
		{
			name: "example4",
			nums: []int{-1},
			want: -1,
		},
		{
			name: "example5",
			nums: []int{1, 2},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
			if got := maxSubArrayOpt(tt.nums); got != tt.want {
				t.Errorf("maxSubArrayOpt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partitionLabels(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []int
	}{
		{
			name: "example1",
			s:    "ababcbacadefegdehijhklij",
			want: []int{9, 7, 8},
		},
		{
			name:"example2",
			s: "caedbdedda",
			want: []int{1,9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partitionLabels(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partitionLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
