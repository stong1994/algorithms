package str

import "testing"

func Test_manacher(t *testing.T) {
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
			args: args{s: "ababbac"},
			want: "abba",
		},
		{
			name: "example2",
			args: args{s: "abb"},
			want: "bb",
		},
		{
			name: "example3",
			args: args{s: "aabccbab"},
			want: "abccba",
		},
		{
			name: "example4",
			args: args{s: "bab"},
			want: "bab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := manacher(tt.args.s); got != tt.want {
				t.Errorf("manacher() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_preProcess(t *testing.T) {
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
			args: args{s: ""},
			want: "^$",
		},
		{
			name: "example2",
			args: args{s: "abc"},
			want: "^#a#b#c#$",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preProcess(tt.args.s); got != tt.want {
				t.Errorf("preProcess() = %v, want %v", got, tt.want)
			}
		})
	}
}
