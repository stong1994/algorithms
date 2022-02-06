package str

import (
	"reflect"
	"testing"
)

func Test_getPMT(t *testing.T) {
	type args struct {
		pat string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example0",
			args: args{pat: "aa"},
			want: []int{0, 1},
		},
		{
			name: "example1",
			args: args{pat: "abc"},
			want: []int{0, 0, 0},
		},
		{
			name: "example2",
			args: args{pat: "abca"},
			want: []int{0, 0, 0, 1},
		},
		{
			name: "example3",
			args: args{pat: "abcabc"},
			want: []int{0, 0, 0, 1, 2, 3},
		},
		{
			name: "example4",
			args: args{pat: "abcabcd"},
			want: []int{0, 0, 0, 1, 2, 3, 0},
		},
		{
			name: "example5",
			args: args{pat: "aaa"},
			want: []int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPMT(tt.args.pat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPMT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getNext(t *testing.T) {
	type args struct {
		pat string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example0",
			args: args{pat: "aa"},
			want: []int{0, 0},
		},
		{
			name: "example1",
			args: args{pat: "abc"},
			want: []int{0, 0, 0},
		},
		{
			name: "example2",
			args: args{pat: "abca"},
			want: []int{0, 0, 0, 0},
		},
		{
			name: "example3",
			args: args{pat: "abcabc"},
			want: []int{0, 0, 0, 0, 1, 2},
		},
		{
			name: "example4",
			args: args{pat: "abcabcd"},
			want: []int{0, 0, 0, 0, 1, 2, 3},
		},
		{
			name: "example5",
			args: args{pat: "aaa"},
			want: []int{0, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNext_old(tt.args.pat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNext_old() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_match(t *testing.T) {
	type args struct {
		txt string
		pat string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				txt: "aaa",
				pat: "a",
			},
			want: 0,
		},
		{
			name: "example2",
			args: args{
				txt: "baaa",
				pat: "a",
			},
			want: 1,
		},
		{
			name: "example3",
			args: args{
				txt: "baaa",
				pat: "ab",
			},
			want: -1,
		},
		{
			name: "example4",
			args: args{
				txt: "abcabcd",
				pat: "abcd",
			},
			want: 3,
		},
		{
			name: "example5",
			args: args{
				txt: "aaacaaabaaaaaa",
				pat: "aaabaaaa",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := match_old1(tt.args.txt, tt.args.pat); got != tt.want {
				t.Errorf("match_old1() = %v, want %v", got, tt.want)
			}
			if got := match_old2(tt.args.txt, tt.args.pat); got != tt.want {
				t.Errorf("match_old2() = %v, want %v", got, tt.want)
			}
			if got := match(tt.args.txt, tt.args.pat); got != tt.want {
				t.Errorf("match() = %v, want %v", got, tt.want)
			}
		})
	}
}
