package math

import "testing"

func Test_convertToBase7(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example1",
			args: args{num: 100},
			want: "202",
		},
		{
			name: "example2",
			args: args{num: -7},
			want: "-10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToBase7(tt.args.num); got != tt.want {
				t.Errorf("convertToBase7() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toHex(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example1",
			args: args{num: 26},
			want: "1a",
		},
		{
			name: "example2",
			args: args{num: -1},
			want: "ffffffff",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toHex(tt.args.num); got != tt.want {
				t.Errorf("toHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertToTitle(t *testing.T) {
	type args struct {
		columnNumber int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example1",
			args: args{columnNumber: 1},
			want: "A",
		},
		{
			name: "example2",
			args: args{columnNumber: 28},
			want: "AB",
		},
		{
			name: "example3",
			args: args{columnNumber: 701},
			want: "ZY",
		},
		{
			name: "example4",
			args: args{columnNumber: 2147483647},
			want: "FXSHRXW",
		},
		{
			name: "example5",
			args: args{columnNumber: 1},
			want: "A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToTitle(tt.args.columnNumber); got != tt.want {
				t.Errorf("convertToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
