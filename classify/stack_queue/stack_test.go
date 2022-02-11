package stack_queue

import (
	"reflect"
	"testing"
)

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example1",
			args: args{s: "()[]{}"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dailyTemperatures(t *testing.T) {
	type args struct {
		temperatures []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example1",
			args: args{temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73}},
			want: []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			name: "example2",
			args: args{temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73}},
			want: []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dailyTemperatures(tt.args.temperatures); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dailyTemperatures() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextGreaterElements(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example1",
			args: args{nums: []int{1, 2, 1}},
			want: []int{2, -1, 2},
		},
		{
			name: "example2",
			args: args{nums: []int{1, 2, 3, 4, 3}},
			want: []int{2, 3, 4, -1, 4},
		},
		{
			name: "example3",
			args: args{nums: []int{1, 2, 3, 2, 1}},
			want: []int{2, 3, -1, 3, 2},
		},
		{
			name: "example4",
			args: args{nums: []int{-1, 0}},
			want: []int{0, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElements(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
