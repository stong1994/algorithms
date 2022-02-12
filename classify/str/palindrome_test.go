package str

import "testing"

func Test_findPalindrome(t *testing.T) {
	type args struct {
		s string
		l int
		r int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example1",
			args: args{
				s: "abbc",
				l: 1,
				r: 2,
			},
			want: "bb",
		},
		{
			name: "example1",
			args: args{
				s: "abba",
				l: 1,
				r: 2,
			},
			want: "abba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPalindrome(tt.args.s, tt.args.l, tt.args.r); got != tt.want {
				t.Errorf("findPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findLongestPalindrome(t *testing.T) {
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
			args: args{s: "abba"},
			want: "abba",
		},
		{
			name: "example2",
			args: args{s: "a"},
			want: "a",
		},
		{
			name: "example3",
			args: args{s: "cccabbaee"},
			want: "abba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLongestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("findLongestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
