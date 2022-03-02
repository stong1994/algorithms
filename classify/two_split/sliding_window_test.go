package two_split

import (
	"reflect"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{s: "abcabcbb"},
			want: 3,
		},
		{
			name: "example2",
			args: args{s: "bbbbb"},
			want: 1,
		},
		{
			name: "example3",
			args: args{s: "pwwkew"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring2(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findRepeatedDnaSequences(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "example1",
			args: args{s: "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"},
			want: []string{"AAAAACCCCC", "CCCCCAAAAA"},
		},
		{
			name: "example2",
			args: args{s: "AAAAAAAAAAAAA"},
			want: []string{"AAAAAAAAAA"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRepeatedDnaSequences(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRepeatedDnaSequences() = %v, want %v", got, tt.want)
			}
			if got := findRepeatedDnaSequences2(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRepeatedDnaSequences2() = %v, want %v", got, tt.want)
			}
		})
	}
}
