package fourth_edition

import "testing"

func Test_violentSearchSubStr(t *testing.T) {
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
			name: "empty",
			args: args{
				txt: "",
				pat: "a",
			},
			want: -1,
		},
		{
			name: "bc",
			args: args{
				txt: "abc",
				pat: "bc",
			},
			want: 1,
		},
		{
			name: "dup",
			args: args{
				txt: "abcbc",
				pat: "bc",
			},
			want: 1,
		},
		{
			name: "abc",
			args: args{
				txt: "bc",
				pat: "abc",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := violentSearchSubStr(tt.args.txt, tt.args.pat); got != tt.want {
				t.Errorf("violentSearchSubStr() = %v, want %v", got, tt.want)
			}
			if got := violentSearchSubStr2(tt.args.txt, tt.args.pat); got != tt.want {
				t.Errorf("violentSearchSubStr2() = %v, want %v", got, tt.want)
			}
			if got := searchSubStrKMP(tt.args.txt, tt.args.pat); got != tt.want {
				t.Errorf("searchSubStrKMP() = %v, want %v", got, tt.want)
			}
			if got := searchSubStrBM(tt.args.txt, tt.args.pat); got != tt.want {
				t.Errorf("searchSubStrBM() = %v, want %v", got, tt.want)
			}
			rk := NewRabinKarp(tt.args.pat)
			if got := rk.search(tt.args.txt); got != tt.want {
				t.Errorf("RabinKarp() = %v, want %v", got, tt.want)
			}
		})
	}
}
