package str

import (
	"reflect"
	"testing"
)

func Test_strRotate(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example1",
			args: args{
				s1: "AABCD",
				s2: "CDAA",
			},
			want: true,
		},
		{
			name: "example2",
			args: args{
				s1: "ABCD",
				s2: "ACBD",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strRotate(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("strRotate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseShift(t *testing.T) {
	type args struct {
		arr []int
		N   int
		K   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example1",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				N:   5,
				K:   2,
			},
			want: []int{4, 5, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseShift(tt.args.arr, tt.args.N, tt.args.K)
			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("got %v but want %v", tt.args.arr, tt.want)
			}

		})
	}
}

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example1",
			args: args{s: "I am a student"},
			want: "student a am I",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
