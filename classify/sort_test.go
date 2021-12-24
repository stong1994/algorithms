package classify

import (
	"reflect"
	"testing"
)

func Test_topN(t *testing.T) {
	type args struct {
		list []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "topN",
			args: args{
				[]int{2, 4, 1, 2, 9, 11, 5, 6, 3, 1, 0, 10, 8, -1},
				5,
			},
			want: 6,
		},
		{
			name: "topN2",
			args: args{
				[]int{3, 2, 1, 5, 6, 4},
				2,
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topN(tt.args.list, tt.args.k); got != tt.want {
				t.Errorf("topN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name  string
		args  args
		check func(got []int) bool
	}{
		{
			name: "empty",
			args: args{
				nums: nil,
				k:    1,
			},
			check: func(got []int) bool {
				return got == nil
			},
		},
		{
			name: "k > len(list)",
			args: args{
				nums: []int{1, 2, 3},
				k:    4,
			},
			check: func(got []int) bool {
				if len(got) != 3 {
					return false
				}
				return got[0]*got[1]*got[2] == 6
			},
		},
		{
			name: "k == len(list)",
			args: args{
				nums: []int{1, 2, 3},
				k:    3,
			},
			check: func(got []int) bool {
				if len(got) != 3 {
					return false
				}
				return got[0]*got[1]*got[2] == 6
			},
		},
		{
			name: "normal",
			args: args{
				nums: []int{1, 2, 2, 3, 3},
				k:    2,
			},
			check: func(got []int) bool {
				if len(got) != 2 {
					return false
				}
				return got[0]*got[1] == 6
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.args.nums, tt.args.k); !tt.check(got) {
				t.Errorf("failed: topKFrequent() = %v", got)
			}
		})
	}
}

func Test_frequencySort(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		check func(got string) bool
	}{
		{
			name: "empty",
			args: args{s: ""},
			check: func(got string) bool {
				return got == ""
			},
		},
		{
			name: "simple",
			args: args{s: "aab"},
			check: func(got string) bool {
				return got == "aab"
			},
		},
		{
			name: "cccaaa",
			args: args{s: "cccaaa"},
			check: func(got string) bool {
				return got == "cccaaa" || got == "aaaccc"
			},
		},
		{
			name: "tree",
			args: args{s: "tree"},
			check: func(got string) bool {
				return got == "eert" || got == "eetr"
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := frequencySort(tt.args.s); !tt.check(got) {
				t.Errorf("failed: frequencySort() = %v", got)
			}
		})
	}
}

func Test_sortColors(t *testing.T) {
	tests := []struct {
		name  string
		nums  []int
		check func(afterNums []int) bool
	}{
		{
			name: "empty",
			nums: []int{},
			check: func(afterNums []int) bool {
				return len(afterNums) == 0
			},
		},
		{
			name: "2,0,2,1,1,0",
			nums: []int{2, 0, 2, 1, 1, 0},
			check: func(afterNums []int) bool {
				dst := []int{0, 0, 1, 1, 2, 2}
				return reflect.DeepEqual(afterNums, dst)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.nums)
			if !tt.check(tt.nums) {
				t.Errorf("failed: name = %s, nums = %v", tt.name, tt.nums)
			}
		})
	}
}
