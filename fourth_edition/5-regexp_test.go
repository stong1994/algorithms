package fourth_edition

import "testing"

func TestNFA_recognizes(t *testing.T) {


	tests := []struct {
		name   string
		nfa *NFA
		txt string
		want   bool
	}{
		{
			name: "abc",
			nfa: NewNFA("abc"),
			txt: "abc",
			want: true,
		},
		{
			name: "a.bc",
			nfa: NewNFA("a.bc"),
			txt: "adbc",
			want: true,
		},
		{
			name: "a*bc",
			nfa: NewNFA("a*bc"),
			txt: "aaaabc",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.nfa.recognizes(tt.txt); got != tt.want {
				t.Errorf("recognizes() = %v, want %v", got, tt.want)
			}
		})
	}
}
