package search

import (
	"reflect"
	"testing"
)

func Test_mySqrt(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{
			name: "4",
			x:    4,
			want: 2,
		},
		{
			name: "8",
			x:    8,
			want: 2,
		},
		{
			name: "1",
			x:    1,
			want: 1,
		},
		{
			name: "2",
			x:    2,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextGreatestLetter(t *testing.T) {
	type args struct {
		letters []byte
		target  byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "example1",
			args: args{
				letters: []byte{'c', 'f', 'j'},
				target:  'a',
			},
			want: 'c',
		},
		{
			name: "example2",
			args: args{
				letters: []byte{'c', 'f', 'j'},
				target:  'c',
			},
			want: 'f',
		},
		{
			name: "example3",
			args: args{
				letters: []byte{'c', 'f', 'j'},
				target:  'd',
			},
			want: 'f',
		},
		{
			name: "example4",
			args: args{
				letters: []byte{'c', 'f', 'j'},
				target:  'g',
			},
			want: 'j',
		},
		{
			name: "example5",
			args: args{
				letters: []byte{'c', 'f', 'j'},
				target:  'j',
			},
			want: 'c',
		},
		{
			name: "example6",
			args: args{
				letters: []byte{'c', 'f', 'j'},
				target:  'k',
			},
			want: 'c',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreatestLetter(tt.args.letters, tt.args.target); got != tt.want {
				t.Errorf("nextGreatestLetter() = %x, want %x", got, tt.want)
			}
			if got := nextGreatestLetterOpt(tt.args.letters, tt.args.target); got != tt.want {
				t.Errorf("nextGreatestLetterOpt() = %x, want %x", got, tt.want)
			}
		})
	}
}

func Test_singleNonDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example1",
			nums: []int{1, 1, 2, 3, 3, 4, 4, 8, 8},
			want: 2,
		},
		{
			name: "example2",
			nums: []int{3, 3, 7, 7, 10, 11, 11},
			want: 10,
		},
		{
			name: "example3",
			nums: []int{1},
			want: 1,
		},
		{
			name: "example4",
			nums: []int{1, 1, 2},
			want: 2,
		},
		{
			name: "example5",
			nums: []int{1, 1, 2, 3, 3, 4, 4, 8, 8},
			want: 2,
		},
		{
			name: "example6",
			nums: []int{1, 1, 2, 3, 3},
			want: 2,
		},
		{
			name: "example7",
			nums: []int{1, 1, 2, 3, 3, 4, 4, 8, 8},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNonDuplicate(tt.nums); got != tt.want {
				t.Errorf("singleNonDuplicate() = %v, want %v", got, tt.want)
			}
			if got := singleNonDuplicateOpt(tt.nums); got != tt.want {
				t.Errorf("singleNonDuplicateOpt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMin(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example1",
			nums: []int{3, 4, 5, 1, 2},
			want: 1,
		},
		{
			name: "example2",
			nums: []int{4, 5, 6, 7, 0, 1, 2},
			want: 0,
		},
		{
			name: "example1",
			nums: []int{11, 13, 15, 17},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example1",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 8,
			},
			want: []int{3, 4},
		},
		{
			name: "example2",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 6,
			},
			want: []int{-1, -1},
		},
		{
			name: "example3",
			args: args{
				nums:   []int{},
				target: 0,
			},
			want: []int{-1, -1},
		},
		{
			name: "example4",
			args: args{
				nums:   []int{2, 2},
				target: 2,
			},
			want: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
